package ordering

import (
	"context"
	"fmt"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/monolith"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/docs"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/infrastructure/client/grpc"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/infrastructure/logging"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/infrastructure/persistence/gorm"
	grpcv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/interface/grpc/v1"
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

	application := logging.NewLogApplicationAccess(
		application.NewOrderApplication(
			orderRepository,
			customerClient,
			invoiceClient,
			notificationClient,
			paymentClient,
			shoppingClient),
		mono.Logger(),
	)

	if err := grpcv1.RegisterServer(ctx, application, mono.RPC().GRPCServer()); err != nil {
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
