package store

import (
	"context"
	"fmt"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/monolith"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/docs"
	applicationservice "github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/application/service"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/infrastructure/logging"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/infrastructure/persistence/gorm"
	v1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/interface/grpc/v1"
	restv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/interface/http/rest/v1"
)

type Module struct{}

var _ monolith.Module = (*Module)(nil)

func NewModule() *Module {
	return &Module{}
}

func (m *Module) PrepareRun(ctx context.Context, mono monolith.Monolith) error {
	db := mono.Database()

	storeRepository := gorm.NewGormStoreRepository(db)
	participatingStoreRepository := gorm.NewGormParticipatingStoreRepository(db)
	productRepository := gorm.NewGormProductRepository(db)

	logApplication := logging.NewLogApplicationAccess(
		applicationservice.NewStoreApplication(storeRepository, participatingStoreRepository, productRepository),
		mono.Logger())

	if err := v1.RegisterServer(ctx, logApplication, mono.RPC().GRPCServer()); err != nil {
		return err
	}

	endpoint := fmt.Sprintf("%s:%d", mono.Config().Server.GPPC.Host, mono.Config().Server.GPPC.Port)
	if err := restv1.RegisterGateway(ctx, mono.Gin(), endpoint); err != nil {
		return err
	}

	if err := docs.RegisterSwagger(mono.Gin()); err != nil {
		return err
	}

	return nil
}

func (m *Module) Name() string {
	return "store"
}
