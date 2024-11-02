package application

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	orderv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/api/order/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/application/port/int/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/application/usecase"
)

type OrderDomainEventHandler[T ddd.Event] struct {
	app usecase.PaymentUseCase
}

var _ ddd.EventHandler[ddd.Event] = (*OrderDomainEventHandler[ddd.Event])(nil)

func NewOrderDomainEventHandler(app usecase.PaymentUseCase) *OrderDomainEventHandler[ddd.Event] {
	return &OrderDomainEventHandler[ddd.Event]{
		app: app,
	}
}

func (h *OrderDomainEventHandler[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case orderv1.OrderReadiedEvent:
		return h.onOrderReadied(ctx, event)
	case orderv1.OrderCanceledEvent:
		return h.onOrderCanceled(ctx, event)
	}
	return nil
}

func (h *OrderDomainEventHandler[T]) onOrderReadied(ctx context.Context, event T) error {
	payload := event.Payload().(*orderv1.OrderReadied)
	return h.app.CreateInvoice(ctx, command.NewCreateInvoice(payload.GetId(), payload.GetId(), payload.GetPaymentId(), payload.GetTotal()))
}

func (h *OrderDomainEventHandler[T]) onOrderCanceled(ctx context.Context, event T) error {
	payload := event.Payload().(*orderv1.OrderCanceled)
	return h.app.CancelInvoice(ctx, command.NewCancelInvoice(payload.GetId()))
}
