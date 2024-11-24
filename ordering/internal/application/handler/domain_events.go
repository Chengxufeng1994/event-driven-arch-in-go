package handler

import (
	"context"
	"time"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/errorsotel"
	orderv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/api/order/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/event"
	domainevent "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/event"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type domainEventHandlers[T ddd.Event] struct {
	publisher am.EventPublisher
}

var _ ddd.EventHandler[ddd.Event] = (*domainEventHandlers[ddd.Event])(nil)

func NewDomainEventHandler(publisher am.EventPublisher) domainEventHandlers[ddd.Event] {
	return domainEventHandlers[ddd.Event]{publisher: publisher}
}

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

func (h domainEventHandlers[T]) HandleEvent(ctx context.Context, event T) (err error) {
	span := trace.SpanFromContext(ctx)
	defer func(started time.Time) {
		if err != nil {
			span.AddEvent(
				"Encountered an error handling domain event",
				trace.WithAttributes(errorsotel.ErrAttrs(err)...),
			)
		}
		span.AddEvent("Handled domain event", trace.WithAttributes(
			attribute.Int64("TookMS", time.Since(started).Milliseconds()),
		))
	}(time.Now())

	span.AddEvent("Handling domain event", trace.WithAttributes(
		attribute.String("Event", event.EventName()),
	))

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
