package notification

import (
	"context"

	customerv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/customer/api/customer/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/amotel"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/amprom"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/broker/nats"
	outboxstoregorm "github.com/Chengxufeng1994/event-driven-arch-in-go/internal/outboxstore/gorm"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/system"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/tm"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/application"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/application/handler"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/infrastructure/client/grpc"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/infrastructure/constants"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/infrastructure/persistence/gorm"
	v1 "github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/interface/grpc/v1"
	orderv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/api/order/v1"
)

type Module struct{}

var _ system.Module = Module{}

func NewModule() *Module { return &Module{} }

func Root(ctx context.Context, svc system.Service) error {
	// setup Driven adapters
	reg := registry.New()
	if err := customerv1.Registrations(reg); err != nil {
		return err
	}
	if err := orderv1.Registrations(reg); err != nil {
		return err
	}
	stream := nats.NewStream(svc.Config().Infrastructure.Nats.Stream, svc.JetStream(), svc.Logger())
	inboxStore := outboxstoregorm.NewInboxStore(constants.InboxTableName, svc.Database())
	messageSubscriber := am.NewMessageSubscriber(stream,
		amotel.OtelMessageContextExtractor(),
		amprom.ReceivedMessagesCounter(constants.ServiceName),
	)
	customer := gorm.NewGormCustomerCacheRepository(
		svc.Database(),
		grpc.NewGrpcCustomerRepository(svc.Config().Server.GRPC.Service(constants.CustomersServiceName)))

	// setup application
	app := application.New(customer)
	integrationEventHandlers := handler.NewIntegrationEventHandler(
		reg, app, customer,
		tm.InboxHandler(inboxStore),
	)

	// setup Driver adapters
	if err := v1.RegisterServer(ctx, app, svc.RPC().GRPCServer()); err != nil {
		return err
	}
	if err := handler.RegisterIntegrationEventHandlers(messageSubscriber, integrationEventHandlers); err != nil {
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
