package depot

import (
	"context"
	"fmt"

	depotv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/depot/api/depot/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/docs"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/application"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/application/handler"
	infragrpc "github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/infastructure/client/grpc"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/infastructure/logging"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/infastructure/persistence/gorm"
	grpcv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/interfaces/grpc/v1"
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
	reg := registry.New()
	if err := storev1.Registrations(reg); err != nil {
		return err
	}
	if err := depotv1.Registrations(reg); err != nil {
		return err
	}
	stream := nats.NewStream(mono.Config().Infrastructure.Nats.Stream, mono.JetStream(), mono.Logger())
	eventStream := am.NewEventStream(reg, stream)
	commandStream := am.NewCommandStream(reg, stream)
	conn, err := infragrpc.Dial(ctx, endpoint)
	if err != nil {
		return err
	}
	domainEventDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()
	shoppingListRepository := gorm.NewGormShoppingListRepository(mono.Database())
	orderClient := infragrpc.NewGrpcOrderClient(conn)
	storeClient := infragrpc.NewGrpcStoreClient(conn)
	productClient := infragrpc.NewGrpcProductClient(conn)
	storeCacheRepository := gorm.NewGormStoreCacheRepository(mono.Database(), storeClient)
	productCacheRepository := gorm.NewGormProductCacheRepository(mono.Database(), productClient)

	// setup application
	app := logging.NewLogApplicationAccess(
		application.NewShoppingListApplication(
			shoppingListRepository,
			storeClient,
			productClient,
			orderClient,
			domainEventDispatcher,
		),
		mono.Logger(),
	)
	domainEventHandler := logging.NewLogEventHandlerAccess[ddd.AggregateEvent](
		handler.NewDomainEventHandlers(orderClient),
		"DomainEvents", mono.Logger(),
	)
	integrationEventHandler := logging.NewLogEventHandlerAccess[ddd.Event](
		handler.NewIntegrationEventHandlers(productCacheRepository, storeCacheRepository),
		"IntegrationEvent", mono.Logger(),
	)
	commandHandler := logging.NewLogCommandHandlerAccess(
		handler.NewCommandHandlers(app),
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

	handler.RegisterDomainEventHandlers(domainEventDispatcher, domainEventHandler)
	if err := handler.RegisterIntegrationEventHandlers(eventStream, integrationEventHandler); err != nil {
		return err
	}
	if err := handler.RegisterCommandHandlers(commandStream, commandHandler); err != nil {
		return err
	}

	return nil
}

func (m *Module) Name() string {
	return "depot"
}
