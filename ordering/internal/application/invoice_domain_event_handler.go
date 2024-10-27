package application

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application/port/in/event"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application/port/out/client"
	domainevent "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/event"
)

type (
	InvoiceDomainEventHandler struct {
		invoiceClient client.InvoiceClient
		event.IgnoreUnimplementedDomainEventHandler
	}
)

var _ event.DomainEventHandlers = (*InvoiceDomainEventHandler)(nil)

func NewInvoiceDomainEventHandler(invoiceClient client.InvoiceClient) *InvoiceDomainEventHandler {
	return &InvoiceDomainEventHandler{
		invoiceClient: invoiceClient,
	}
}

func (h InvoiceDomainEventHandler) OnOrderReadied(ctx context.Context, e ddd.DomainEvent) error {
	orderReadied := e.(*domainevent.OrderReadied)
	return h.invoiceClient.Save(ctx, orderReadied.OrderID, orderReadied.PaymentID, orderReadied.Total)
}
