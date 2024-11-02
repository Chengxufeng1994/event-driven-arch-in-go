package aggregate

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/es"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/event"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/valueobject"
	"github.com/stackus/errors"
)

const OrderAggregate = "ordering.Order"

var (
	ErrOrderAlreadyCreated     = errors.Wrap(errors.ErrBadRequest, "the order cannot be recreated")
	ErrOrderHasNoItems         = errors.Wrap(errors.ErrBadRequest, "the order has no items")
	ErrOrderCannotBeCancelled  = errors.Wrap(errors.ErrBadRequest, "the order cannot be cancelled")
	ErrCustomerIDCannotBeBlank = errors.Wrap(errors.ErrBadRequest, "the customer id cannot be blank")
	ErrPaymentIDCannotBeBlank  = errors.Wrap(errors.ErrBadRequest, "the payment id cannot be blank")
)

type Order struct {
	es.AggregateBase
	CustomerID string
	PaymentID  string
	InvoiceID  string
	ShoppingID string
	Items      []valueobject.Item
	Status     valueobject.OrderStatus
}

var _ interface {
	es.EventApplier
	es.Snapshotter
} = (*Order)(nil)

func NewOrder(id string) *Order {
	return &Order{
		AggregateBase: es.NewAggregateBase(id, OrderAggregate),
	}
}

func (o *Order) CreateOrder(_, customerID, paymentID, shoppingID string, items []valueobject.Item) error {
	if o.Status != valueobject.OrderUnknown {
		return ErrOrderAlreadyCreated
	}

	if len(items) == 0 {
		return ErrOrderHasNoItems
	}

	if customerID == "" {
		return ErrCustomerIDCannotBeBlank
	}

	if paymentID == "" {
		return ErrPaymentIDCannotBeBlank
	}

	o.AddEvent(event.OrderCreatedEvent, event.NewOrderCreated(customerID, paymentID, shoppingID, items))

	return nil
}

func (Order) Key() string { return OrderAggregate }

func (o *Order) Cancel() error {
	if o.Status != valueobject.OrderIsPending {
		return ErrOrderCannotBeCancelled
	}

	o.AddEvent(event.OrderCanceledEvent, event.NewOrderCanceled(o.CustomerID, o.PaymentID))

	return nil
}

func (o *Order) Ready() error {
	// validate status

	o.AddEvent(event.OrderReadiedEvent, event.NewOrderReadied(o.CustomerID, o.PaymentID, o.GetTotal()))

	return nil
}

func (o *Order) Complete(invoiceID string) error {
	// validate invoice exists

	// validate status

	o.AddEvent(event.OrderCompletedEvent, event.NewOrderCompleted(o.CustomerID, invoiceID))

	return nil
}

func (o *Order) GetTotal() float64 {
	var total float64

	for _, item := range o.Items {
		total += item.Price * float64(item.Quantity)
	}

	return total
}

func (o *Order) ApplyEvent(evt ddd.Event) error {
	switch payload := evt.Payload().(type) {
	case *event.OrderCreated:
		o.CustomerID = payload.CustomerID
		o.PaymentID = payload.PaymentID
		o.ShoppingID = payload.ShoppingID
		o.Items = payload.Items
		o.Status = valueobject.OrderIsPending

	case *event.OrderCanceled:
		o.Status = valueobject.OrderIsCancelled

	case *event.OrderReadied:
		o.Status = valueobject.OrderIsReady

	case *event.OrderCompleted:
		o.InvoiceID = payload.InvoiceID
		o.Status = valueobject.OrderIsCompleted

	default:
		return errors.ErrInternal.Msgf("%T received the event %s with unexpected payload %T", o, evt.EventName(), payload)
	}

	return nil
}

func (o *Order) ApplySnapshot(snapshot es.Snapshot) error {
	switch ss := snapshot.(type) {
	case *OrderV1:
		o.CustomerID = ss.CustomerID
		o.PaymentID = ss.PaymentID
		o.InvoiceID = ss.InvoiceID
		o.ShoppingID = ss.ShoppingID
		o.Items = ss.Items
		o.Status = ss.Status
	default:
		return errors.ErrInternal.Msgf("%T received the unexpected snapshot %T", o, snapshot)
	}
	return nil
}

func (o *Order) ToSnapshot() es.Snapshot {
	return &OrderV1{
		CustomerID: o.CustomerID,
		PaymentID:  o.PaymentID,
		InvoiceID:  o.InvoiceID,
		ShoppingID: o.ShoppingID,
		Items:      o.Items,
		Status:     o.Status,
	}
}
