package application

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/application/port/in/event"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/application/port/out/client"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
)

type (
	ShoppingListDomainEventHandler struct {
		appDomainEvents
	}

	appDomainEvents struct {
		event.OnShoppingListCompletedHandler
		event.IgnoreUnimplementedDomainEventHandler
	}
)

var _ event.DomainEventHandlers = (*ShoppingListDomainEventHandler)(nil)

func NewShoppingListDomainEventHandler(orderClient client.OrderClient) *ShoppingListDomainEventHandler {
	return &ShoppingListDomainEventHandler{
		appDomainEvents: appDomainEvents{
			OnShoppingListCompletedHandler:        event.NewOnShoppingListCompletedHandler(orderClient),
			IgnoreUnimplementedDomainEventHandler: event.NewIgnoreUnimplementedDomainEventHandler(),
		},
	}
}

func (h *ShoppingListDomainEventHandler) OnShoppingListCompleted(ctx context.Context, event ddd.DomainEvent) error {
	return h.OnShoppingListCompletedHandler.OnShoppingListCompleted(ctx, event)
}
