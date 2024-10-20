package aggregate

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/valueobject"
	"github.com/stackus/errors"
)

var (
	ErrOrderHasNoItems         = errors.Wrap(errors.ErrBadRequest, "the order has no items")
	ErrOrderCannotBeCancelled  = errors.Wrap(errors.ErrBadRequest, "the order cannot be canceled")
	ErrCustomerIDCannotBeBlank = errors.Wrap(errors.ErrBadRequest, "the customer id cannot be blank")
	ErrPaymentIDCannotBeBlank  = errors.Wrap(errors.ErrBadRequest, "the payment id cannot be blank")
)

type OrderAgg struct {
	ID         string
	CustomerID string
	PaymentID  string
	InvoiceID  string
	ShoppingID string
	Items      []*valueobject.Item
	Status     valueobject.OrderStatus
}

func CreateOrder(id, customerID, paymentID string, items []*valueobject.Item) (*OrderAgg, error) {
	if len(items) == 0 {
		return nil, ErrOrderHasNoItems
	}

	if customerID == "" {
		return nil, ErrCustomerIDCannotBeBlank
	}

	if paymentID == "" {
		return nil, ErrPaymentIDCannotBeBlank
	}

	order := &OrderAgg{
		ID:         id,
		CustomerID: customerID,
		PaymentID:  paymentID,
		Items:      items,
		Status:     valueobject.OrderPending,
	}

	return order, nil
}

func (o *OrderAgg) Cancel() error {
	if o.Status != valueobject.OrderPending {
		return ErrOrderCannotBeCancelled
	}

	o.Status = valueobject.OrderCancelled

	return nil
}

func (o *OrderAgg) Ready() error {
	// validate status

	o.Status = valueobject.OrderReady

	return nil
}

func (o *OrderAgg) Complete(invoiceID string) error {
	// validate invoice exists

	// validate status

	o.InvoiceID = invoiceID
	o.Status = valueobject.OrderCompleted

	return nil
}

func (o *OrderAgg) GetTotal() float64 {
	var total float64

	for _, item := range o.Items {
		total += item.Price * float64(item.Quantity)
	}

	return total
}
