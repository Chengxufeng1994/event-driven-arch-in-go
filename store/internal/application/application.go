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
		command.RebrandStoreHandler
		command.EnableParticipationHandler
		command.DisableParticipationHandler
		command.AddProductHandler
		command.RemoveProductHandler
		command.RebrandProductHandler
		command.IncreaseProductPriceHandler
		command.DecreaseProductPriceHandler
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

func NewStoreApplication(
	storeRepository repository.StoreRepository,
	productRepository repository.ProductRepository,
	mallRepository repository.MallRepository,
	catalogRepository repository.CatalogRepository,
) *StoreApplication {
	return &StoreApplication{
		appCommands: appCommands{
			CreateStoreHandler:          command.NewCreateStoreHandler(storeRepository),
			RebrandStoreHandler:         command.NewRebrandStoreHandler(storeRepository),
			EnableParticipationHandler:  command.NewEnableParticipationHandler(storeRepository),
			DisableParticipationHandler: command.NewDisableParticipationHandler(storeRepository),
			AddProductHandler:           command.NewAddProductHandler(storeRepository, productRepository),
			RemoveProductHandler:        command.NewRemoveProductHandler(productRepository),
			RebrandProductHandler:       command.NewRebrandProductHandler(productRepository),
			IncreaseProductPriceHandler: command.NewIncreaseProductPriceHandler(productRepository),
			DecreaseProductPriceHandler: command.NewDecreaseProductPriceHandler(productRepository),
		},
		appQueries: appQueries{
			GetStoreHandler:               query.NewGetStoreHandler(mallRepository),
			GetStoresHandler:              query.NewGetStoresHandler(mallRepository),
			GetParticipatingStoresHandler: query.NewGetParticipatingStoresHandler(mallRepository),
			GetCatalogHandler:             query.NewGetCatalogHandler(catalogRepository),
			GetProductHandler:             query.NewGetProductHandler(catalogRepository),
		},
	}
}
