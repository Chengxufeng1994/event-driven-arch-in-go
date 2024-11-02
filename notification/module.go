package notification

import (
	"context"
	"fmt"

	customerv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/customer/api/customer/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/broker/nats"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/monolith"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/application"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/infrastructure/client/grpc"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/infrastructure/logging"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/infrastructure/persistence/gorm"
	v1 "github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/interface/grpc/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/interface/handler"
)

type Module struct{}

var _ monolith.Module = Module{}

func NewModule() *Module {
	return &Module{}
}

func (m Module) PrepareRun(ctx context.Context, mono monolith.Monolith) error {
	// setup Driven adapters
	endpoint := fmt.Sprintf("%s:%d", mono.Config().Server.GPPC.Host, mono.Config().Server.GPPC.Port)
	conn, err := grpc.Dial(ctx, endpoint)
	if err != nil {
		return err
	}
	reg := registry.New()
	if err := customerv1.Registrations(reg); err != nil {
		return err
	}
	eventStream := am.NewEventStream(reg, nats.NewStream(mono.Config().Infrastructure.Nats.Stream, mono.JetStream()))
	customer := gorm.NewGormCustomerCacheRepository(mono.Database(), grpc.NewCustomerClient(conn))

	// setup application
	app := logging.NewLogApplicationAccess(application.New(customer), mono.Logger())
	customerHandler := logging.NewLogEventHandlerAccess[ddd.Event](
		application.NewCustomerIntegrationEventHandler(customer),
		"Customer",
		mono.Logger())
	orderHandler := logging.NewLogEventHandlerAccess[ddd.Event](
		application.NewOrderIntegrationEventHandler(app),
		"Order",
		mono.Logger())

	// setup Driver adapters
	if err := v1.RegisterServer(ctx, app, mono.RPC().GRPCServer()); err != nil {
		return err
	}

	if err := handler.RegisterCustomerIntegrationEventHandlers(customerHandler, eventStream); err != nil {
		return err
	}

	if err := handler.RegisterOrderIntegrationEventHandlers(orderHandler, eventStream); err != nil {
		return err
	}

	return nil
}

func (m Module) Name() string {
	return "notification"
}
