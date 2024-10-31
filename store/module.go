package store

import (
	"context"
	"fmt"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/broker/nats"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/es"
	evenstoregorm "github.com/Chengxufeng1994/event-driven-arch-in-go/internal/eventstore/gorm"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/monolith"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry/serdes"
	snapshotstoregorm "github.com/Chengxufeng1994/event-driven-arch-in-go/internal/snapshotstore/gorm"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/docs"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/application"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/event"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/infrastructure/logging"
	persistencegorm "github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/infrastructure/persistence/gorm"
	v1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/interface/grpc/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/interface/handler"
	restv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/interface/rest/v1"
)

type Module struct{}

var _ monolith.Module = (*Module)(nil)

func NewModule() *Module {
	return &Module{}
}

func (m *Module) PrepareRun(ctx context.Context, mono monolith.Monolith) error {
	// setup Driven adapters
	// serialize, deserialize register
	endpoint := fmt.Sprintf("%s:%d", mono.Config().Server.GPPC.Host, mono.Config().Server.GPPC.Port)
	reg := registry.New()
	if err := registrations(reg); err != nil {
		return err
	}
	if err := storev1.Registrations(reg); err != nil {
		return err
	}
	eventStream := am.NewEventStream(reg, nats.NewStream(mono.Config().Infrastructure.Nats.Stream, mono.JetStream()))
	domainEventDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()
	aggregateStore := es.AggregateStoreWithMiddleware(
		evenstoregorm.NewGormEventStore("stores.events", mono.Database(), reg),
		es.NewEventPublisher(domainEventDispatcher),
		snapshotstoregorm.NewGormSnapshotStore("stores.snapshots", mono.Database(), reg),
	)
	storeRepository := es.NewAggregateRepository[*aggregate.Store](aggregate.StoreAggregate, reg, aggregateStore)
	productRepository := es.NewAggregateRepository[*aggregate.Product](aggregate.ProductAggregate, reg, aggregateStore)
	mallRepository := persistencegorm.NewGormMallRepository(mono.Database())
	catalogRepository := persistencegorm.NewGormCatalogRepository(mono.Database())

	// setup application
	app := logging.NewLogApplicationAccess(
		application.NewStoreApplication(
			storeRepository,
			productRepository,
			mallRepository,
			catalogRepository,
		),
		mono.Logger())

	catalogHandler := logging.NewLogHandlerAccess(
		application.NewCatalogDomainEventHandler(catalogRepository),
		"Catalog",
		mono.Logger(),
	)

	mallHandler := logging.NewLogHandlerAccess(
		application.NewMallDomainEventHandler(mallRepository),
		"Mall",
		mono.Logger(),
	)

	storeHandler := logging.NewLogHandlerAccess(
		application.NewStoreIntegrationEventHandler(eventStream),
		"Store",
		mono.Logger(),
	)

	// setup Driver adapters
	if err := v1.RegisterServer(app, mono.RPC().GRPCServer()); err != nil {
		return err
	}

	if err := restv1.RegisterGateway(ctx, mono.Gin(), endpoint); err != nil {
		return err
	}

	if err := docs.RegisterSwagger(mono.Gin()); err != nil {
		return err
	}

	handler.RegisterCatalogHandler(catalogHandler, domainEventDispatcher)
	handler.RegisterMallHandler(mallHandler, domainEventDispatcher)
	handler.RegisterStoreIntegrationHandler(storeHandler, domainEventDispatcher)

	return nil
}

func (m *Module) Name() string {
	return "store"
}

func registrations(reg registry.Registry) error {
	serde := serdes.NewJsonSerde(reg)

	// store
	if err := serde.Register(aggregate.Store{}, func(v any) error {
		store := v.(*aggregate.Store)
		store.AggregateBase = es.NewAggregateBase("", aggregate.StoreAggregate)
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
		product.AggregateBase = es.NewAggregateBase("", aggregate.ProductAggregate)
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
