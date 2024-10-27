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
	restv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/interfaces/http/rest/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/monolith"
)

type Module struct{}

var _ monolith.Module = (*Module)(nil)

func NewModule() *Module {
	return &Module{}
}

func (m *Module) PrepareRun(ctx context.Context, mono monolith.Monolith) error {
	endpoint := fmt.Sprintf("%s:%d", mono.Config().Server.GPPC.Host, mono.Config().Server.GPPC.Port)

	shoppingListRepository := gorm.NewGormShoppingListRepository(mono.Database())
	conn, err := infragrpc.Dial(ctx, endpoint)
	if err != nil {
		return err
	}

	domainEventDispatcher := ddd.NewEventDispatcher()
	grpcOrderClient := infragrpc.NewGrpcOrderClient(conn)
	grpcStoreClient := infragrpc.NewGrpcStoreClient(conn)
	grpcProductClient := infragrpc.NewGrpcProductClient(conn)

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

	logDomainEventHandler := logging.NewLogDomainEventHandlerAccess(
		application.NewShoppingListDomainEventHandler(grpcOrderClient),
		mono.Logger())

	handler.RegisterDomainEventHandlers(ctx, logDomainEventHandler, domainEventDispatcher)

	if err := grpcv1.RegisterServer(ctx, logApplication, mono.RPC().GRPCServer()); err != nil {
		return err
	}

	if err := restv1.RegisterGateway(ctx, mono.Gin(), endpoint); err != nil {
		return err
	}

	return docs.RegisterSwagger(mono.Gin())
}

func (m *Module) Name() string {
	return "depot"
}
