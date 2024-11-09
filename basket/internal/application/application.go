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

func NewBasketApplication(
	basketRepository repository.BasketRepository,
	productRepository repository.ProductRepository,
	storeRepository repository.StoreRepository,
	publisher ddd.EventPublisher[ddd.Event],
) *BasketApplication {
	return &BasketApplication{
		appCommands: appCommands{
			StartBasketHandler:    command.NewStartBasketHandler(basketRepository, publisher),
			CancelBasketHandler:   command.NewCancelBasketHandler(basketRepository, publisher),
			CheckoutBasketHandler: command.NewCheckoutBasketHandler(basketRepository, publisher),
			AddItemHandler:        command.NewAddItemHandler(basketRepository, productRepository, storeRepository),
			RemoveItemHandler:     command.NewRemoveItemHandler(basketRepository, productRepository),
		},
		appQueries: appQueries{
			GetBasketHandler: query.NewGetBasketHandler(basketRepository),
		},
	}
}
