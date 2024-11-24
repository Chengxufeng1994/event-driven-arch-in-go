package basket

import (
	"context"

	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"

	basketv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/basket/api/basket/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/docs"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/handler"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/repository"
	infragrpc "github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/infrastructure/client/grpc"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/infrastructure/constants"
	persistencegorm "github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/infrastructure/persistence/gorm"
	grpcv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/interface/grpc/v1"
	restv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/interface/rest/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/amotel"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/amprom"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/broker/nats"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/di"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/es"
	evenstoregorm "github.com/Chengxufeng1994/event-driven-arch-in-go/internal/eventstore/gorm"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/logger"
	outboxstoregorm "github.com/Chengxufeng1994/event-driven-arch-in-go/internal/outboxstore/gorm"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"
	snapshotstoregorm "github.com/Chengxufeng1994/event-driven-arch-in-go/internal/snapshotstore/gorm"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/system"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/tm"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Module struct{}

var _ system.Module = (*Module)(nil)

func NewModule() *Module { return &Module{} }

func Root(ctx context.Context, svc system.Service) error {
	container := di.New()

	// setup Driven adapters
	container.AddSingleton(constants.RegistryKey, func(c di.Container) (any, error) {
		reg := registry.New()
		if err := domain.Registrations(reg); err != nil {
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
	stream := nats.NewStream(svc.Config().Infrastructure.Nats.Stream, svc.JetStream(), svc.Logger())
	container.AddSingleton(constants.DomainDispatcherKey, func(c di.Container) (any, error) {
		return ddd.NewEventDispatcher[ddd.Event](), nil
	})
	container.AddScoped(constants.DatabaseTransactionKey, func(c di.Container) (any, error) {
		return svc.Database().Begin(), nil
	})
	sentCounter := amprom.SentMessagesCounter(constants.ServiceName)
	container.AddScoped(constants.MessagePublisherKey, func(c di.Container) (any, error) {
		tx := c.Get(constants.DatabaseTransactionKey).(*gorm.DB)
		outboxStore := outboxstoregorm.NewOutboxStore(constants.OutboxTableName, tx)
		return am.NewMessagePublisher(
			stream,
			amotel.OtelMessageContextInjector(),
			sentCounter,
			tm.OutboxPublisher(outboxStore),
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
		tx := c.Get(constants.DatabaseTransactionKey).(*gorm.DB)
		return outboxstoregorm.NewInboxStore(constants.InboxTableName, tx), nil
	})
	container.AddScoped(constants.BasketsRepoKey, func(c di.Container) (any, error) {
		tx := c.Get(constants.DatabaseTransactionKey).(*gorm.DB)
		reg := c.Get(constants.RegistryKey).(registry.Registry)
		store := es.AggregateStoreWithMiddleware(
			evenstoregorm.NewEventStore(constants.EventsTableName, tx, reg),
			snapshotstoregorm.NewSnapshotStore(constants.SnapshotsTableName, tx, reg),
		)
		return es.NewAggregateRepository[*aggregate.Basket](
			aggregate.BasketAggregate,
			reg,
			store,
		), nil
	})
	container.AddScoped(constants.StoresRepoKey, func(c di.Container) (any, error) {
		return persistencegorm.NewGormStoreCacheRepository(
				c.Get(constants.DatabaseTransactionKey).(*gorm.DB),
				infragrpc.NewGrpcStoreRepository(svc.Config().Server.GRPC.Service(constants.StoresServiceName))),
			nil
	})
	container.AddScoped(constants.ProductsRepoKey, func(c di.Container) (any, error) {
		return persistencegorm.NewGormProductCacheRepository(
				c.Get(constants.DatabaseTransactionKey).(*gorm.DB),
				infragrpc.NewGrpcProductRepository(svc.Config().Server.GRPC.Service(constants.StoresServiceName))),
			nil
	})
	// Prometheus counters
	basketsStarted := promauto.NewCounter(prometheus.CounterOpts{
		Name: constants.BasketsStartedCount,
	})
	basketsCheckedOut := promauto.NewCounter(prometheus.CounterOpts{
		Name: constants.BasketsCheckedOutCount,
	})
	basketsCanceled := promauto.NewCounter(prometheus.CounterOpts{
		Name: constants.BaksetsCanceledCount,
	})

	// setup application
	container.AddScoped(constants.ApplicationKey, func(c di.Container) (any, error) {
		return application.NewInstrumentedApp(
			application.New(
				c.Get(constants.BasketsRepoKey).(repository.BasketRepository),
				c.Get(constants.StoresRepoKey).(repository.StoreCacheRepository),
				c.Get(constants.ProductsRepoKey).(repository.ProductCacheRepository),
				c.Get(constants.DomainDispatcherKey).(ddd.EventDispatcher[ddd.Event]),
			),
			basketsStarted, basketsCheckedOut, basketsCanceled,
		), nil
	})
	container.AddScoped(constants.DomainEventHandlersKey, func(c di.Container) (any, error) {
		return handler.NewDomainEventHandlers(c.Get(constants.EventPublisherKey).(am.EventPublisher)), nil
	})
	container.AddScoped(constants.IntegrationEventHandlersKey, func(c di.Container) (any, error) {
		return handler.NewIntegrationEventHandlers(
				c.Get(constants.RegistryKey).(registry.Registry),
				c.Get(constants.StoresRepoKey).(repository.StoreCacheRepository),
				c.Get(constants.ProductsRepoKey).(repository.ProductCacheRepository),
				tm.InboxHandler(c.Get(constants.InboxStoreKey).(tm.InboxStore)),
			),
			nil
	})
	outboxProcessor := tm.NewOutboxProcessor(
		stream,
		outboxstoregorm.NewOutboxStore(constants.OutboxTableName, svc.Database()),
	)

	// setup Driver adapters
	if err := grpcv1.RegisterServerTx(container, svc.RPC().GRPCServer()); err != nil {
		return err
	}
	if err := restv1.RegisterGateway(ctx, svc.Gin(), svc.Config().Server.GRPC.Address()); err != nil {
		return err
	}
	if err := docs.RegisterSwagger(svc.Gin()); err != nil {
		return err
	}
	handler.RegisterDomainEventHandlersTx(container)
	if err := handler.RegisterIntegrationEventHandlersTx(container); err != nil {
		return err
	}

	go startOutboxProcessor(ctx, outboxProcessor, svc.Logger())

	return nil
}

func (m *Module) Startup(ctx context.Context, svc system.Service) error {
	return Root(ctx, svc)
}

func (m *Module) Name() string {
	return "baskets"
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
		logger.WithError(err).Error("baskets outbox processor encountered an error")
	}
}
