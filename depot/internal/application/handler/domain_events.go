package handler

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/application/port/out/client"
	domainevent "github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/event"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
)

type DomainEventHandlers[T ddd.AggregateEvent] struct {
	orders client.OrderClient
}

var _ ddd.EventHandler[ddd.AggregateEvent] = (*DomainEventHandlers[ddd.AggregateEvent])(nil)

func RegisterDomainEventHandlers(subscriber ddd.EventSubscriber[ddd.AggregateEvent], handler ddd.EventHandler[ddd.AggregateEvent]) {
	subscriber.Subscribe(handler, domainevent.ShoppingListCompletedEvent)
}

func NewDomainEventHandlers(orders client.OrderClient) ddd.EventHandler[ddd.AggregateEvent] {
	return &DomainEventHandlers[ddd.AggregateEvent]{
		orders: orders,
	}
}

func (h *DomainEventHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case domainevent.ShoppingListCompletedEvent:
		return h.onShoppingListCompleted(ctx, event)
	}

	return nil
}

func (h DomainEventHandlers[T]) onShoppingListCompleted(ctx context.Context, event ddd.AggregateEvent) error {
	completed := event.Payload().(*domainevent.ShoppingListCompleted)
	return h.orders.Ready(ctx, completed.OrderID)
}
