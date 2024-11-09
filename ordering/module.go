package ordering

import (
	"context"
	"fmt"

	basketv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/basket/api/basket/v1"
	depotv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/depot/api/depot/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/broker/nats"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/es"
	evenstoregorm "github.com/Chengxufeng1994/event-driven-arch-in-go/internal/eventstore/gorm"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/monolith"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry/serdes"
	snapshotstoregorm "github.com/Chengxufeng1994/event-driven-arch-in-go/internal/snapshotstore/gorm"
	orderv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/api/order/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/docs"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application/handler"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/event"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/infrastructure/client/grpc"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/infrastructure/logging"
	grpcv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/interface/grpc/v1"
	v1 "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/interface/rest/v1"
)

type Module struct{}

var _ monolith.Module = (*Module)(nil)

func NewModule() *Module {
	return &Module{}
}

func (m *Module) PrepareRun(ctx context.Context, mono monolith.Monolith) error {
	// setup Driver adapters
	endpoint := fmt.Sprintf("%s:%d", mono.Config().Server.GPPC.Host, mono.Config().Server.GPPC.Port)
	reg := registry.New()
	if err := registrations(reg); err != nil {
		return err
	}
	if err := orderv1.Registrations(reg); err != nil {
		return err
	}
	if err := basketv1.Registrations(reg); err != nil {
		return err
	}
	if err := depotv1.Registrations(reg); err != nil {
		return err
	}
	stream := nats.NewStream(mono.Config().Infrastructure.Nats.Stream, mono.JetStream(), mono.Logger())
	eventStream := am.NewEventStream(reg, stream)
	commandStream := am.NewCommandStream(reg, stream)
	conn, err := grpc.Dial(ctx, endpoint)
	if err != nil {
		return err
	}
	aggregateStore := es.AggregateStoreWithMiddleware(
		evenstoregorm.NewGormEventStore("ordering.events", mono.Database(), reg),
		snapshotstoregorm.NewGormSnapshotStore("ordering.snapshots", mono.Database(), reg),
	)
	domainEventDispatcher := ddd.NewEventDispatcher[ddd.Event]()
	orderRepository := es.NewAggregateRepository[*aggregate.Order](aggregate.OrderAggregate, reg, aggregateStore)
	shoppingClient := grpc.NewGrpcShoppingClient(conn)

	// setup application
	app := logging.NewLogApplicationAccess(
		application.NewOrderApplication(orderRepository, shoppingClient, domainEventDispatcher),
		mono.Logger(),
	)
	// setup application handlers
	domainEventHandlers := logging.NewLogEventHandlerAccess[ddd.Event](
		handler.NewDomainEventHandler(eventStream),
		"DomainEvents", mono.Logger(),
	)
	integrationEventHandlers := logging.NewLogEventHandlerAccess[ddd.Event](
		handler.NewIntegrationEventHandlers(app),
		"IntegrationEvents", mono.Logger(),
	)
	commandHandlers := logging.NewLogCommandHandlerAccess[ddd.Command](
		handler.NewCommandHandlers(app),
		"Commands", mono.Logger(),
	)

	// setup Driver adapters
	if err := grpcv1.RegisterServer(ctx, app, mono.RPC().GRPCServer()); err != nil {
		return err
	}
	if v1.RegisterGateway(ctx, mono.Gin(), endpoint) != nil {
		return err
	}
	if err := docs.RegisterSwagger(mono.Gin()); err != nil {
		return err
	}
	handler.RegisterDomainEventHandlers(domainEventDispatcher, domainEventHandlers)
	if err := handler.RegisterIntegrationEventHandler(integrationEventHandlers, eventStream); err != nil {
		return err
	}
	if err := handler.RegisterCommandHandlers(commandStream, commandHandlers); err != nil {
		return err
	}

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
		order.AggregateBase = es.NewAggregateBase("", aggregate.OrderAggregate)
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
