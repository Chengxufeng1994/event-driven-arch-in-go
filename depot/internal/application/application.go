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
		command.InitiateShoppingHandler
		command.AssignShoppingListHandler
		command.CompleteShoppingListHandler
	}

	appQueries struct {
		query.GetShoppingListHandler
	}
)

var _ usecase.ShoppingListUseCase = (*ShoppingListApplication)(nil)

func NewShoppingListApplication(
	shoppingRepository repository.ShoppingListRepository,
	storeClient client.StoreClient,
	productClient client.ProductClient,
	orderClient client.OrderClient,
	publisher ddd.EventDispatcher[ddd.AggregateEvent],
) *ShoppingListApplication {
	return &ShoppingListApplication{
		appCommands: appCommands{
			CreateShoppingListHandler:   command.NewCreateShoppingListHandler(shoppingRepository, storeClient, productClient, publisher),
			CancelShoppingListHandler:   command.NewCancelShoppingListHandler(shoppingRepository, publisher),
			InitiateShoppingHandler:     command.NewInitiateShoppingHandler(shoppingRepository, publisher),
			AssignShoppingListHandler:   command.NewAssignShoppingListHandler(shoppingRepository, publisher),
			CompleteShoppingListHandler: command.NewCompleteShoppingListHandler(shoppingRepository, orderClient, publisher),
		},
		appQueries: appQueries{
			GetShoppingListHandler: query.NewGetShoppingListHandler(shoppingRepository),
		},
	}
}
