package depot

import (
	"context"
	"fmt"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/docs"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/application"
	infragrpc "github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/infastructure/client/grpc"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/infastructure/logging"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/infastructure/persistence/gorm"
	grpcv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/interfaces/grpc/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/interfaces/handler"
	restv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/interfaces/rest/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/broker/nats"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/monolith"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
)

type Module struct{}

var _ monolith.Module = (*Module)(nil)

func NewModule() *Module { return &Module{} }

func (m *Module) PrepareRun(ctx context.Context, mono monolith.Monolith) error {
	// setup Driver adapters
	endpoint := fmt.Sprintf("%s:%d", mono.Config().Server.GPPC.Host, mono.Config().Server.GPPC.Port)
	conn, err := infragrpc.Dial(ctx, endpoint)
	if err != nil {
		return err
	}
	reg := registry.New()
	if err = storev1.Registrations(reg); err != nil {
		return err
	}

	eventStream := am.NewEventStream(reg, nats.NewStream(mono.Config().Infrastructure.Nats.Stream, mono.JetStream()))
	domainEventDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()
	shoppingListRepository := gorm.NewGormShoppingListRepository(mono.Database())
	grpcOrderClient := infragrpc.NewGrpcOrderClient(conn)
	grpcStoreClient := infragrpc.NewGrpcStoreClient(conn)
	grpcProductClient := infragrpc.NewGrpcProductClient(conn)
	storeCacheRepository := gorm.NewGormStoreCacheRepository(mono.Database(), grpcStoreClient)
	productCacheRepository := gorm.NewGormProductCacheRepository(mono.Database(), grpcProductClient)

	// setup application
	logApplication := logging.NewLogApplicationAccess(
		application.NewShoppingListApplication(
			shoppingListRepository,
			grpcStoreClient,
			grpcProductClient,
			grpcOrderClient,
			domainEventDispatcher,
		),
		mono.Logger(),
	)
	orderHandler := logging.NewLogDomainEventHandlerAccess[ddd.AggregateEvent](
		application.NewShoppingListDomainEventHandler(grpcOrderClient),
		"Order",
		mono.Logger())
	storeHandler := logging.NewLogDomainEventHandlerAccess[ddd.Event](
		application.NewStoreIntegrationEventHandler(storeCacheRepository),
		"Store",
		mono.Logger())
	productHandler := logging.NewLogDomainEventHandlerAccess[ddd.Event](
		application.NewProductIntegrationEventHandler(productCacheRepository),
		"Store",
		mono.Logger())

	// setup Driver adapters
	if err := grpcv1.RegisterServer(ctx, logApplication, mono.RPC().GRPCServer()); err != nil {
		return err
	}

	if err := restv1.RegisterGateway(ctx, mono.Gin(), endpoint); err != nil {
		return err
	}

	if err := docs.RegisterSwagger(mono.Gin()); err != nil {
		return err
	}

	handler.RegisterOrderDomainEventHandlers(orderHandler, domainEventDispatcher)
	if err := handler.RegisterStoreIntegrationEventHandlers(storeHandler, eventStream); err != nil {
		return err
	}
	if err := handler.RegisterProductIntegrationEventHandlers(productHandler, eventStream); err != nil {
		return err
	}

	return nil
}

func (m *Module) Name() string {
	return "depot"
}
