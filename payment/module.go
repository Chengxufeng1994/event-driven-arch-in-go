package payment

import (
	"context"
	"fmt"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/broker/nats"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/monolith"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"
	orderv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/api/order/v1"
	paymentv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/payment/api/payment/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/docs"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/application"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/application/handler"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/infrastructure/logging"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/infrastructure/persistence/gorm"
	v1 "github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/interface/grpc/v1"
	restv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/interface/rest/v1"
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
	if err := orderv1.Registrations(reg); err != nil {
		return err
	}
	if err := paymentv1.Registrations(reg); err != nil {
		return err
	}
	stream := nats.NewStream(mono.Config().Infrastructure.Nats.Stream, mono.JetStream(), mono.Logger())
	eventStream := am.NewEventStream(reg, stream)
	commandStream := am.NewCommandStream(reg, stream)
	domainDispatcher := ddd.NewEventDispatcher[ddd.Event]()
	invoiceRepository := gorm.NewGormInvoiceRepository(mono.Database())
	paymentRepository := gorm.NewGormPaymentRepository(mono.Database())

	// setup app
	app := logging.NewLogApplicationAccess(
		application.NewPaymentApplication(invoiceRepository, paymentRepository, domainDispatcher),
		mono.Logger(),
	)
	domainEventHandler := logging.NewLogEventHandlerAccess[ddd.Event](
		handler.NewDomainEventHandler(eventStream),
		"DomainEvents", mono.Logger(),
	)
	integrationEventHandler := logging.NewLogEventHandlerAccess[ddd.Event](
		handler.NewIntegrationEventHandlers(app),
		"IntegrationEvents", mono.Logger(),
	)
	commandHandler := logging.NewLogCommandHandlerAccess(
		handler.NewCommandHandler(app),
		"Commands", mono.Logger(),
	)

	// setup Driver adapters
	if err := v1.RegisterServer(ctx, app, mono.RPC().GRPCServer()); err != nil {
		return err
	}
	if err := restv1.RegisterGateway(ctx, mono.Gin(), endpoint); err != nil {
		return err
	}
	if err := docs.RegisterSwagger(mono.Gin()); err != nil {
		return err
	}
	if err := handler.RegisterIntegrationEventHandler(integrationEventHandler, eventStream); err != nil {
		return err
	}
	handler.RegisterDomainEventHandler(domainEventHandler, domainDispatcher)
	if err := handler.RegisterCommandHandler(commandStream, commandHandler); err != nil {
		return err
	}

	return nil
}

func (m *Module) Name() string {
	return "payment"
}
