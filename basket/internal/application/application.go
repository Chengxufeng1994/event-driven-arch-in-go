package application

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/port/in/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/port/in/query"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/port/out/client"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/usecase"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/repository"
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
	orderClient client.OrderClient,
	productClient client.ProductClient,
	storeClient client.StoreClient,
) *BasketApplication {
	return &BasketApplication{
		appCommands: appCommands{
			StartBasketHandler:    command.NewStartBasketHandler(basketRepository),
			CancelBasketHandler:   command.NewCancelBasketHandler(basketRepository),
			CheckoutBasketHandler: command.NewCheckoutBasketHandler(basketRepository, orderClient),
			AddItemHandler:        command.NewAddItemHandler(basketRepository, productClient, storeClient),
			RemoveItemHandler:     command.NewRemoveItemHandler(basketRepository, productClient),
		},
		appQueries: appQueries{
			GetBasketHandler: query.NewGetBasketHandler(basketRepository),
		},
	}
}
