package application

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/application/port/in/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/application/port/in/query"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/application/port/out/client"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/application/usecase"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
)

type (
	ShoppingListApplication struct {
		appCommands
		appQueries
	}

	appCommands struct {
		command.CreateShoppingListHandler
		command.CancelShoppingListHandler
		command.AssignShoppingListHandler
		command.CompleteShoppingListHandler
	}

	appQueries struct {
		query.GetShoppingListHandler
	}
)

var _ usecase.ShoppingListUseCase = (*ShoppingListApplication)(nil)

func NewShoppingListApplication(
	shoppingListRepository repository.ShoppingListRepository,
	storeClient client.StoreClient,
	productClient client.ProductClient,
	orderClient client.OrderClient,
	domainEventDispatcher ddd.EventDispatcherIntf,
) *ShoppingListApplication {
	return &ShoppingListApplication{
		appCommands: appCommands{
			CreateShoppingListHandler:   command.NewCreateShoppingListHandler(shoppingListRepository, storeClient, productClient, domainEventDispatcher),
			AssignShoppingListHandler:   command.NewAssignShoppingListHandler(shoppingListRepository, domainEventDispatcher),
			CancelShoppingListHandler:   command.NewCancelShoppingListHandler(shoppingListRepository, domainEventDispatcher),
			CompleteShoppingListHandler: command.NewCompleteShoppingListHandler(shoppingListRepository, orderClient, domainEventDispatcher),
		},
		appQueries: appQueries{
			GetShoppingListHandler: query.NewGetShoppingListHandler(shoppingListRepository),
		},
	}
}
