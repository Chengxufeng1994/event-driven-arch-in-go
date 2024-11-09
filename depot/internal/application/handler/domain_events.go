package handler

import (
	"context"

	depotv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/depot/api/depot/v1"
	domainevent "github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/event"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
)

type DomainEventHandlers[T ddd.AggregateEvent] struct {
	publisher am.MessagePublisher[ddd.Event]
}

var _ ddd.EventHandler[ddd.AggregateEvent] = (*DomainEventHandlers[ddd.AggregateEvent])(nil)

func RegisterDomainEventHandlers(subscriber ddd.EventSubscriber[ddd.AggregateEvent], handler ddd.EventHandler[ddd.AggregateEvent]) {
	subscriber.Subscribe(handler, domainevent.ShoppingListCompletedEvent)
}

func NewDomainEventHandlers(publisher am.MessagePublisher[ddd.Event]) ddd.EventHandler[ddd.AggregateEvent] {
	return &DomainEventHandlers[ddd.AggregateEvent]{
		publisher: publisher,
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
	return h.publisher.Publish(ctx, depotv1.ShoppingListAggregateChannel,
		ddd.NewEvent(depotv1.ShoppingListCompletedEvent, &depotv1.ShoppingListCompleted{
			Id:      event.AggregateID(),
			OrderId: completed.OrderID,
		}))
}
