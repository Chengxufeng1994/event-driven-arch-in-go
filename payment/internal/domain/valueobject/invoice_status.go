package valueobject

type InvoiceStatus string

const (
	InvoiceUnknown  InvoiceStatus = ""
	InvoicePending  InvoiceStatus = "pending"
	InvoicePaid     InvoiceStatus = "paid"
	InvoiceCanceled InvoiceStatus = "canceled"
)

func NewInvoiceStatus(status string) (InvoiceStatus, error) {
	switch status {
	case InvoicePending.String():
		return InvoicePending, nil
	case InvoicePaid.String():
		return InvoicePaid, nil
	case InvoiceCanceled.String():
		return InvoiceCanceled, nil
	default:
		return InvoiceUnknown, nil
	}
}

func (s InvoiceStatus) String() string {
	switch s {
	case InvoicePending, InvoicePaid, InvoiceCanceled:
		return string(s)
	default:
		return ""
	}
}
