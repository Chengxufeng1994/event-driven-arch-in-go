package application

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/application/port/in/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/application/port/in/query"
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
		command.InitiateShoppingHandler
		command.AssignShoppingListHandler
		command.CompleteShoppingListHandler
	}

	appQueries struct {
		query.GetShoppingListHandler
	}
)

var _ usecase.ShoppingListUseCase = (*ShoppingListApplication)(nil)

func New(
	shoppingList repository.ShoppingListRepository,
	stores repository.StoreCacheRepository,
	products repository.ProductCacheRepository,
	publisher ddd.EventDispatcher[ddd.AggregateEvent],
) *ShoppingListApplication {
	return &ShoppingListApplication{
		appCommands: appCommands{
			CreateShoppingListHandler:   command.NewCreateShoppingListHandler(shoppingList, stores, products, publisher),
			CancelShoppingListHandler:   command.NewCancelShoppingListHandler(shoppingList, publisher),
			InitiateShoppingHandler:     command.NewInitiateShoppingHandler(shoppingList, publisher),
			AssignShoppingListHandler:   command.NewAssignShoppingListHandler(shoppingList, publisher),
			CompleteShoppingListHandler: command.NewCompleteShoppingListHandler(shoppingList, publisher),
		},
		appQueries: appQueries{
			GetShoppingListHandler: query.NewGetShoppingListHandler(shoppingList),
		},
	}
}
