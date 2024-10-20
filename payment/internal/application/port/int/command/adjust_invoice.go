package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/domain/repository"
	"github.com/stackus/errors"
)

type AdjustInvoice struct {
	ID     string
	Amount float64
}

func NewAdjustInvoice(id string, amount float64) AdjustInvoice {
	return AdjustInvoice{
		ID:     id,
		Amount: amount,
	}
}

type AdjustInvoiceHandler struct {
	invoiceRepository repository.InvoiceRepository
}

func NewAdjustInvoiceHandler(invoiceRepository repository.InvoiceRepository) AdjustInvoiceHandler {
	return AdjustInvoiceHandler{
		invoiceRepository: invoiceRepository,
	}
}

func (h AdjustInvoiceHandler) AdjustInvoice(ctx context.Context, adjust AdjustInvoice) error {
	invoice, err := h.invoiceRepository.Find(ctx, adjust.ID)
	if err != nil {
		return err
	}
	if invoice == nil {
		return errors.Wrap(errors.ErrNotFound, "adjust invoice command")
	}

	invoice.Amount = adjust.Amount

	return h.invoiceRepository.Update(ctx, invoice)
}
