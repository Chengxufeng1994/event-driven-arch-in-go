package event

const InvoicePaidEvent = "payments.InvoicePaid"

type InvoicePaid struct {
	ID      string
	OrderID string
}
