package application

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/application/port/in/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/application/port/in/query"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/application/usecase"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/repository"
)

type (
	StoreApplication struct {
		appCommands
		appQueries
	}

	appCommands struct {
		command.CreateStoreHandler
		command.EnableParticipationHandler
		command.DisableParticipationHandler
		command.AddProductHandler
		command.RemoveProductHandler
	}

	appQueries struct {
		query.GetStoreHandler
		query.GetStoresHandler
		query.GetParticipatingStoresHandler
		query.GetCatalogHandler
		query.GetProductHandler
	}
)

var _ usecase.StoreUseCase = (*StoreApplication)(nil)

func NewStoreApplication(storeRepository repository.StoreRepository, participatingStoreRepository repository.ParticipatingStoreRepository, productRepository repository.ProductRepository) *StoreApplication {
	return &StoreApplication{
		appCommands: appCommands{
			CreateStoreHandler:          command.NewCreateStoreHandler(storeRepository),
			EnableParticipationHandler:  command.NewEnableParticipationHandler(storeRepository),
			DisableParticipationHandler: command.NewDisableParticipationHandler(storeRepository),
			AddProductHandler:           command.NewAddProductHandler(storeRepository, productRepository),
			RemoveProductHandler:        command.NewRemoveProductHandler(storeRepository, productRepository),
		},
		appQueries: appQueries{
			GetStoreHandler:               query.NewGetStoreHandler(storeRepository),
			GetStoresHandler:              query.NewGetStoresHandler(storeRepository),
			GetParticipatingStoresHandler: query.NewGetParticipatingStoresHandler(participatingStoreRepository),
			GetCatalogHandler:             query.NewGetCatalogHandler(productRepository),
			GetProductHandler:             query.NewGetProductHandler(productRepository),
		},
	}
}
