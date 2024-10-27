package customer

import (
	"context"
	"fmt"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/docs"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/application"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/infrastructure/logging"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/infrastructure/persistence/gorm"
	grpcv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/interface/grpc/v1"
	restv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/interface/http/rest/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/monolith"
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
	customerRepository := gorm.NewGormCustomerRepository(db)

	logApplication := logging.NewLogApplicationAccess(
		application.NewCustomerApplication(customerRepository, domainEventDispatcher),
		mono.Logger(),
	)
	// TODO: new domain event handlers, register domain events...

	if err := grpcv1.RegisterServer(ctx, logApplication, mono.RPC().GRPCServer()); err != nil {
		return err
	}

	if err := restv1.RegisterGateway(ctx, mono.Gin(), endpoint); err != nil {
		return err
	}

	return docs.RegisterSwagger(mono.Gin())
}

func (m *Module) Name() string {
	return "customer"
}
