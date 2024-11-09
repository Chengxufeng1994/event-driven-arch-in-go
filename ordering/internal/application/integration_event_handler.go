package application

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	orderv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/api/order/v1"
	domainevent "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/event"
)

type IntegrationEventHandlers[T ddd.AggregateEvent] struct {
	publisher am.MessagePublisher[ddd.Event]
}

var _ ddd.EventHandler[ddd.AggregateEvent] = (*IntegrationEventHandlers[ddd.AggregateEvent])(nil)

func RegisterIntegrationEventHandlers(eventHandlers ddd.EventHandler[ddd.AggregateEvent], domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent]) {
	domainSubscriber.Subscribe(eventHandlers,
		domainevent.OrderCreatedEvent,
		domainevent.OrderReadiedEvent,
		domainevent.OrderCanceledEvent,
		domainevent.OrderCompletedEvent,
	)
}

func NewIntegrationEventHandlers(publisher am.MessagePublisher[ddd.Event]) *IntegrationEventHandlers[ddd.AggregateEvent] {
	return &IntegrationEventHandlers[ddd.AggregateEvent]{
		publisher: publisher,
	}
}

func (h IntegrationEventHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case domainevent.OrderCreatedEvent:
		return h.onOrderCreated(ctx, event)
	case domainevent.OrderReadiedEvent:
		return h.onOrderReadied(ctx, event)
	case domainevent.OrderCanceledEvent:
		return h.onOrderCanceled(ctx, event)
	case domainevent.OrderCompletedEvent:
		return h.onOrderCompleted(ctx, event)
	}
	return nil
}

func (h IntegrationEventHandlers[T]) onOrderCreated(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domainevent.OrderCreated)
	items := make([]*orderv1.OrderCreated_Item, len(payload.Items))
	for i, item := range payload.Items {
		items[i] = &orderv1.OrderCreated_Item{
			ProductId: item.ProductID,
			Price:     item.Price,
			Quantity:  int32(item.Quantity),
		}
	}
	return h.publisher.Publish(ctx, orderv1.OrderAggregateChannel,
		ddd.NewEvent(orderv1.OrderCreatedEvent, &orderv1.OrderCreated{
			Id:         event.AggregateID(),
			CustomerId: payload.CustomerID,
			PaymentId:  payload.PaymentID,
			ShoppingId: payload.ShoppingID,
			Items:      items,
		}),
	)
}

func (h IntegrationEventHandlers[T]) onOrderReadied(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domainevent.OrderReadied)
	return h.publisher.Publish(ctx, orderv1.OrderAggregateChannel,
		ddd.NewEvent(orderv1.OrderReadiedEvent, &orderv1.OrderReadied{
			Id:         event.AggregateID(),
			CustomerId: payload.CustomerID,
			PaymentId:  payload.PaymentID,
			Total:      payload.Total,
		}),
	)
}

func (h IntegrationEventHandlers[T]) onOrderCanceled(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domainevent.OrderCanceled)
	return h.publisher.Publish(ctx, orderv1.OrderAggregateChannel,
		ddd.NewEvent(orderv1.OrderCanceledEvent, &orderv1.OrderCanceled{
			Id:         event.AggregateID(),
			CustomerId: payload.CustomerID,
			PaymentId:  payload.PaymentID,
		}),
	)
}

func (h IntegrationEventHandlers[T]) onOrderCompleted(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domainevent.OrderCompleted)
	return h.publisher.Publish(ctx, orderv1.OrderAggregateChannel,
		ddd.NewEvent(orderv1.OrderCompletedEvent, &orderv1.OrderCompleted{
			Id:         event.AggregateID(),
			CustomerId: payload.CustomerID,
			InvoiceId:  payload.InvoiceID,
		}),
	)
}
