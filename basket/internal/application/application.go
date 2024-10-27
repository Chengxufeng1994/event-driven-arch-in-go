package application

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/port/in/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/port/in/event"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/port/in/query"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/port/out/client"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/usecase"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
)

type (
	BasketApplication struct {
		appCommands
		appQueries
		appDomainEvents
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
	domainEventDispatcher ddd.EventDispatcherIntf,
) *BasketApplication {
	return &BasketApplication{
		appCommands: appCommands{
			StartBasketHandler:    command.NewStartBasketHandler(basketRepository, domainEventDispatcher),
			CancelBasketHandler:   command.NewCancelBasketHandler(basketRepository, domainEventDispatcher),
			CheckoutBasketHandler: command.NewCheckoutBasketHandler(basketRepository, orderClient, domainEventDispatcher),
			AddItemHandler:        command.NewAddItemHandler(basketRepository, productClient, storeClient, domainEventDispatcher),
			RemoveItemHandler:     command.NewRemoveItemHandler(basketRepository, productClient, domainEventDispatcher),
		},
		appQueries: appQueries{
			GetBasketHandler: query.NewGetBasketHandler(basketRepository),
		},
		appDomainEvents: appDomainEvents{
			OnBasketCheckOutEventHandler: event.NewOnBasketCheckOutEventHandler(orderClient),
		},
	}
}
