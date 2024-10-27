package mapper

import (
	"encoding/json"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/valueobject"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/infrastructure/persistence/gorm/po"
	"github.com/stackus/errors"
)

type OrderMapperIntf interface {
	ToPersistence(order *aggregate.Order) (*po.Order, error)
	ToDomain(order *po.Order) (*aggregate.Order, error)
}

type OrderMapper struct{}

var _ OrderMapperIntf = (*OrderMapper)(nil)

func NewOrderMapper() *OrderMapper {
	return &OrderMapper{}
}

func (m *OrderMapper) ToPersistence(order *aggregate.Order) (*po.Order, error) {
	byt, err := json.Marshal(order.Items)
	if err != nil {
		return nil, errors.Wrap(err, "json marshal")
	}

	return &po.Order{
		ID:         order.ID,
		CustomerID: order.CustomerID,
		PaymentID:  order.PaymentID,
		InvoiceID:  order.InvoiceID,
		ShoppingID: order.ShoppingID,
		Items:      byt,
		Status:     order.Status.String(),
	}, nil
}

func (m *OrderMapper) ToDomain(order *po.Order) (*aggregate.Order, error) {
	var items []valueobject.Item
	if err := json.Unmarshal(order.Items, &items); err != nil {
		return nil, errors.Wrap(err, "json unmarshal")
	}

	return &aggregate.Order{
		AggregateBase: ddd.NewAggregateBase(order.ID),
		CustomerID:    order.CustomerID,
		PaymentID:     order.PaymentID,
		InvoiceID:     order.InvoiceID,
		ShoppingID:    order.ShoppingID,
		Items:         items,
		Status:        valueobject.OrderStatus(order.Status),
	}, nil
}
