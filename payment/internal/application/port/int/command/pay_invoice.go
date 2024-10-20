package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/application/port/out/client"
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
	orderClient       client.OrderClient
}

func NewPayInvoiceHandler(invoiceRepository repository.InvoiceRepository, orderClient client.OrderClient) PayInvoiceHandler {
	return PayInvoiceHandler{
		invoiceRepository: invoiceRepository,
		orderClient:       orderClient,
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

	if err := h.orderClient.Complete(ctx, invoice.ID, invoice.OrderID); err != nil {
		return errors.Wrap(err, "pay invoice command")
	}

	return errors.Wrap(h.invoiceRepository.Update(ctx, invoice), "pay invoice command")
}
