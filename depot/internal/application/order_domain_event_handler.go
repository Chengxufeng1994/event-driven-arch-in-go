package application

import (
	"context"
	"fmt"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/application/port/out/client"
	domainevent "github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/event"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
)

type OrderDomainEventHandler[T ddd.AggregateEvent] struct {
	orderClient client.OrderClient
}

var _ ddd.EventHandler[ddd.AggregateEvent] = (*OrderDomainEventHandler[ddd.AggregateEvent])(nil)

func NewShoppingListDomainEventHandler(orderClient client.OrderClient) *OrderDomainEventHandler[ddd.AggregateEvent] {
	return &OrderDomainEventHandler[ddd.AggregateEvent]{orderClient: orderClient}
}

// HandleEvent implements ddd.EventHandler.
func (h *OrderDomainEventHandler[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case domainevent.ShoppingListCompletedEvent:
		return h.OnShoppingListCompleted(ctx, event)
	default:
		return fmt.Errorf("unexpected event type: %T", event)
	}
}

func (h *OrderDomainEventHandler[T]) OnShoppingListCompleted(ctx context.Context, event ddd.AggregateEvent) error {
	shoppingListCompleted, ok := event.Payload().(*domainevent.ShoppingListCompleted)
	if !ok {
		return fmt.Errorf("unexpected event type: %T", event)
	}
	return h.orderClient.Ready(ctx, shoppingListCompleted.OrderID)
}
