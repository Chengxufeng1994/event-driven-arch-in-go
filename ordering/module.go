package ordering

import (
	"context"
	"fmt"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/es"
	evenstoregorm "github.com/Chengxufeng1994/event-driven-arch-in-go/internal/eventstore/gorm"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/monolith"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry/serdes"
	snapshotstoregorm "github.com/Chengxufeng1994/event-driven-arch-in-go/internal/snapshotstore/gorm"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/docs"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/event"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/infrastructure/client/grpc"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/infrastructure/logging"
	grpcv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/interface/grpc/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/interface/handler"
	v1 "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/interface/http/rest/v1"
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
	conn, err := grpc.Dial(ctx, endpoint)
	if err != nil {
		return err
	}

	domainEventDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()
	aggregateStore := es.AggregateStoreWithMiddleware(
		evenstoregorm.NewGormEventStore("ordering.events", mono.Database(), reg),
		es.NewEventPublisher(domainEventDispatcher),
		snapshotstoregorm.NewGormSnapshotStore("ordering.snapshots", mono.Database(), reg),
	)
	orderRepository := es.NewAggregateRepository[*aggregate.Order](aggregate.OrderAggregate, reg, aggregateStore)
	customerClient := grpc.NewGrpcCustomerClient(conn)
	invoiceClient := grpc.NewGrpcInvoiceClient(conn)
	notificationClient := grpc.NewGrpcNotificationClient(conn)
	paymentClient := grpc.NewGrpcPaymentClient(conn)
	shoppingClient := grpc.NewGrpcShoppingClient(conn)

	// setup application
	logApplication := logging.NewLogApplicationAccess(
		application.NewOrderApplication(
			orderRepository,
			customerClient,
			paymentClient,
			shoppingClient,
		),
		mono.Logger(),
	)
	// setup application handlers
	notificationDomainEventHandler := logging.NewLogNotificationEventHandlerAccess[ddd.AggregateEvent](
		application.NewNotificationDomainEventHandler(notificationClient),
		"Notification",
		mono.Logger())
	invoiceDomainEventHandler := logging.NewLogInvoiceEventHandlerAccess[ddd.AggregateEvent](
		application.NewInvoiceDomainEventHandler(invoiceClient),
		"Invoice",
		mono.Logger(),
	)

	// setup Driver adapters
	if err := grpcv1.RegisterServer(ctx, logApplication, mono.RPC().GRPCServer()); err != nil {
		return err
	}

	if v1.RegisterGateway(ctx, mono.Gin(), endpoint) != nil {
		return err
	}

	if err := docs.RegisterSwagger(mono.Gin()); err != nil {
		return err
	}

	handler.RegisterNotificationHandler(notificationDomainEventHandler, domainEventDispatcher)
	handler.RegisterInvoiceHandler(invoiceDomainEventHandler, domainEventDispatcher)

	return nil
}

func (m *Module) Name() string {
	return "order"
}

func registrations(reg registry.Registry) (err error) {
	serde := serdes.NewJsonSerde(reg)

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
