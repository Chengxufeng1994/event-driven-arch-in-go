package application

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application/port/out/client"
	domainevent "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/event"
)

type NotificationDomainEventHandler[T ddd.AggregateEvent] struct {
	notificationClient client.NotificationClient
}

var _ ddd.EventHandler[ddd.AggregateEvent] = (*NotificationDomainEventHandler[ddd.AggregateEvent])(nil)

func NewNotificationDomainEventHandler(notificationClient client.NotificationClient) *NotificationDomainEventHandler[ddd.AggregateEvent] {
	return &NotificationDomainEventHandler[ddd.AggregateEvent]{
		notificationClient: notificationClient,
	}
}

func (h *NotificationDomainEventHandler[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case domainevent.OrderCreatedEvent:
		return h.OnOrderCreated(ctx, event)
	case domainevent.OrderReadiedEvent:
		return h.OnOrderReadied(ctx, event)
	case domainevent.OrderCanceledEvent:
		return h.OnOrderCanceled(ctx, event)
	}
	return nil
}

func (h NotificationDomainEventHandler[T]) OnOrderCreated(ctx context.Context, event T) error {
	orderCreated := event.Payload().(*domainevent.OrderCreated)
	return h.notificationClient.NotifyOrderCreated(ctx, event.AggregateID(), orderCreated.CustomerID)
}

func (h NotificationDomainEventHandler[T]) OnOrderReadied(ctx context.Context, event T) error {
	orderReadied := event.Payload().(*domainevent.OrderReadied)
	return h.notificationClient.NotifyOrderReady(ctx, event.AggregateID(), orderReadied.CustomerID)
}

func (h NotificationDomainEventHandler[T]) OnOrderCanceled(ctx context.Context, event T) error {
	orderCanceled := event.Payload().(*domainevent.OrderCanceled)
	return h.notificationClient.NotifyOrderCanceled(ctx, event.AggregateID(), orderCanceled.CustomerID)
}
