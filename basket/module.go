package basket

import (
	"context"
	"fmt"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/docs"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application"
	infragrpc "github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/infrastructure/client/grpc"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/infrastructure/logging"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/infrastructure/persistence/gorm"
	grpcv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/interface/grpc/v1"
	restv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/interface/http/rest/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/monolith"
)

type Module struct{}

var _ monolith.Module = (*Module)(nil)

func NewModule() *Module {
	return &Module{}
}

func (m *Module) PrepareRun(ctx context.Context, mono monolith.Monolith) error {
	db := mono.Database()
	endpoint := fmt.Sprintf("%s:%d", mono.Config().Server.GPPC.Host, mono.Config().Server.GPPC.Port)

	conn, err := infragrpc.Dial(ctx, endpoint)
	if err != nil {
		return err
	}
	basketRepository := gorm.NewGormBasketRepository(db)
	grpcOrderClient := infragrpc.NewGrpcOrderClient(conn)
	grpcProductClient := infragrpc.NewGrpcProductClient(conn)
	grpcStoreClient := infragrpc.NewGrpcStoreClient(conn)

	application := logging.NewLogApplicationAccess(
		application.NewBasketApplication(basketRepository, grpcOrderClient, grpcProductClient, grpcStoreClient),
		mono.Logger())

	if err := grpcv1.RegisterServer(ctx, application, mono.RPC().GRPCServer()); err != nil {
		return err
	}

	if err := restv1.RegisterGateway(ctx, mono.Gin(), endpoint); err != nil {
		return err
	}

	if err := docs.RegisterSwagger(mono.Gin()); err != nil {
		return err
	}

	return nil
}

func (m *Module) Name() string {
	return "basket"
}
