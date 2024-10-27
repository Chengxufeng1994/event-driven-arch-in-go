package basket

import (
	"context"
	"fmt"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/docs"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/entity"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/event"
	infragrpc "github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/infrastructure/client/grpc"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/infrastructure/logging"
	grpcv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/interface/grpc/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/interface/handler"
	restv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/interface/http/rest/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/es"
	evenstoregorm "github.com/Chengxufeng1994/event-driven-arch-in-go/internal/eventstore/gorm"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/monolith"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry/serdes"
	snapshotstoregorm "github.com/Chengxufeng1994/event-driven-arch-in-go/internal/snapshotstore/gorm"
)

type Module struct{}

var _ monolith.Module = (*Module)(nil)

func NewModule() *Module { return &Module{} }

func (m *Module) PrepareRun(ctx context.Context, mono monolith.Monolith) error {
	// setup Driven adapters
	endpoint := fmt.Sprintf("%s:%d", mono.Config().Server.GPPC.Host, mono.Config().Server.GPPC.Port)
	conn, err := infragrpc.Dial(ctx, endpoint)
	if err != nil {
		return err
	}
	reg := registry.New()
	if err := registrations(reg); err != nil {
		return err
	}

	domainEventDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()
	aggregateStore := es.AggregateStoreWithMiddleware(
		evenstoregorm.NewGormEventStore("baskets.events", mono.Database(), reg),
		es.NewEventPublisher(domainEventDispatcher),
		snapshotstoregorm.NewGormSnapshotStore("baskets.snapshots", mono.Database(), reg),
	)

	basketRepository := es.NewAggregateRepository[*aggregate.Basket](aggregate.BasketAggregate, reg, aggregateStore)
	grpcOrderClient := infragrpc.NewGrpcOrderClient(conn)
	grpcProductClient := infragrpc.NewGrpcProductClient(conn)
	grpcStoreClient := infragrpc.NewGrpcStoreClient(conn)

	// setup application
	logApplication := logging.NewLogApplicationAccess(
		application.NewBasketApplication(basketRepository, grpcOrderClient, grpcProductClient, grpcStoreClient),
		mono.Logger())
	logDomainEventHandlers := logging.NewLogDomainEventHandlerAccess(
		application.NewOrderDomainEventHandler(grpcOrderClient),
		"Order",
		mono.Logger(),
	)

	// setup Driver adapters
	if err := grpcv1.RegisterServer(ctx, logApplication, mono.RPC().GRPCServer()); err != nil {
		return err
	}

	if err := restv1.RegisterGateway(ctx, mono.Gin(), endpoint); err != nil {
		return err
	}

	if err := docs.RegisterSwagger(mono.Gin()); err != nil {
		return err
	}

	handler.RegisterOrderDomainEventHandlers(logDomainEventHandlers, domainEventDispatcher)

	return nil
}

func (m *Module) Name() string {
	return "basket"
}

func registrations(reg registry.Registry) error {
	serde := serdes.NewJsonSerde(reg)
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
