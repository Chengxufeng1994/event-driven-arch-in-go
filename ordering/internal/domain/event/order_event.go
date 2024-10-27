package event

import "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/valueobject"

const (
	OrderCreatedEvent   = "ordering.OrderCreated"
	OrderCanceledEvent  = "ordering.OrderCanceled"
	OrderReadiedEvent   = "ordering.OrderReadied"
	OrderCompletedEvent = "ordering.OrderCompleted"
)

type OrderCreated struct {
	OrderID    string
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

func (OrderCreated) EventName() string { return OrderCreatedEvent }

type OrderCanceled struct {
	OrderID    string
	CustomerID string
}

func NewOrderCanceled(customerID string) *OrderCanceled {
	return &OrderCanceled{
		CustomerID: customerID,
	}
}

func (OrderCanceled) EventName() string { return OrderCanceledEvent }

type OrderReadied struct {
	OrderID    string
	CustomerID string
	PaymentID  string
	Total      float64
}

func NewOrderReadied(orderID, customerID, paymentID string, total float64) *OrderReadied {
	return &OrderReadied{
		OrderID:    orderID,
		CustomerID: customerID,
		PaymentID:  paymentID,
		Total:      total,
	}
}

func (OrderReadied) EventName() string { return OrderReadiedEvent }

type OrderCompleted struct {
	OrderID   string
	InvoiceID string
}

func NewOrderCompleted(invoiceID string) *OrderCompleted {
	return &OrderCompleted{
		InvoiceID: invoiceID,
	}
}

func (OrderCompleted) EventName() string { return OrderCompletedEvent }
