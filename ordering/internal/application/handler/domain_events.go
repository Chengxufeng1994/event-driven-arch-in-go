package handler

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	orderv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/api/order/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/event"
)

type domainEventHandlers[T ddd.Event] struct {
	publisher am.MessagePublisher[ddd.Event]
}

var _ ddd.EventHandler[ddd.Event] = (*domainEventHandlers[ddd.Event])(nil)

func RegisterDomainEventHandlers(subscriber ddd.EventSubscriber[ddd.Event], handlers ddd.EventHandler[ddd.Event]) {
	subscriber.Subscribe(handlers,
		event.OrderCreatedEvent,
		event.OrderRejectedEvent,
		event.OrderApprovedEvent,
		event.OrderReadiedEvent,
		event.OrderCanceledEvent,
		event.OrderCompletedEvent,
	)
}

func NewDomainEventHandler(publisher am.MessagePublisher[ddd.Event]) domainEventHandlers[ddd.Event] {
	return domainEventHandlers[ddd.Event]{publisher: publisher}
}

func (h domainEventHandlers[T]) HandleEvent(ctx context.Context, e T) error {
	switch e.EventName() {
	case event.OrderCreatedEvent:
		return h.onOrderCreated(ctx, e)
	case event.OrderReadiedEvent:
		return h.onOrderReadied(ctx, e)
	case event.OrderCanceledEvent:
		return h.onOrderCanceled(ctx, e)
	case event.OrderCompletedEvent:
		return h.onOrderCompleted(ctx, e)
	}
	return nil
}

func (h domainEventHandlers[T]) onOrderCreated(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*aggregate.Order)
	items := make([]*orderv1.OrderCreated_Item, len(payload.Items))
	for i, item := range payload.Items {
		items[i] = &orderv1.OrderCreated_Item{
			ProductId: item.ProductID,
			StoreId:   item.StoreID,
			Price:     item.Price,
			Quantity:  int32(item.Quantity),
		}
	}
	return h.publisher.Publish(ctx, orderv1.OrderAggregateChannel,
		ddd.NewEvent(orderv1.OrderCreatedEvent, &orderv1.OrderCreated{
			Id:         payload.ID(),
			CustomerId: payload.CustomerID,
			PaymentId:  payload.PaymentID,
			ShoppingId: payload.ShoppingID,
			Items:      items,
		}),
	)
}

func (h domainEventHandlers[T]) onOrderRejected(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*aggregate.Order)
	return h.publisher.Publish(ctx, orderv1.OrderAggregateChannel,
		ddd.NewEvent(orderv1.OrderRejectedEvent, &orderv1.OrderRejected{
			Id:         payload.ID(),
			CustomerId: payload.CustomerID,
			PaymentId:  payload.PaymentID,
		}),
	)
}

func (h domainEventHandlers[T]) onOrderApproved(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*aggregate.Order)
	return h.publisher.Publish(ctx, orderv1.OrderAggregateChannel,
		ddd.NewEvent(orderv1.OrderApprovedEvent, &orderv1.OrderApproved{
			Id:         payload.ID(),
			CustomerId: payload.CustomerID,
			PaymentId:  payload.PaymentID,
		}),
	)
}

func (h domainEventHandlers[T]) onOrderReadied(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*aggregate.Order)
	return h.publisher.Publish(ctx, orderv1.OrderAggregateChannel,
		ddd.NewEvent(orderv1.OrderReadiedEvent, &orderv1.OrderReadied{
			Id:         payload.ID(),
			CustomerId: payload.CustomerID,
			PaymentId:  payload.PaymentID,
			Total:      payload.GetTotal(),
		}),
	)
}

func (h domainEventHandlers[T]) onOrderCanceled(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*aggregate.Order)
	return h.publisher.Publish(ctx, orderv1.OrderAggregateChannel,
		ddd.NewEvent(orderv1.OrderCanceledEvent, &orderv1.OrderCanceled{
			Id:         payload.ID(),
			CustomerId: payload.CustomerID,
			PaymentId:  payload.PaymentID,
		}),
	)
}

func (h domainEventHandlers[T]) onOrderCompleted(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*aggregate.Order)
	return h.publisher.Publish(ctx, orderv1.OrderAggregateChannel,
		ddd.NewEvent(orderv1.OrderCompletedEvent, &orderv1.OrderCompleted{
			Id:         payload.ID(),
			CustomerId: payload.CustomerID,
			InvoiceId:  payload.InvoiceID,
		}),
	)
}
