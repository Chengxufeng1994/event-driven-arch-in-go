package basket

import (
	"context"
	"fmt"

	basketv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/basket/api/basket/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/docs"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/handler"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/entity"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/event"
	infragrpc "github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/infrastructure/client/grpc"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/infrastructure/logging"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/infrastructure/persistence/gorm"
	grpcv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/interface/grpc/v1"
	restv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/interface/rest/v1"
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
)

type Module struct{}

var _ monolith.Module = (*Module)(nil)

func NewModule() *Module { return &Module{} }

func (m *Module) PrepareRun(ctx context.Context, mono monolith.Monolith) error {
	// setup Driven adapters
	endpoint := fmt.Sprintf("%s:%d", mono.Config().Server.GPPC.Host, mono.Config().Server.GPPC.Port)
	reg := registry.New()
	if err := registrations(reg); err != nil {
		return err
	}
	if err := basketv1.Registrations(reg); err != nil {
		return err
	}
	if err := storev1.Registrations(reg); err != nil {
		return err
	}
	stream := nats.NewStream(mono.Config().Infrastructure.Nats.Stream, mono.JetStream(), mono.Logger())
	eventStream := am.NewEventStream(reg, stream)
	conn, err := infragrpc.Dial(ctx, endpoint)
	if err != nil {
		return err
	}
	domainEventDispatcher := ddd.NewEventDispatcher[ddd.Event]()
	aggregateStore := es.AggregateStoreWithMiddleware(
		evenstoregorm.NewGormEventStore("baskets.events", mono.Database(), reg),
		snapshotstoregorm.NewGormSnapshotStore("baskets.snapshots", mono.Database(), reg),
	)
	basketRepository := es.NewAggregateRepository[*aggregate.Basket](aggregate.BasketAggregate, reg, aggregateStore)
	storeRepository := gorm.NewGormStoreCacheRepository(mono.Database(), infragrpc.NewGrpcStoreClient(conn))
	productRepository := gorm.NewGormProductCacheRepository(mono.Database(), infragrpc.NewGrpcProductClient(conn))

	// setup application
	app := logging.NewLogApplicationAccess(
		application.NewBasketApplication(basketRepository, productRepository, storeRepository, domainEventDispatcher),
		mono.Logger())
	domainEventHandler := logging.NewLogEventHandlerAccess[ddd.Event](
		handler.NewDomainEventHandler(eventStream),
		"DomainEvents", mono.Logger(),
	)
	integrationEventHandler := logging.NewLogEventHandlerAccess[ddd.Event](
		handler.NewIntegrationEventHandler(storeRepository, productRepository),
		"IntegrationEvents", mono.Logger(),
	)

	// setup Driver adapters
	if err := grpcv1.RegisterServer(ctx, app, mono.RPC().GRPCServer()); err != nil {
		return err
	}
	if err := restv1.RegisterGateway(ctx, mono.Gin(), endpoint); err != nil {
		return err
	}
	if err := docs.RegisterSwagger(mono.Gin()); err != nil {
		return err
	}

	handler.RegisterDomainEventHandlers(domainEventDispatcher, domainEventHandler)
	return handler.RegisterIntegrationEventHandlers(eventStream, integrationEventHandler)
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
