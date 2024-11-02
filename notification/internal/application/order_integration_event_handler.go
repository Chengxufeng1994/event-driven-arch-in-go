package application

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/application/port/in/command"
	orderv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/api/order/v1"
)

type OrderIntegrationEventHandler[T ddd.Event] struct {
	app NotificationUseCase
}

var _ ddd.EventHandler[ddd.Event] = (*OrderIntegrationEventHandler[ddd.Event])(nil)

func NewOrderIntegrationEventHandler(app NotificationUseCase) *OrderIntegrationEventHandler[ddd.Event] {
	return &OrderIntegrationEventHandler[ddd.Event]{
		app: app,
	}
}

func (h *OrderIntegrationEventHandler[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case orderv1.OrderCreatedEvent:
		return h.onOrderCreated(ctx, event)
	case orderv1.OrderReadiedEvent:
		return h.onOrderReadied(ctx, event)
	case orderv1.OrderCanceledEvent:
		return h.onOrderCanceled(ctx, event)
	}

	return nil
}

func (h *OrderIntegrationEventHandler[T]) onOrderCreated(ctx context.Context, event T) error {
	payload := event.Payload().(*orderv1.OrderCreated)
	return h.app.NotifyOrderCreated(ctx, command.NewOrderCreated(payload.GetId(), payload.GetCustomerId()))
}

func (h *OrderIntegrationEventHandler[T]) onOrderReadied(ctx context.Context, event T) error {
	payload := event.Payload().(*orderv1.OrderCreated)
	return h.app.NotifyOrderReady(ctx, command.NewOrderReady(payload.GetId(), payload.GetCustomerId()))
}

func (h *OrderIntegrationEventHandler[T]) onOrderCanceled(ctx context.Context, event T) error {
	payload := event.Payload().(*orderv1.OrderCanceled)
	return h.app.NotifyOrderCanceled(ctx, command.NewOrderCanceled(payload.GetId(), payload.GetCustomerId()))
}
