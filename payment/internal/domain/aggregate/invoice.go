package aggregate

import "github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/domain/valueobject"

type InvoiceAgg struct {
	ID      string
	OrderID string
	Status  valueobject.InvoiceStatus
	Amount  float64
}

func NewInvoice(id, orderID string, amount float64, status valueobject.InvoiceStatus) *InvoiceAgg {
	return &InvoiceAgg{
		ID:      id,
		OrderID: orderID,
		Amount:  amount,
		Status:  status,
	}
}

func (invoice *InvoiceAgg) AdjustAmount(amount float64) {
	invoice.Amount = amount
}

func (invoice *InvoiceAgg) Paid() {
	invoice.Status = valueobject.InvoicePaid
}

func (invoice *InvoiceAgg) Cancel() {
	invoice.Status = valueobject.InvoiceCanceled
}

func (invoice *InvoiceAgg) IsPending() bool {
	return invoice.Status == valueobject.InvoicePending
}
