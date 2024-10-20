package mapper

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/infrastructure/persistence/gorm/po"
	"github.com/shopspring/decimal"
)

type PaymentMapperIntf interface {
	ToPersistent(payment *aggregate.PaymentAgg) (*po.Payment, error)
	ToDomain(payment *po.Payment) (*aggregate.PaymentAgg, error)
}

type PaymentMapper struct{}

var _ PaymentMapperIntf = (*PaymentMapper)(nil)

func NewPaymentMapper() *PaymentMapper {
	return &PaymentMapper{}
}

func (p *PaymentMapper) ToPersistent(payment *aggregate.PaymentAgg) (*po.Payment, error) {
	amount := decimal.NewFromFloat(payment.Amount)
	return &po.Payment{
		ID:         payment.ID,
		CustomerID: payment.CustomerID,
		Amount:     amount,
	}, nil
}

func (p *PaymentMapper) ToDomain(payment *po.Payment) (*aggregate.PaymentAgg, error) {
	amount, _ := payment.Amount.Float64()
	return &aggregate.PaymentAgg{
		ID:         payment.ID,
		CustomerID: payment.CustomerID,
		Amount:     amount,
	}, nil
}
