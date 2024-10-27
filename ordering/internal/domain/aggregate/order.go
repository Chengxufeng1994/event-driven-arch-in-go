package aggregate

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/event"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/valueobject"
	"github.com/stackus/errors"
)

var (
	ErrOrderHasNoItems         = errors.Wrap(errors.ErrBadRequest, "the order has no items")
	ErrOrderCannotBeCancelled  = errors.Wrap(errors.ErrBadRequest, "the order cannot be canceled")
	ErrCustomerIDCannotBeBlank = errors.Wrap(errors.ErrBadRequest, "the customer id cannot be blank")
	ErrPaymentIDCannotBeBlank  = errors.Wrap(errors.ErrBadRequest, "the payment id cannot be blank")
)

type Order struct {
	ddd.AggregateBase
	CustomerID string
	PaymentID  string
	InvoiceID  string
	ShoppingID string
	Items      []valueobject.Item
	Status     valueobject.OrderStatus
}

func CreateOrder(id, customerID, paymentID string, items []valueobject.Item) (*Order, error) {
	if len(items) == 0 {
		return nil, ErrOrderHasNoItems
	}

	if customerID == "" {
		return nil, ErrCustomerIDCannotBeBlank
	}

	if paymentID == "" {
		return nil, ErrPaymentIDCannotBeBlank
	}

	order := &Order{
		AggregateBase: ddd.NewAggregateBase(id),
		CustomerID:    customerID,
		PaymentID:     paymentID,
		Items:         items,
		Status:        valueobject.OrderPending,
	}

	order.AddEvent(event.NewOrderCreated(customerID, paymentID, id, items))

	return order, nil
}

func (o *Order) Cancel() error {
	if o.Status != valueobject.OrderPending {
		return ErrOrderCannotBeCancelled
	}

	o.Status = valueobject.OrderCancelled

	o.AddEvent(event.NewOrderCanceled(o.CustomerID))

	return nil
}

func (o *Order) Ready() error {
	// validate status

	o.Status = valueobject.OrderReady

	o.AddEvent(event.NewOrderReadied(o.ID, o.CustomerID, o.PaymentID, o.GetTotal()))

	return nil
}

func (o *Order) Complete(invoiceID string) error {
	// validate invoice exists

	// validate status

	o.InvoiceID = invoiceID
	o.Status = valueobject.OrderCompleted

	o.AddEvent(event.NewOrderCompleted(invoiceID))

	return nil
}

func (o *Order) GetTotal() float64 {
	var total float64

	for _, item := range o.Items {
		total += item.Price * float64(item.Quantity)
	}

	return total
}
