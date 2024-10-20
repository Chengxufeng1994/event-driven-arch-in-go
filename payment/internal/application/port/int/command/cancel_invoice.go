package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/domain/repository"
	"github.com/stackus/errors"
)

type CancelInvoice struct {
	ID string
}

func NewCancelInvoice(id string) CancelInvoice {
	return CancelInvoice{
		ID: id,
	}
}

type CancelInvoiceHandler struct {
	invoiceRepository repository.InvoiceRepository
}

func NewCancelInvoiceHandler(invoiceRepository repository.InvoiceRepository) CancelInvoiceHandler {
	return CancelInvoiceHandler{
		invoiceRepository: invoiceRepository,
	}
}

func (h CancelInvoiceHandler) CancelInvoice(ctx context.Context, cancel CancelInvoice) error {
	invoice, err := h.invoiceRepository.Find(ctx, cancel.ID)
	if err != nil {
		return errors.Wrap(err, "cancel invoice command")
	}

	if !invoice.IsPending() {
		return errors.Wrap(errors.ErrBadRequest, "invoice cannot be canceled")
	}

	invoice.Cancel()

	return errors.Wrap(h.invoiceRepository.Update(ctx, invoice), "cancel invoice command")
}
