package ordering

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"

	basketv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/basket/api/basket/v1"
	depotv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/depot/api/depot/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/broker/nats"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/di"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/es"
	evenstoregorm "github.com/Chengxufeng1994/event-driven-arch-in-go/internal/eventstore/gorm"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/logger"
	outboxstoregorm "github.com/Chengxufeng1994/event-driven-arch-in-go/internal/outboxstore/gorm"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry/serdes"
	snapshotstoregorm "github.com/Chengxufeng1994/event-driven-arch-in-go/internal/snapshotstore/gorm"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/system"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/tm"
	orderv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/api/order/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/docs"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application/handler"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application/usecase"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/event"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/infrastructure/client/grpc"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/infrastructure/logging"
	grpcv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/interface/grpc/v1"
	v1 "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/interface/rest/v1"
)

type Module struct{}

var _ system.Module = (*Module)(nil)

func NewModule() *Module { return &Module{} }

func (m *Module) Startup(ctx context.Context, mono system.Service) error {
	container := di.New()
	// setup Driver adapters
	endpoint := fmt.Sprintf("%s:%d", mono.Config().Server.GPPC.Host, mono.Config().Server.GPPC.Port)
	container.AddSingleton("registry", func(c di.Container) (any, error) {
		reg := registry.New()
		if err := registrations(reg); err != nil {
			return nil, err
		}
		if err := orderv1.Registrations(reg); err != nil {
			return nil, err
		}
		if err := basketv1.Registrations(reg); err != nil {
			return nil, err
		}
		if err := depotv1.Registrations(reg); err != nil {
			return nil, err
		}
		return reg, nil
	})
	container.AddSingleton("logger", func(c di.Container) (any, error) {
		return mono.Logger(), nil
	})
	container.AddSingleton("stream", func(c di.Container) (any, error) {
		return nats.NewStream(mono.Config().Infrastructure.Nats.Stream, mono.JetStream(), mono.Logger()), nil
	})
	container.AddSingleton("domainEventDispatcher", func(c di.Container) (any, error) {
		return ddd.NewEventDispatcher[ddd.Event](), nil
	})
	container.AddSingleton("db", func(c di.Container) (any, error) {
		return mono.Database(), nil
	})
	container.AddSingleton("conn", func(c di.Container) (any, error) {
		return grpc.Dial(ctx, endpoint)
	})
	container.AddSingleton("outboxProcessor", func(c di.Container) (any, error) {
		return tm.NewOutboxProcessor(
			c.Get("stream").(am.RawMessageStream),
			outboxstoregorm.NewOutboxStore("ordering.outbox", c.Get("db").(*gorm.DB)),
		), nil
	})
	container.AddScoped("tx", func(c di.Container) (any, error) {
		db := c.Get("db").(*gorm.DB)
		return db.Begin(), nil
	})
	container.AddScoped("txStream", func(c di.Container) (any, error) {
		tx := c.Get("tx").(*gorm.DB)
		outboxStore := outboxstoregorm.NewOutboxStore("ordering.outbox", tx)
		return am.RawMessageStreamWithMiddleware(
			c.Get("stream").(am.RawMessageStream),
			tm.NewOutboxStreamMiddleware(outboxStore),
		), nil
	})
	container.AddScoped("eventStream", func(c di.Container) (any, error) {
		return am.NewEventStream(c.Get("registry").(registry.Registry), c.Get("txStream").(am.RawMessageStream)), nil
	})
	container.AddScoped("replyStream", func(c di.Container) (any, error) {
		return am.NewReplyStream(c.Get("registry").(registry.Registry), c.Get("txStream").(am.RawMessageStream)), nil
	})
	container.AddScoped("inboxMiddleware", func(c di.Container) (any, error) {
		tx := c.Get("tx").(*gorm.DB)
		inboxStore := outboxstoregorm.NewInboxStore("ordering.inbox", tx)
		return tm.NewInboxHandlerMiddleware(inboxStore), nil
	})
	container.AddScoped("aggregateStore", func(c di.Container) (any, error) {
		tx := c.Get("tx").(*gorm.DB)
		reg := c.Get("registry").(registry.Registry)
		return es.AggregateStoreWithMiddleware(
			evenstoregorm.NewEventStore("ordering.events", tx, reg),
			snapshotstoregorm.NewSnapshotStore("ordering.snapshots", tx, reg),
		), nil
	})
	container.AddScoped("orders", func(c di.Container) (any, error) {
		aggregateStore := c.Get("aggregateStore").(es.AggregateStore)
		reg := c.Get("registry").(registry.Registry)
		return es.NewAggregateRepository[*aggregate.Order](aggregate.OrderAggregate, reg, aggregateStore), nil
	})

	// setup application
	container.AddScoped("app", func(c di.Container) (any, error) {
		return logging.NewLogApplicationAccess(
			application.NewOrderApplication(
				c.Get("orders").(repository.OrderRepository),
				c.Get("domainEventDispatcher").(ddd.EventPublisher[ddd.Event]),
			),
			c.Get("logger").(logger.Logger)), nil
	})
	// setup application handlers
	container.AddScoped("domainEventHandlers", func(c di.Container) (any, error) {
		return logging.NewLogEventHandlerAccess[ddd.Event](
			handler.NewDomainEventHandler(c.Get("eventStream").(am.EventStream)),
			"DomainEvents", c.Get("logger").(logger.Logger),
		), nil
	})
	container.AddScoped("integrationEventHandlers", func(c di.Container) (any, error) {
		return logging.NewLogEventHandlerAccess[ddd.Event](
			handler.NewIntegrationEventHandlers(c.Get("app").(usecase.OrderUseCase)),
			"IntegrationEvents", c.Get("logger").(logger.Logger),
		), nil
	})
	container.AddScoped("commandHandlers", func(c di.Container) (any, error) {
		return logging.NewLogCommandHandlerAccess[ddd.Command](
			handler.NewCommandHandlers(c.Get("app").(usecase.OrderUseCase)),
			"Commands", c.Get("logger").(logger.Logger),
		), nil
	})

	// setup Driver adapters
	if err := grpcv1.RegisterServerTx(container, mono.RPC().GRPCServer()); err != nil {
		return err
	}
	if err := v1.RegisterGateway(ctx, mono.Gin(), endpoint); err != nil {
		return err
	}
	if err := docs.RegisterSwagger(mono.Gin()); err != nil {
		return err
	}
	handler.RegisterDomainEventHandlersTx(container)
	if err := handler.RegisterIntegrationEventHandlersTx(container); err != nil {
		return err
	}
	if err := handler.RegisterCommandHandlersTx(container); err != nil {
		return err
	}

	go m.startOutboxProcessor(ctx, container)

	return nil
}

func (m *Module) Name() string {
	return "order"
}

func registrations(reg registry.Registry) (err error) {
	serde := serdes.NewJSONSerde(reg)

	// Order
	if err := serde.Register(aggregate.Order{}, func(v any) error {
		order := v.(*aggregate.Order)
		order.Aggregate = es.NewAggregate("", aggregate.OrderAggregate)
		return nil
	}); err != nil {
		return err
	}

	// order events
	if err := serde.Register(event.OrderCreated{}); err != nil {
		return err
	}
	if err := serde.Register(event.OrderRejected{}); err != nil {
		return err
	}
	if err := serde.Register(event.OrderApproved{}); err != nil {
		return err
	}
	if err := serde.Register(event.OrderReadied{}); err != nil {
		return err
	}
	if err := serde.Register(event.OrderCanceled{}); err != nil {
		return err
	}
	if err := serde.Register(event.OrderCompleted{}); err != nil {
		return err
	}

	// order snapshots
	if err := serde.RegisterKey(aggregate.OrderV1{}.SnapshotName(), aggregate.OrderV1{}); err != nil {
		return err
	}

	return nil
}

func (m *Module) startOutboxProcessor(ctx context.Context, container di.Container) {
	outboxProcessor := container.Get("outboxProcessor").(tm.OutboxProcessor)
	logger := container.Get("logger").(logger.Logger)

	eg := errgroup.Group{}
	eg.Go(func() error {
		return outboxProcessor.Start(ctx)
	})

	eg.Go(func() error {
		<-ctx.Done()
		return nil
	})

	if err := eg.Wait(); err != nil {
		logger.WithError(err).Error("ordering outbox processor encountered an error")
	}
}
