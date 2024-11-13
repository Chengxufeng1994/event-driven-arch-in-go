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
	es.Aggregate
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
		Aggregate: es.NewAggregate(id, OrderAggregate),
	}
}

func (Order) Key() string { return OrderAggregate }

func (o *Order) CreateOrder(id, customerID, paymentID string, items []valueobject.Item) (ddd.Event, error) {
	if o.Status != valueobject.OrderUnknown {
		return nil, ErrOrderAlreadyCreated
	}

	if len(items) == 0 {
		return nil, ErrOrderHasNoItems
	}

	if customerID == "" {
		return nil, ErrCustomerIDCannotBeBlank
	}

	if paymentID == "" {
		return nil, ErrPaymentIDCannotBeBlank
	}

	o.AddEvent(event.OrderCreatedEvent, &event.OrderCreated{
		CustomerID: customerID,
		PaymentID:  paymentID,
		Items:      items,
	})

	return ddd.NewEvent(event.OrderCreatedEvent, o), nil
}

func (o *Order) Reject() (ddd.Event, error) {
	// validate status

	o.AddEvent(event.OrderRejectedEvent, event.NewOrderRejected())

	return ddd.NewEvent(event.OrderRejectedEvent, o), nil
}

func (o *Order) Approve(shoppingID string) (ddd.Event, error) {
	// validate status

	o.AddEvent(event.OrderApprovedEvent, event.NewOrderApproved(shoppingID))

	return ddd.NewEvent(event.OrderApprovedEvent, o), nil
}

func (o *Order) Cancel() (ddd.Event, error) {
	if o.Status != valueobject.OrderIsPending {
		return nil, ErrOrderCannotBeCancelled
	}

	o.AddEvent(event.OrderCanceledEvent, event.NewOrderCanceled(o.CustomerID, o.PaymentID))

	return ddd.NewEvent(event.OrderCanceledEvent, o), nil
}

func (o *Order) Ready() (ddd.Event, error) {
	// validate status

	o.AddEvent(event.OrderReadiedEvent, event.NewOrderReadied(
		o.CustomerID,
		o.PaymentID,
		o.GetTotal()))

	return ddd.NewEvent(event.OrderReadiedEvent, o), nil
}

func (o *Order) Complete(invoiceID string) (ddd.Event, error) {
	// validate invoice exists

	// validate status

	o.AddEvent(event.OrderCompletedEvent, event.NewOrderCompleted(
		o.CustomerID,
		invoiceID))

	return ddd.NewEvent(event.OrderCompletedEvent, o), nil
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

	case *event.OrderRejected:
		o.Status = valueobject.OrderIsRejected

	case *event.OrderApproved:
		o.ShoppingID = payload.ShoppingID
		o.Status = valueobject.OrderIsApproved

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
