package valueobject

type InvoiceID string

func NewInvoiceID(id string) InvoiceID {
	return InvoiceID(id)
}

func (i InvoiceID) String() string {
	return string(i)
}
