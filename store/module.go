package store

import (
	"context"

	"golang.org/x/sync/errgroup"
	pkggorm "gorm.io/gorm"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/amotel"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/amprom"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/broker/nats"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/di"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/es"
	evenstoregorm "github.com/Chengxufeng1994/event-driven-arch-in-go/internal/eventstore/gorm"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/logger"
	outboxgorm "github.com/Chengxufeng1994/event-driven-arch-in-go/internal/outboxstore/gorm"
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
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/infrastructure/constants"
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
	container.AddSingleton(constants.RegistryKey, func(c di.Container) (any, error) {
		reg := registry.New()
		if err := registrations(reg); err != nil {
			return nil, err
		}
		if err := storev1.Registrations(reg); err != nil {
			return nil, err
		}
		return reg, nil
	})
	stream := nats.NewStream(svc.Config().Infrastructure.Nats.Stream, svc.JetStream(), svc.Logger())
	container.AddSingleton(constants.DomainDispatcherKey, func(c di.Container) (any, error) {
		return ddd.NewEventDispatcher[ddd.Event](), nil
	})
	container.AddScoped(constants.DatabaseTransactionKey, func(c di.Container) (any, error) {
		return svc.Database().Begin(), nil
	})
	sentCounter := amprom.SentMessagesCounter(constants.ServiceName)
	container.AddScoped(constants.MessagePublisherKey, func(c di.Container) (any, error) {
		tx := c.Get(constants.DatabaseTransactionKey).(*pkggorm.DB)
		outboxstore := outboxgorm.NewOutboxStore(constants.OutboxTableName, tx)
		return am.NewMessagePublisher(
			stream,
			amotel.OtelMessageContextInjector(),
			sentCounter,
			tm.OutboxPublisher(outboxstore),
		), nil
	})
	container.AddSingleton(constants.MessageSubscriberKey, func(c di.Container) (any, error) {
		return am.NewMessageSubscriber(
			stream,
			amotel.OtelMessageContextExtractor(),
			amprom.ReceivedMessagesCounter(constants.ServiceName),
		), nil
	})
	container.AddScoped(constants.EventPublisherKey, func(c di.Container) (any, error) {
		return am.NewEventPublisher(
			c.Get(constants.RegistryKey).(registry.Registry),
			c.Get(constants.MessagePublisherKey).(am.MessagePublisher),
		), nil
	})
	container.AddScoped(constants.InboxStoreKey, func(c di.Container) (any, error) {
		tx := c.Get(constants.DatabaseTransactionKey).(*pkggorm.DB)
		return outboxgorm.NewInboxStore(constants.InboxTableName, tx), nil
	})
	container.AddScoped(constants.AggregateStoreKey, func(c di.Container) (any, error) {
		tx := c.Get(constants.DatabaseTransactionKey).(*pkggorm.DB)
		reg := c.Get(constants.RegistryKey).(registry.Registry)
		return es.AggregateStoreWithMiddleware(
			evenstoregorm.NewEventStore(constants.EventsTableName, tx, reg),
			snapshotstoregorm.NewSnapshotStore(constants.SnapshotsTableName, tx, reg),
		), nil
	})
	container.AddScoped(constants.StoresRepoKey, func(c di.Container) (any, error) {
		return es.NewAggregateRepository[*aggregate.Store](
			aggregate.StoreAggregate,
			c.Get(constants.RegistryKey).(registry.Registry),
			c.Get(constants.AggregateStoreKey).(es.AggregateStore),
		), nil
	})
	container.AddScoped(constants.ProductsRepoKey, func(c di.Container) (any, error) {
		return es.NewAggregateRepository[*aggregate.Product](
			aggregate.ProductAggregate,
			c.Get(constants.RegistryKey).(registry.Registry),
			c.Get(constants.AggregateStoreKey).(es.AggregateStore),
		), nil
	})
	container.AddScoped(constants.CatalogRepoKey, func(c di.Container) (any, error) {
		return persistencegorm.NewGormCatalogRepository(
			c.Get(constants.DatabaseTransactionKey).(*pkggorm.DB),
		), nil
	})
	container.AddScoped(constants.MallRepoKey, func(c di.Container) (any, error) {
		return persistencegorm.NewGormMallRepository(
			c.Get(constants.DatabaseTransactionKey).(*pkggorm.DB),
		), nil
	})

	// setup application
	container.AddScoped(constants.ApplicationKey, func(c di.Container) (any, error) {
		return application.New(
			c.Get(constants.StoresRepoKey).(repository.StoreRepository),
			c.Get(constants.ProductsRepoKey).(repository.ProductRepository),
			c.Get(constants.MallRepoKey).(repository.MallRepository),
			c.Get(constants.CatalogRepoKey).(repository.CatalogRepository),
			c.Get(constants.DomainDispatcherKey).(ddd.EventPublisher[ddd.Event]),
		), nil
	})
	container.AddScoped(constants.CatalogHandlersKey, func(c di.Container) (any, error) {
		return handler.NewCatalogDomainEventHandler(c.Get(constants.CatalogRepoKey).(repository.CatalogRepository)), nil
	})
	container.AddScoped(constants.MallHandlersKey, func(c di.Container) (any, error) {
		return handler.NewMallDomainEventHandler(c.Get(constants.MallRepoKey).(repository.MallRepository)), nil
	})
	container.AddScoped(constants.DomainEventHandlersKey, func(c di.Container) (any, error) {
		return handler.NewDomainEventHandlers(c.Get(constants.EventPublisherKey).(am.EventPublisher)), nil
	})
	outboxProcessor := tm.NewOutboxProcessor(
		stream,
		outboxgorm.NewOutboxStore(constants.OutboxTableName, svc.Database()),
	)

	// setup Driver adapters
	if err := v1.RegisterServerTx(container, svc.RPC().GRPCServer()); err != nil {
		return err
	}
	if err := restv1.RegisterGateway(ctx, svc.Gin(), svc.Config().Server.GRPC.Address()); err != nil {
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

	go startOutboxProcessor(ctx, outboxProcessor, svc.Logger())

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

func startOutboxProcessor(ctx context.Context, outboxProcessor tm.OutboxProcessor, logger logger.Logger) {
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
