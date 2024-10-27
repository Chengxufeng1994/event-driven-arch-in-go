package application

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application/port/out/client"
	domainevent "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/event"
)

type InvoiceDomainEventHandler[T ddd.AggregateEvent] struct {
	invoiceClient client.InvoiceClient
}

var _ ddd.EventHandler[ddd.AggregateEvent] = (*InvoiceDomainEventHandler[ddd.AggregateEvent])(nil)

func NewInvoiceDomainEventHandler(invoiceClient client.InvoiceClient) *InvoiceDomainEventHandler[ddd.AggregateEvent] {
	return &InvoiceDomainEventHandler[ddd.AggregateEvent]{
		invoiceClient: invoiceClient,
	}
}

func (h *InvoiceDomainEventHandler[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case domainevent.OrderReadiedEvent:
		return h.OnOrderReadied(ctx, event)
	}

	return nil
}

func (h InvoiceDomainEventHandler[T]) OnOrderReadied(ctx context.Context, event T) error {
	orderReadied := event.Payload().(*domainevent.OrderReadied)
	return h.invoiceClient.Save(ctx, event.AggregateID(), orderReadied.PaymentID, orderReadied.Total)
}
