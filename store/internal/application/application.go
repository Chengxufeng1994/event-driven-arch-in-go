package application

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
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

func New(
	stores repository.StoreRepository,
	products repository.ProductRepository,
	mall repository.MallRepository,
	catalog repository.CatalogRepository,
	publisher ddd.EventPublisher[ddd.Event],
) *StoreApplication {
	return &StoreApplication{
		appCommands: appCommands{
			CreateStoreHandler:          command.NewCreateStoreHandler(stores, publisher),
			EnableParticipationHandler:  command.NewEnableParticipationHandler(stores, publisher),
			DisableParticipationHandler: command.NewDisableParticipationHandler(stores, publisher),
			RebrandStoreHandler:         command.NewRebrandStoreHandler(stores, publisher),
			AddProductHandler:           command.NewAddProductHandler(products, publisher),
			RebrandProductHandler:       command.NewRebrandProductHandler(products, publisher),
			IncreaseProductPriceHandler: command.NewIncreaseProductPriceHandler(products, publisher),
			DecreaseProductPriceHandler: command.NewDecreaseProductPriceHandler(products, publisher),
			RemoveProductHandler:        command.NewRemoveProductHandler(products, publisher),
		},
		appQueries: appQueries{
			GetStoreHandler:               query.NewGetStoreHandler(mall),
			GetStoresHandler:              query.NewGetStoresHandler(mall),
			GetParticipatingStoresHandler: query.NewGetParticipatingStoresHandler(mall),
			GetCatalogHandler:             query.NewGetCatalogHandler(catalog),
			GetProductHandler:             query.NewGetProductHandler(catalog),
		},
	}
}
