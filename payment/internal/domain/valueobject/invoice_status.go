package valueobject

type InvoiceStatus string

const (
	InvoiceIsUnknown  InvoiceStatus = ""
	InvoiceIsPending  InvoiceStatus = "pending"
	InvoiceIsPaid     InvoiceStatus = "paid"
	InvoiceIsCanceled InvoiceStatus = "canceled"
)

func NewInvoiceStatus(status string) (InvoiceStatus, error) {
	switch status {
	case InvoiceIsPending.String():
		return InvoiceIsPending, nil
	case InvoiceIsPaid.String():
		return InvoiceIsPaid, nil
	case InvoiceIsCanceled.String():
		return InvoiceIsCanceled, nil
	default:
		return InvoiceIsUnknown, nil
	}
}

func (s InvoiceStatus) String() string {
	switch s {
	case InvoiceIsPending, InvoiceIsPaid, InvoiceIsCanceled:
		return string(s)
	default:
		return ""
	}
}
