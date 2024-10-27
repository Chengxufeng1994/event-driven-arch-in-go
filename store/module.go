package store

import (
	"context"
	"fmt"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/monolith"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/docs"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/application"
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
	endpoint := fmt.Sprintf("%s:%d", mono.Config().Server.GPPC.Host, mono.Config().Server.GPPC.Port)
	db := mono.Database()
	domainEventDispatcher := ddd.NewEventDispatcher()

	storeRepository := gorm.NewGormStoreRepository(db)
	participatingStoreRepository := gorm.NewGormParticipatingStoreRepository(db)
	productRepository := gorm.NewGormProductRepository(db)

	logApplication := logging.NewLogApplicationAccess(
		application.NewStoreApplication(
			storeRepository,
			participatingStoreRepository,
			productRepository,
			domainEventDispatcher,
		),
		mono.Logger())

	if err := v1.RegisterServer(ctx, logApplication, mono.RPC().GRPCServer()); err != nil {
		return err
	}

	if err := restv1.RegisterGateway(ctx, mono.Gin(), endpoint); err != nil {
		return err
	}

	return docs.RegisterSwagger(mono.Gin())
}

func (m *Module) Name() string {
	return "store"
}
