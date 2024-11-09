package basket

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	basketv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/basket/api/basket/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/docs"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/handler"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/entity"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/event"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/repository"
	infragrpc "github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/infrastructure/client/grpc"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/infrastructure/logging"
	persistencegorm "github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/infrastructure/persistence/gorm"
	grpcv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/interface/grpc/v1"
	restv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/interface/rest/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/broker/nats"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/di"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/es"
	evenstoregorm "github.com/Chengxufeng1994/event-driven-arch-in-go/internal/eventstore/gorm"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/logger"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/monolith"
	outboxstoregorm "github.com/Chengxufeng1994/event-driven-arch-in-go/internal/outboxstore/gorm"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry/serdes"
	snapshotstoregorm "github.com/Chengxufeng1994/event-driven-arch-in-go/internal/snapshotstore/gorm"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/tm"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
)

type Module struct{}

var _ monolith.Module = (*Module)(nil)

func NewModule() *Module { return &Module{} }

func (m *Module) PrepareRun(ctx context.Context, mono monolith.Monolith) error {
	container := di.New()
	// setup Driven adapters
	endpoint := fmt.Sprintf("%s:%d", mono.Config().Server.GPPC.Host, mono.Config().Server.GPPC.Port)
	container.AddSingleton("registry", func(c di.Container) (any, error) {
		reg := registry.New()
		if err := registrations(reg); err != nil {
			return nil, err
		}
		if err := basketv1.Registrations(reg); err != nil {
			return nil, err
		}
		if err := storev1.Registrations(reg); err != nil {
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
		return infragrpc.Dial(ctx, endpoint)
	})
	container.AddSingleton("outboxProcessor", func(c di.Container) (any, error) {
		return tm.NewOutboxProcessor(
			c.Get("stream").(am.RawMessageStream),
			outboxstoregorm.NewOutboxStore("baskets.outbox", c.Get("db").(*gorm.DB)),
		), nil
	})
	container.AddScoped("tx", func(c di.Container) (any, error) {
		db := c.Get("db").(*gorm.DB)
		return db.Begin(), nil
	})
	container.AddScoped("txStream", func(c di.Container) (any, error) {
		tx := c.Get("tx").(*gorm.DB)
		outboxStore := outboxstoregorm.NewOutboxStore("baskets.outbox", tx)
		return am.RawMessageStreamWithMiddleware(
			c.Get("stream").(am.RawMessageStream),
			tm.NewOutboxStreamMiddleware(outboxStore),
		), nil
	})

	container.AddScoped("eventStream", func(c di.Container) (any, error) {
		return am.NewEventStream(c.Get("registry").(registry.Registry), c.Get("txStream").(am.RawMessageStream)), nil
	})
	container.AddScoped("inboxMiddleware", func(c di.Container) (any, error) {
		tx := c.Get("tx").(*gorm.DB)
		inboxStore := outboxstoregorm.NewInboxStore("baskets.inbox", tx)
		return tm.NewInboxHandlerMiddleware(inboxStore), nil
	})
	container.AddScoped("aggregateStore", func(c di.Container) (any, error) {
		tx := c.Get("tx").(*gorm.DB)
		reg := c.Get("registry").(registry.Registry)
		return es.AggregateStoreWithMiddleware(
			evenstoregorm.NewEventStore("baskets.events", tx, reg),
			snapshotstoregorm.NewSnapshotStore("baskets.snapshots", tx, reg),
		), nil
	})
	container.AddScoped("baskets", func(c di.Container) (any, error) {
		return es.NewAggregateRepository[*aggregate.Basket](
				aggregate.BasketAggregate,
				c.Get("registry").(registry.Registry),
				c.Get("aggregateStore").(es.AggregateStore)),
			nil
	})
	container.AddScoped("stores", func(c di.Container) (any, error) {
		return persistencegorm.NewGormStoreCacheRepository(
				c.Get("tx").(*gorm.DB),
				infragrpc.NewGrpcStoreClient(c.Get("conn").(*grpc.ClientConn))),
			nil
	})
	container.AddScoped("products", func(c di.Container) (any, error) {
		return persistencegorm.NewGormProductCacheRepository(
				c.Get("tx").(*gorm.DB),
				infragrpc.NewGrpcProductClient(c.Get("conn").(*grpc.ClientConn))),
			nil
	})

	// setup application
	container.AddScoped("app", func(c di.Container) (any, error) {
		return logging.NewLogApplicationAccess(
			application.NewBasketApplication(
				c.Get("baskets").(es.AggregateRepository[*aggregate.Basket]),
				c.Get("stores").(repository.StoreCacheRepository),
				c.Get("products").(repository.ProductCacheRepository),
				c.Get("domainEventDispatcher").(ddd.EventDispatcher[ddd.Event]),
			),
			c.Get("logger").(logger.Logger)), nil
	})
	container.AddScoped("domainEventHandlers", func(c di.Container) (any, error) {
		return logging.NewLogEventHandlerAccess[ddd.Event](
			handler.NewDomainEventHandler(c.Get("eventStream").(am.EventStream)),
			"DomainEvents", c.Get("logger").(logger.Logger),
		), nil
	})
	container.AddScoped("integrationEventHandlers", func(c di.Container) (any, error) {
		return logging.NewLogEventHandlerAccess[ddd.Event](
			handler.NewIntegrationEventHandler(
				c.Get("stores").(repository.StoreCacheRepository),
				c.Get("products").(repository.ProductCacheRepository)),
			"IntegrationEvents", c.Get("logger").(logger.Logger),
		), nil
	})

	// setup Driver adapters
	if err := grpcv1.RegisterServerTx(container, mono.RPC().GRPCServer()); err != nil {
		return err
	}
	if err := restv1.RegisterGateway(ctx, mono.Gin(), endpoint); err != nil {
		return err
	}
	if err := docs.RegisterSwagger(mono.Gin()); err != nil {
		return err
	}
	handler.RegisterDomainEventHandlersTx(container)
	if err := handler.RegisterIntegrationEventHandlersTx(container); err != nil {
		return err
	}
	go m.startOutboxProcessor(ctx, container)

	return nil
}

func (m *Module) Name() string {
	return "basket"
}

func registrations(reg registry.Registry) error {
	serde := serdes.NewJSONSerde(reg)
	// Basket
	if err := serde.Register(&aggregate.Basket{}, func(v interface{}) error {
		basket := v.(*aggregate.Basket)
		basket.Items = make(map[string]*entity.Item)
		return nil
	}); err != nil {
		return err
	}

	// basket events
	if err := serde.Register(event.BasketStarted{}); err != nil {
		return err
	}
	if err := serde.Register(event.BasketCanceled{}); err != nil {
		return err
	}
	if err := serde.Register(event.BasketCheckedOut{}); err != nil {
		return err
	}
	if err := serde.Register(event.BasketItemAdded{}); err != nil {
		return err
	}
	if err := serde.Register(event.BasketItemRemoved{}); err != nil {
		return err
	}

	// basket snapshots
	if err := serde.RegisterKey(aggregate.BasketV1{}.SnapshotName(), aggregate.BasketV1{}); err != nil {
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
		logger.WithError(err).Error("baskets outbox processor encountered an error")
	}
}
