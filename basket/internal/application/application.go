package application

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/port/in/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/port/in/query"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/usecase"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
)

type (
	BasketApplication struct {
		appCommands
		appQueries
	}

	appCommands struct {
		command.StartBasketHandler
		command.CancelBasketHandler
		command.CheckoutBasketHandler
		command.AddItemHandler
		command.RemoveItemHandler
	}

	appQueries struct {
		query.GetBasketHandler
	}
)

var _ usecase.BasketUseCase = (*BasketApplication)(nil)

func New(
	baskets repository.BasketRepository,
	stores repository.StoreRepository,
	products repository.ProductRepository,
	publisher ddd.EventPublisher[ddd.Event],
) *BasketApplication {
	return &BasketApplication{
		appCommands: appCommands{
			StartBasketHandler:    command.NewStartBasketHandler(baskets, publisher),
			CancelBasketHandler:   command.NewCancelBasketHandler(baskets, publisher),
			CheckoutBasketHandler: command.NewCheckoutBasketHandler(baskets, publisher),
			AddItemHandler:        command.NewAddItemHandler(baskets, products, stores),
			RemoveItemHandler:     command.NewRemoveItemHandler(baskets, products),
		},
		appQueries: appQueries{
			GetBasketHandler: query.NewGetBasketHandler(baskets),
		},
	}
}
