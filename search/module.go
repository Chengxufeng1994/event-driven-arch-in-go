package search

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"gorm.io/gorm"

	customerv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/customer/api/customer/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/broker/nats"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/di"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/logger"
	outboxstoregorm "github.com/Chengxufeng1994/event-driven-arch-in-go/internal/outboxstore/gorm"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/system"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/tm"
	orderv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/api/order/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/docs"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/application"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/application/handler"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/application/port/out"
	infragrpc "github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/infrastructure/client/grpc"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/infrastructure/logging"
	persistencegorm "github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/infrastructure/persistence/gorm"
	v1 "github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/interface/grpc/v1"
	restv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/interface/rest/v1"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
)

type Module struct{}

var _ system.Module = (*Module)(nil)

func NewModule() *Module { return &Module{} }

func (m *Module) Startup(ctx context.Context, mono system.Service) error {
	container := di.New()
	// setup Driven adapters
	endpoint := fmt.Sprintf("%s:%d", mono.Config().Server.GPPC.Host, mono.Config().Server.GPPC.Port)
	container.AddSingleton("registry", func(c di.Container) (any, error) {
		reg := registry.New()
		if err := orderv1.Registrations(reg); err != nil {
			return nil, err
		}
		if err := customerv1.Registrations(reg); err != nil {
			return nil, err
		}
		if err := storev1.Registrations(reg); err != nil {
			return nil, err
		}
		return reg, nil
	})
	container.AddSingleton("logger", func(c di.Container) (any, error) {
		return mono.Logger(), nil
	})
	container.AddSingleton("stream", func(c di.Container) (any, error) {
		return nats.NewStream(mono.Config().Infrastructure.Nats.Stream, mono.JetStream(), mono.Logger()), nil
	})
	container.AddSingleton("db", func(c di.Container) (any, error) {
		return mono.Database(), nil
	})
	container.AddSingleton("conn", func(c di.Container) (any, error) {
		return infragrpc.Dial(ctx, endpoint)
	})
	container.AddScoped("tx", func(c di.Container) (any, error) {
		db := c.Get("db").(*gorm.DB)
		return db.Begin(), nil
	})
	container.AddScoped("inboxMiddleware", func(c di.Container) (any, error) {
		tx := c.Get("tx").(*gorm.DB)
		inboxStore := outboxstoregorm.NewInboxStore("search.inbox", tx)
		return tm.NewInboxHandlerMiddleware(inboxStore), nil
	})
	container.AddScoped("customers", func(c di.Container) (any, error) {
		return persistencegorm.NewGormCustomerCacheRepository(
			c.Get("tx").(*gorm.DB),
			infragrpc.NewCustomerClient(c.Get("conn").(*grpc.ClientConn)),
		), nil
	})
	container.AddScoped("stores", func(c di.Container) (any, error) {
		return persistencegorm.NewGormStoreCacheRepository(
			c.Get("tx").(*gorm.DB),
			infragrpc.NewStoreClient(c.Get("conn").(*grpc.ClientConn)),
		), nil
	})
	container.AddScoped("products", func(c di.Container) (any, error) {
		return persistencegorm.NewGormProductCacheRepository(
			c.Get("tx").(*gorm.DB),
			infragrpc.NewProductClient(c.Get("conn").(*grpc.ClientConn)),
		), nil
	})
	container.AddScoped("orders", func(c di.Container) (any, error) {
		return persistencegorm.NewGormOrderRepository(c.Get("tx").(*gorm.DB)), nil
	})

	// setup application
	container.AddScoped("app", func(c di.Container) (any, error) {
		return logging.NewLogApplicationAccess(
			application.New(c.Get("orders").(*persistencegorm.GormOrderRepository)),
			c.Get("logger").(logger.Logger),
		), nil
	})
	container.AddScoped("integrationEventHandlers", func(c di.Container) (any, error) {
		return logging.NewLogEventHandlerAccess[ddd.Event](
			handler.NewIntegrationEventHandlers(
				c.Get("orders").(out.OrderRepository),
				c.Get("customers").(out.CustomerCacheRepository),
				c.Get("products").(out.ProductCacheRepository),
				c.Get("stores").(out.StoreCacheRepository),
			),
			"IntegrationEvents", c.Get("logger").(logger.Logger),
		), nil
	})

	// setup Driver adapters
	if err := v1.RegisterServerTx(container, mono.RPC().GRPCServer()); err != nil {
		return err
	}
	if err := restv1.RegisterGateway(ctx, mono.Gin(), endpoint); err != nil {
		return err
	}
	if err := docs.RegisterSwagger(mono.Gin()); err != nil {
		return err
	}
	if err := handler.RegisterIntegrationEventHandlersTx(container); err != nil {
		return err
	}

	return nil
}

func (m *Module) Name() string {
	return "search"
}
