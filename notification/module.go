package notification

import (
	"context"
	"fmt"

	customerv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/customer/api/customer/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/broker/nats"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/system"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/application"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/application/handler"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/infrastructure/client/grpc"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/infrastructure/logging"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/infrastructure/persistence/gorm"
	v1 "github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/interface/grpc/v1"
	orderv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/api/order/v1"
)

type Module struct{}

var _ system.Module = Module{}

func NewModule() *Module { return &Module{} }

func Root(ctx context.Context, svc system.Service) error {
	// setup Driven adapters
	endpoint := fmt.Sprintf("%s:%d", svc.Config().Server.GPPC.Host, svc.Config().Server.GPPC.Port)
	reg := registry.New()
	if err := customerv1.Registrations(reg); err != nil {
		return err
	}
	if err := orderv1.Registrations(reg); err != nil {
		return err
	}
	eventStream := am.NewEventStream(reg, nats.NewStream(svc.Config().Infrastructure.Nats.Stream, svc.JetStream(), svc.Logger()))
	conn, err := grpc.Dial(ctx, endpoint)
	if err != nil {
		return err
	}
	customer := gorm.NewGormCustomerCacheRepository(svc.Database(), grpc.NewCustomerClient(conn))

	// setup application
	app := logging.NewLogApplicationAccess(application.New(customer), svc.Logger())
	integrationEventHandler := logging.NewLogEventHandlerAccess(
		handler.NewIntegrationEventHandler(app, customer),
		"IntegrationEvents", svc.Logger(),
	)

	// setup Driver adapters
	if err := v1.RegisterServer(ctx, app, svc.RPC().GRPCServer()); err != nil {
		return err
	}
	if err := handler.RegisterIntegrationEventHandlers(eventStream, integrationEventHandler); err != nil {
		return err
	}

	return nil
}

func (m Module) Startup(ctx context.Context, svc system.Service) error {
	return Root(ctx, svc)
}

func (m Module) Name() string {
	return "notification"
}
