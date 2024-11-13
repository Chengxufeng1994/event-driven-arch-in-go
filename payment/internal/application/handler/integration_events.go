package handler

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	orderv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/api/order/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/application/port/int/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/application/usecase"
)

type IntegrationEventHandlers[T ddd.Event] struct {
	app usecase.PaymentUseCase
}

var _ ddd.EventHandler[ddd.Event] = (*IntegrationEventHandlers[ddd.Event])(nil)

func NewIntegrationEventHandlers(app usecase.PaymentUseCase) *IntegrationEventHandlers[ddd.Event] {
	return &IntegrationEventHandlers[ddd.Event]{
		app: app,
	}
}

func RegisterIntegrationEventHandlers(subscriber am.RawMessageStream, handlers am.RawMessageHandler) error {
	_, err := subscriber.Subscribe(orderv1.OrderAggregateChannel, handlers, am.MessageFilter{
		orderv1.OrderReadiedEvent,
	}, am.GroupName("payment-orders"))
	return err
}

func (h IntegrationEventHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case orderv1.OrderReadiedEvent:
		return h.onOrderReadied(ctx, event)
	case orderv1.OrderCanceledEvent:
		return h.onOrderCanceled(ctx, event)
	}
	return nil
}

func (h IntegrationEventHandlers[T]) onOrderReadied(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*orderv1.OrderReadied)
	return h.app.CreateInvoice(ctx, command.NewCreateInvoice(
		payload.GetId(),
		payload.GetId(),
		payload.GetPaymentId(),
		payload.GetTotal()))
}

func (h IntegrationEventHandlers[T]) onOrderCanceled(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*orderv1.OrderCanceled)
	return h.app.CancelInvoice(ctx, command.NewCancelInvoice(payload.GetId()))
}