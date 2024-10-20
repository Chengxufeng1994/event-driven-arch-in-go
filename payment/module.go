package payment

import (
	"context"
	"fmt"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/monolith"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/docs"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/application"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/infrastructure/client/grpc"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/infrastructure/logging"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/infrastructure/persistence/gorm"
	v1 "github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/interface/grpc/v1"
	restv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/interface/http/rest/v1"
)

type Module struct{}

var _ monolith.Module = (*Module)(nil)

func NewModule() *Module {
	return &Module{}
}

func (m *Module) PrepareRun(ctx context.Context, mono monolith.Monolith) error {
	endpoint := fmt.Sprintf("%s:%d", mono.Config().Server.GPPC.Host, mono.Config().Server.GPPC.Port)
	db := mono.Database()
	conn, err := grpc.Dial(ctx, endpoint)
	if err != nil {
		return err
	}

	invoiceRepository := gorm.NewGormInvoiceRepository(db)
	paymentRepository := gorm.NewGormPaymentRepository(db)
	orderClient := grpc.NewGrpcOrderClient(conn)

	application := logging.NewLogApplicationAccess(
		application.NewPaymentApplication(invoiceRepository, paymentRepository, orderClient),
		mono.Logger(),
	)

	if err := v1.RegisterServer(ctx, application, mono.RPC().GRPCServer()); err != nil {
		return err
	}

	if err := restv1.RegisterGateway(ctx, mono.Gin(), endpoint); err != nil {
		return err
	}

	return docs.RegisterSwagger(mono.Gin())
}

func (m *Module) Name() string {
	return "payment"
}
