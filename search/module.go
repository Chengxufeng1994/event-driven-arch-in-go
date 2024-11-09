package search

import (
	"context"
	"fmt"

	customerv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/customer/api/customer/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/broker/nats"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/monolith"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"
	orderv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/api/order/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/docs"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/application"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/application/handler"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/infrastructure/client/grpc"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/infrastructure/logging"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/infrastructure/persistence/gorm"
	v1 "github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/interface/grpc/v1"
	restv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/interface/rest/v1"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
)

type Module struct{}

var _ monolith.Module = (*Module)(nil)

func NewModule() *Module {
	return &Module{}
}

func (m *Module) PrepareRun(ctx context.Context, mono monolith.Monolith) error {
	// setup Driven adapters
	endpoint := fmt.Sprintf("%s:%d", mono.Config().Server.GPPC.Host, mono.Config().Server.GPPC.Port)
	reg := registry.New()
	if err := orderv1.Registrations(reg); err != nil {
		return err
	}
	if err := customerv1.Registrations(reg); err != nil {
		return err
	}
	if err := storev1.Registrations(reg); err != nil {
		return err
	}
	stream := nats.NewStream(mono.Config().Infrastructure.Nats.Stream, mono.JetStream(), mono.Logger())
	eventStream := am.NewEventStream(reg, stream)
	conn, err := grpc.Dial(ctx, endpoint)
	if err != nil {
		return err
	}
	customers := gorm.NewGormCustomerCacheRepository(mono.Database(), grpc.NewCustomerClient(conn))
	stores := gorm.NewGormStoreCacheRepository(mono.Database(), grpc.NewStoreClient(conn))
	products := gorm.NewGormProductCacheRepository(mono.Database(), grpc.NewProductClient(conn))
	orders := gorm.NewGormOrderRepository(mono.Database())

	// setup application
	app := logging.NewLogApplicationAccess(
		application.New(orders),
		mono.Logger(),
	)
	integrationEventHandlers := logging.NewLogEventHandlerAccess[ddd.Event](
		handler.NewIntegrationEventHandlers(orders, customers, products, stores),
		"IntegrationEvents", mono.Logger(),
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
	if err := handler.RegisterIntegrationEventHandlers(eventStream, integrationEventHandlers); err != nil {
		return err
	}

	return nil
}

func (m *Module) Name() string {
	return "search"
}
