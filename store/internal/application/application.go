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

func NewStoreApplication(
	storeRepository repository.StoreRepository,
	participatingStoreRepository repository.ParticipatingStoreRepository,
	productRepository repository.ProductRepository,
	domainEventDispatcher ddd.EventDispatcherIntf,
) *StoreApplication {
	return &StoreApplication{
		appCommands: appCommands{
			CreateStoreHandler:          command.NewCreateStoreHandler(storeRepository, domainEventDispatcher),
			EnableParticipationHandler:  command.NewEnableParticipationHandler(storeRepository, domainEventDispatcher),
			DisableParticipationHandler: command.NewDisableParticipationHandler(storeRepository, domainEventDispatcher),
			AddProductHandler:           command.NewAddProductHandler(storeRepository, productRepository, domainEventDispatcher),
			RemoveProductHandler:        command.NewRemoveProductHandler(productRepository, domainEventDispatcher),
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
