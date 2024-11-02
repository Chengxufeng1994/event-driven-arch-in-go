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
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/infrastructure/client/grpc"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/infrastructure/logging"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/infrastructure/persistence/gorm"
	v1 "github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/interface/grpc/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/interface/handler"
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
	// serialize, deserialize register
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
	eventStream := am.NewEventStream(reg, nats.NewStream(mono.Config().Infrastructure.Nats.Stream, mono.JetStream()))
	conn, err := grpc.Dial(ctx, endpoint)
	if err != nil {
		return err
	}
	customer := gorm.NewGormCustomerCacheRepository(mono.Database(), grpc.NewCustomerClient(conn))
	store := gorm.NewGormStoreCacheRepository(mono.Database(), grpc.NewStoreClient(conn))
	product := gorm.NewGormProductCacheRepository(mono.Database(), grpc.NewProductClient(conn))
	order := gorm.NewGormOrderRepository(mono.Database())

	// setup application
	app := logging.NewLogApplicationAccess(
		application.New(order),
		mono.Logger(),
	)
	orderHandler := logging.NewLogEventHandlerAccess[ddd.Event](
		application.NewOrderHandlers(order, customer, store, product),
		"Order",
		mono.Logger(),
	)
	customerHandler := logging.NewLogEventHandlerAccess[ddd.Event](
		application.NewCustomerHandlers(customer),
		"Customer",
		mono.Logger(),
	)
	storeHandler := logging.NewLogEventHandlerAccess[ddd.Event](
		application.NewStoreHandlers(store),
		"Store",
		mono.Logger(),
	)
	productHandler := logging.NewLogEventHandlerAccess[ddd.Event](
		application.NewProductHandlers(product),
		"Product",
		mono.Logger(),
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
	if err := handler.RegisterCustomerHandlers(customerHandler, eventStream); err != nil {
		return err
	}
	if err := handler.RegisterOrderHandlers(orderHandler, eventStream); err != nil {
		return err
	}
	if err := handler.RegisterStoreHandlers(storeHandler, eventStream); err != nil {
		return err
	}
	if err := handler.RegisterProductHandlers(productHandler, eventStream); err != nil {
		return err
	}

	return nil
}

func (m *Module) Name() string {
	return "search"
}
