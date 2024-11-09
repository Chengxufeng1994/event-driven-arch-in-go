package handler

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"

	paymentv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/payment/api/payment/v1"
	domainevent "github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/domain/event"
)

type DomainEventHandler[T ddd.Event] struct {
	publisher am.MessagePublisher[ddd.Event]
}

var _ ddd.EventHandler[ddd.Event] = (*DomainEventHandler[ddd.Event])(nil)

func NewDomainEventHandlers(publisher am.MessagePublisher[ddd.Event]) *DomainEventHandler[ddd.Event] {
	return &DomainEventHandler[ddd.Event]{
		publisher: publisher,
	}
}

func RegisterDomainEventHandlers(subscriber ddd.EventSubscriber[ddd.Event], handler ddd.EventHandler[ddd.Event]) {
	subscriber.Subscribe(handler,
		domainevent.InvoicePaidEvent,
	)
}

func (h *DomainEventHandler[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case domainevent.InvoicePaidEvent:
		return h.onInvoicePaid(ctx, event)
	}
	return nil
}

func (h *DomainEventHandler[T]) onInvoicePaid(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*domainevent.InvoicePaid)
	return h.publisher.Publish(ctx, paymentv1.InvoiceAggregateChannel,
		ddd.NewEvent(paymentv1.InvoicePaidEvent, &paymentv1.InvoicePaid{
			Id:      payload.ID,
			OrderId: payload.OrderID,
		}))
}
