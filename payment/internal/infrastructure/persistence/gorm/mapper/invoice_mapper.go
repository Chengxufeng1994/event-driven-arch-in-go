package mapper

import (
	"fmt"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/domain/valueobject"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/infrastructure/persistence/gorm/po"
	"github.com/shopspring/decimal"
)

type InvoiceMapperIntf interface {
	ToPersistence(invoice *aggregate.InvoiceAgg) (*po.Invoice, error)
	ToDomain(invoice *po.Invoice) (*aggregate.InvoiceAgg, error)
}

type InvoiceMapper struct{}

var _ InvoiceMapperIntf = (*InvoiceMapper)(nil)

func NewInvoiceMapper() *InvoiceMapper {
	return &InvoiceMapper{}
}

func (m *InvoiceMapper) ToPersistence(invoice *aggregate.InvoiceAgg) (*po.Invoice, error) {
	amount := decimal.NewFromFloat(invoice.Amount)
	status, err := m.statusFromDomain(invoice.Status.String())
	if err != nil {
		return nil, err
	}

	return &po.Invoice{
		ID:      invoice.ID,
		OrderID: invoice.OrderID,
		Amount:  amount,
		Status:  string(status),
	}, nil
}

func (m *InvoiceMapper) ToDomain(invoice *po.Invoice) (*aggregate.InvoiceAgg, error) {
	amount, _ := invoice.Amount.Float64()
	status, err := valueobject.NewInvoiceStatus(invoice.Status)
	if err != nil {
		return nil, err
	}

	return &aggregate.InvoiceAgg{
		ID:      invoice.ID,
		OrderID: invoice.OrderID,
		Amount:  amount,
		Status:  status,
	}, nil
}

func (m *InvoiceMapper) statusFromDomain(status string) (valueobject.InvoiceStatus, error) {
	switch status {
	case valueobject.InvoiceIsPending.String():
		return valueobject.InvoiceIsPending, nil
	case valueobject.InvoiceIsPaid.String():
		return valueobject.InvoiceIsPaid, nil
	case valueobject.InvoiceIsCanceled.String():
		return valueobject.InvoiceIsCanceled, nil
	default:
		return valueobject.InvoiceIsUnknown, fmt.Errorf("unknown invoice status: %s", status)
	}
}
