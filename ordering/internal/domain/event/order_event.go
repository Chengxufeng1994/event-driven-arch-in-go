package event

import "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/valueobject"

const (
	OrderCreatedEvent   = "ordering.OrderCreated"
	OrderRejectedEvent  = "ordering.OrderRejected"
	OrderApprovedEvent  = "ordering.OrderApproved"
	OrderCanceledEvent  = "ordering.OrderCanceled"
	OrderReadiedEvent   = "ordering.OrderReadied"
	OrderCompletedEvent = "ordering.OrderCompleted"
)

type OrderCreated struct {
	CustomerID string
	PaymentID  string
	ShoppingID string
	Items      []valueobject.Item
}

func NewOrderCreated(customerID, paymentID, shoppingID string, items []valueobject.Item) *OrderCreated {
	return &OrderCreated{
		CustomerID: customerID,
		PaymentID:  paymentID,
		ShoppingID: shoppingID,
		Items:      items,
	}
}

func (OrderCreated) Key() string { return OrderCreatedEvent }

type OrderRejected struct{}

func NewOrderRejected() *OrderRejected {
	return &OrderRejected{}
}

func (OrderRejected) Key() string { return OrderRejectedEvent }

type OrderApproved struct {
	ShoppingID string
}

func NewOrderApproved(shoppingID string) *OrderApproved {
	return &OrderApproved{
		ShoppingID: shoppingID,
	}
}

func (OrderApproved) Key() string { return OrderApprovedEvent }

type OrderCanceled struct {
	CustomerID string
	PaymentID  string
}

func NewOrderCanceled(customerID, paymentID string) *OrderCanceled {
	return &OrderCanceled{
		CustomerID: customerID,
		PaymentID:  paymentID,
	}
}

func (OrderCanceled) Key() string { return OrderCanceledEvent }

type OrderReadied struct {
	CustomerID string
	PaymentID  string
	Total      float64
}

func NewOrderReadied(customerID, paymentID string, total float64) *OrderReadied {
	return &OrderReadied{
		CustomerID: customerID,
		PaymentID:  paymentID,
		Total:      total,
	}
}

func (OrderReadied) Key() string { return OrderReadiedEvent }

type OrderCompleted struct {
	CustomerID string
	InvoiceID  string
}

func NewOrderCompleted(customerID, invoiceID string) *OrderCompleted {
	return &OrderCompleted{
		CustomerID: customerID,
		InvoiceID:  invoiceID,
	}
}

func (OrderCompleted) Key() string { return OrderCompletedEvent }
