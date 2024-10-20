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
	restv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/interfaces/http/rest/v1"
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

	grpcOrderClient := infragrpc.NewGrpcOrderClient(conn)
	grpcStoreClient := infragrpc.NewGrpcStoreClient(conn)
	grpcProductClient := infragrpc.NewGrpcProductClient(conn)

	application := logging.NewLogApplicationAccess(
		application.NewShoppingListApplication(
			shoppingListRepository,
			grpcStoreClient,
			grpcProductClient,
			grpcOrderClient),
		mono.Logger(),
	)

	if err := grpcv1.RegisterServer(ctx, application, mono.RPC().GRPCServer()); err != nil {
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
