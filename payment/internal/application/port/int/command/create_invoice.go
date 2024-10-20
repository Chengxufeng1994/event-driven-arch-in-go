package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/domain/valueobject"
)

type CreateInvoice struct {
	ID        string
	OrderID   string
	PaymentID string
	Amount    float64
}

func NewCreateInvoice(id, orderID, paymentID string, amount float64) CreateInvoice {
	return CreateInvoice{
		ID:        id,
		OrderID:   orderID,
		PaymentID: paymentID,
		Amount:    amount,
	}
}

type CreateInvoiceHandler struct {
	invoiceRepository repository.InvoiceRepository
}

func NewCreateInvoiceHandler(invoiceRepository repository.InvoiceRepository) CreateInvoiceHandler {
	return CreateInvoiceHandler{
		invoiceRepository: invoiceRepository,
	}
}

func (h CreateInvoiceHandler) CreateInvoice(ctx context.Context, create CreateInvoice) error {
	return h.invoiceRepository.Save(ctx,
		aggregate.NewInvoice(
			create.ID,
			create.OrderID,
			create.Amount,
			valueobject.InvoicePending,
		))
}
