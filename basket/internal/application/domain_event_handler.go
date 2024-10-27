package application

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/port/in/event"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/port/out/client"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
)

type (
	BasketDomainEventHandler struct {
		appDomainEvents
	}

	appDomainEvents struct {
		event.IgnoreUnimplementedDomainEventHandler
		event.OnBasketCheckOutEventHandler
	}
)

var _ event.DomainEventHandlers = (*BasketDomainEventHandler)(nil)

func NewBasketDomainEventHandler(orderClient client.OrderClient) *BasketDomainEventHandler {
	return &BasketDomainEventHandler{
		appDomainEvents: appDomainEvents{
			OnBasketCheckOutEventHandler: event.NewOnBasketCheckOutEventHandler(orderClient),
		},
	}
}

func (b *BasketDomainEventHandler) OnBasketCheckedOut(ctx context.Context, event ddd.DomainEvent) error {
	return b.OnBasketCheckOutEventHandler.OnBasketCheckedOut(ctx, event)
}
