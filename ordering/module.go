package ordering

import (
	"context"
	"fmt"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/monolith"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/docs"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/infrastructure/client/grpc"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/infrastructure/logging"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/infrastructure/persistence/gorm"
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
	endpoint := fmt.Sprintf("%s:%d", mono.Config().Server.GPPC.Host, mono.Config().Server.GPPC.Port)
	db := mono.Database()

	domainEventDispatcher := ddd.NewEventDispatcher()
	orderRepository := gorm.NewGormOrderRepository(db)
	conn, err := grpc.Dial(ctx, endpoint)
	if err != nil {
		return err
	}
	customerClient := grpc.NewGrpcCustomerClient(conn)
	invoiceClient := grpc.NewGrpcInvoiceClient(conn)
	notificationClient := grpc.NewGrpcNotificationClient(conn)
	paymentClient := grpc.NewGrpcPaymentClient(conn)
	shoppingClient := grpc.NewGrpcShoppingClient(conn)

	logApplication := logging.NewLogApplicationAccess(
		application.NewOrderApplication(
			orderRepository,
			customerClient,
			paymentClient,
			shoppingClient,
			domainEventDispatcher,
		),
		mono.Logger(),
	)

	invoiceDomainEventHandler := logging.NewLogInvoiceDomainEventHandlerAccess(
		application.NewInvoiceDomainEventHandler(invoiceClient),
		mono.Logger(),
	)
	notificationDomainEventHandler := logging.NewLogNotificationDomainEventHandlerAccess(
		application.NewNotificationDomainEventHandler(notificationClient),
		mono.Logger())

	handler.RegisterInvoiceDomainEventHandlers(invoiceDomainEventHandler, domainEventDispatcher)
	handler.RegisterNotificationDomainEventHandlers(notificationDomainEventHandler, domainEventDispatcher)

	if err := grpcv1.RegisterServer(ctx, logApplication, mono.RPC().GRPCServer()); err != nil {
		return err
	}

	if v1.RegisterGateway(ctx, mono.Gin(), endpoint) != nil {
		return err
	}

	return docs.RegisterSwagger(mono.Gin())
}

func (m *Module) Name() string {
	return "order"
}
