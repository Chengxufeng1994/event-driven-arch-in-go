package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/domain/event"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/domain/repository"
	"github.com/stackus/errors"
)

type PayInvoice struct {
	ID string
}

func NewPayInvoice(id string) PayInvoice {
	return PayInvoice{
		ID: id,
	}
}

type PayInvoiceHandler struct {
	invoiceRepository repository.InvoiceRepository
	publisher         ddd.EventPublisher[ddd.Event]
}

func NewPayInvoiceHandler(invoiceRepository repository.InvoiceRepository, publisher ddd.EventPublisher[ddd.Event]) PayInvoiceHandler {
	return PayInvoiceHandler{
		invoiceRepository: invoiceRepository,
		publisher:         publisher,
	}
}

func (h PayInvoiceHandler) PayInvoice(ctx context.Context, pay PayInvoice) error {
	invoice, err := h.invoiceRepository.Find(ctx, pay.ID)
	if err != nil || invoice == nil {
		return errors.Wrap(errors.ErrNotFound, "pay invoice command")
	}

	if !invoice.IsPending() {
		return errors.Wrap(errors.ErrBadRequest, "invoice cannot be paid for")
	}

	invoice.Paid()

	// Before or after the invoice is saved we still risk something failing which
	// will leave the state change only partially complete
	if err = h.publisher.Publish(ctx, ddd.NewEvent(event.InvoicePaidEvent, &event.InvoicePaid{
		ID:      invoice.ID,
		OrderID: invoice.OrderID,
	})); err != nil {
		return err
	}

	return errors.Wrap(h.invoiceRepository.Update(ctx, invoice), "pay invoice command")
}
