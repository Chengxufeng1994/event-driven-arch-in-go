package application

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	paymentv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/payment/api/payment/v1"
	domainevent "github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/domain/event"
)

type IntegrationEventHandlers[T ddd.Event] struct {
	publisher am.MessagePublisher[ddd.Event]
}

var _ ddd.EventHandler[ddd.Event] = (*IntegrationEventHandlers[ddd.Event])(nil)

func NewIntegrationEventHandlers(publisher am.MessagePublisher[ddd.Event]) *IntegrationEventHandlers[ddd.Event] {
	return &IntegrationEventHandlers[ddd.Event]{
		publisher: publisher,
	}
}

func (h IntegrationEventHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case domainevent.InvoicePaidEvent:
		return h.onInvoicePaid(ctx, event)
	}
	return nil
}

func (h IntegrationEventHandlers[T]) onInvoicePaid(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*domainevent.InvoicePaid)
	return h.publisher.Publish(ctx, paymentv1.InvoiceAggregateChannel,
		ddd.NewEventBase(paymentv1.InvoicePaidEvent, &paymentv1.InvoicePaid{
			Id:      payload.ID,
			OrderId: payload.OrderID,
		}),
	)
}
