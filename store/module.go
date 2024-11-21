package store

import (
	"context"

	"golang.org/x/sync/errgroup"
	pkggorm "gorm.io/gorm"

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
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/docs"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/application"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/application/handler"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/event"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/infrastructure/logging"
	persistencegorm "github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/infrastructure/persistence/gorm"
	v1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/interface/grpc/v1"
	restv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/interface/rest/v1"
)

type Module struct{}

var _ system.Module = (*Module)(nil)

func NewModule() *Module { return &Module{} }

func Root(ctx context.Context, svc system.Service) error {
	container := di.New()

	// setup Driven adapters
	container.AddSingleton("registry", func(c di.Container) (any, error) {
		reg := registry.New()
		if err := registrations(reg); err != nil {
			return nil, err
		}
		if err := storev1.Registrations(reg); err != nil {
			return nil, err
		}
		return reg, nil
	})
	container.AddSingleton("logger", func(c di.Container) (any, error) {
		return svc.Logger(), nil
	})
	container.AddSingleton("stream", func(c di.Container) (any, error) {
		return nats.NewStream(svc.Config().Infrastructure.Nats.Stream, svc.JetStream(), svc.Logger()), nil
	})
	container.AddSingleton("domainEventDispatcher", func(c di.Container) (any, error) {
		return ddd.NewEventDispatcher[ddd.Event](), nil
	})
	container.AddSingleton("db", func(c di.Container) (any, error) {
		return svc.Database(), nil
	})
	container.AddSingleton("outboxProcessor", func(c di.Container) (any, error) {
		return tm.NewOutboxProcessor(
			c.Get("stream").(am.RawMessageStream),
			outboxstoregorm.NewOutboxStore("stores.outbox", c.Get("db").(*pkggorm.DB)),
		), nil
	})
	container.AddScoped("tx", func(c di.Container) (any, error) {
		db := c.Get("db").(*pkggorm.DB)
		return db.Begin(), nil
	})
	container.AddScoped("txStream", func(c di.Container) (any, error) {
		tx := c.Get("tx").(*pkggorm.DB)
		outboxStore := outboxstoregorm.NewOutboxStore("stores.outbox", tx)
		return am.RawMessageStreamWithMiddleware(
			c.Get("stream").(am.RawMessageStream),
			tm.NewOutboxStreamMiddleware(outboxStore),
		), nil
	})
	container.AddScoped("eventStream", func(c di.Container) (any, error) {
		return am.NewEventStream(c.Get("registry").(registry.Registry), c.Get("txStream").(am.RawMessageStream)), nil
	})
	container.AddScoped("aggregateStore", func(c di.Container) (any, error) {
		tx := c.Get("tx").(*pkggorm.DB)
		reg := c.Get("registry").(registry.Registry)
		return es.AggregateStoreWithMiddleware(
			evenstoregorm.NewEventStore("stores.events", tx, reg),
			snapshotstoregorm.NewSnapshotStore("stores.snapshots", tx, reg),
		), nil
	})
	container.AddScoped("stores", func(c di.Container) (any, error) {
		return es.NewAggregateRepository[*aggregate.Store](
			aggregate.StoreAggregate,
			c.Get("registry").(registry.Registry),
			c.Get("aggregateStore").(es.AggregateStore),
		), nil
	})
	container.AddScoped("products", func(c di.Container) (any, error) {
		return es.NewAggregateRepository[*aggregate.Product](
			aggregate.ProductAggregate,
			c.Get("registry").(registry.Registry),
			c.Get("aggregateStore").(es.AggregateStore),
		), nil
	})
	container.AddScoped("catalog", func(c di.Container) (any, error) {
		return persistencegorm.NewGormCatalogRepository(c.Get("tx").(*pkggorm.DB)), nil
	})
	container.AddScoped("mall", func(c di.Container) (any, error) {
		return persistencegorm.NewGormMallRepository(c.Get("tx").(*pkggorm.DB)), nil
	})

	// setup application
	container.AddScoped("app", func(c di.Container) (any, error) {
		return logging.NewLogApplicationAccess(
			application.NewStoreApplication(
				c.Get("stores").(repository.StoreRepository),
				c.Get("products").(repository.ProductRepository),
				c.Get("mall").(repository.MallRepository),
				c.Get("catalog").(repository.CatalogRepository),
				c.Get("domainEventDispatcher").(ddd.EventDispatcher[ddd.Event]),
			),
			c.Get("logger").(logger.Logger),
		), nil
	})
	container.AddScoped("catalogHandlers", func(c di.Container) (any, error) {
		return logging.NewLogEventHandlerAccess[ddd.Event](
			handler.NewCatalogDomainEventHandler(c.Get("catalog").(repository.CatalogRepository)),
			"Catalog", c.Get("logger").(logger.Logger),
		), nil
	})
	container.AddScoped("mallHandlers", func(c di.Container) (any, error) {
		return logging.NewLogEventHandlerAccess[ddd.Event](
			handler.NewMallDomainEventHandler(c.Get("mall").(repository.MallRepository)),
			"Mall", c.Get("logger").(logger.Logger),
		), nil
	})
	container.AddScoped("domainEventHandlers", func(c di.Container) (any, error) {
		return logging.NewLogEventHandlerAccess[ddd.Event](
			handler.NewDomainEventHandler(c.Get("eventStream").(am.EventStream)),
			"DomainEvents", c.Get("logger").(logger.Logger),
		), nil
	})

	// setup Driver adapters
	if err := v1.RegisterServerTx(container, svc.RPC().GRPCServer()); err != nil {
		return err
	}
	if err := restv1.RegisterGateway(ctx, svc.Gin(), svc.Config().Server.GPPC.Address()); err != nil {
		return err
	}
	if err := docs.RegisterSwagger(svc.Gin()); err != nil {
		return err
	}
	handler.RegisterCatalogDomainEventHandlersTx(container)
	handler.RegisterMallDomainEventHandlersTx(container)
	handler.RegisterDomainEventHandlersTx(container)
	if err := storev1.RegisterAsyncAPI(svc.Gin()); err != nil {
		return err
	}

	go startOutboxProcessor(ctx, container)

	return nil
}

func (m *Module) Startup(ctx context.Context, svc system.Service) error {
	return Root(ctx, svc)
}

func (m *Module) Name() string {
	return "stores"
}

func registrations(reg registry.Registry) error {
	serde := serdes.NewJSONSerde(reg)

	// store
	if err := serde.Register(aggregate.Store{}, func(v any) error {
		store := v.(*aggregate.Store)
		store.Aggregate = es.NewAggregate("", aggregate.StoreAggregate)
		return nil
	}); err != nil {
		return err
	}
	// store events
	if err := serde.Register(event.StoreCreated{}); err != nil {
		return err
	}
	if err := serde.RegisterKey(event.StoreParticipationEnabledEvent, event.StoreParticipationToggled{}); err != nil {
		return err
	}
	if err := serde.RegisterKey(event.StoreParticipationDisabledEvent, event.StoreParticipationToggled{}); err != nil {
		return err
	}
	if err := serde.Register(event.StoreRebranded{}); err != nil {
		return err
	}
	// store snapshots
	if err := serde.RegisterKey(aggregate.StoreV1{}.SnapshotName(), aggregate.StoreV1{}); err != nil {
		return err
	}

	// product
	if err := serde.Register(aggregate.Product{}, func(v any) error {
		product := v.(*aggregate.Product)
		product.Aggregate = es.NewAggregate("", aggregate.ProductAggregate)
		return nil
	}); err != nil {
		return err
	}
	// product events
	if err := serde.Register(event.ProductAdded{}); err != nil {
		return err
	}
	if err := serde.Register(event.ProductRebranded{}); err != nil {
		return err
	}
	if err := serde.RegisterKey(event.ProductPriceIncreasedEvent, event.ProductPriceChanged{}); err != nil {
		return err
	}
	if err := serde.RegisterKey(event.ProductPriceDecreasedEvent, event.ProductPriceChanged{}); err != nil {
		return err
	}
	if err := serde.Register(event.ProductRemoved{}); err != nil {
		return err
	}
	// product snapshots
	if err := serde.RegisterKey(aggregate.ProductV1{}.SnapshotName(), aggregate.ProductV1{}); err != nil {
		return err
	}

	return nil
}

func startOutboxProcessor(ctx context.Context, container di.Container) {
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
		logger.WithError(err).Error("stores outbox processor encountered an error")
	}
}
