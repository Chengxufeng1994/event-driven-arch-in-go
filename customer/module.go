package customer

import (
	"context"
	"fmt"

	customerv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/customer/api/customer/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/docs"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/application"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/application/handler"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/infrastructure/logging"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/infrastructure/persistence/gorm"
	grpcv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/interface/grpc/v1"
	restv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/interface/rest/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/broker/nats"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/monolith"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"
)

type Module struct{}

var _ monolith.Module = (*Module)(nil)

func NewModule() *Module { return &Module{} }

func (m *Module) PrepareRun(ctx context.Context, mono monolith.Monolith) error {
	// setup Driver adapters
	endpoint := fmt.Sprintf("%s:%d", mono.Config().Server.GPPC.Host, mono.Config().Server.GPPC.Port)
	reg := registry.New()
	if err := customerv1.Registrations(reg); err != nil {
		return err
	}
	stream := nats.NewStream(mono.Config().Infrastructure.Nats.Stream, mono.JetStream(), mono.Logger())
	eventStream := am.NewEventStream(reg, stream)
	commandStream := am.NewCommandStream(reg, stream)
	domainEventDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()
	customers := gorm.NewGormCustomerRepository(mono.Database())

	// setup application
	app := logging.NewLogApplicationAccess(
		application.NewCustomerApplication(customers, domainEventDispatcher), mono.Logger(),
	)
	domainEventHandler := logging.NewLogEventHandlerAccess(
		handler.NewDomainEventHandler(eventStream),
		"DomainEvents", mono.Logger(),
	)
	commandHandler := logging.NewLogCommandHandlerAccess(
		handler.NewCommandHandler(app),
		"Commands", mono.Logger(),
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
	handler.RegisterDomainEventHandler(domainEventHandler, domainEventDispatcher)
	if err := handler.RegisterCommandHandlers(commandStream, commandHandler); err != nil {
		return err
	}

	return nil
}

func (m *Module) Name() string {
	return "customer"
}
