package event

import "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/valueobject"

const (
	OrderCreatedEvent   = "ordering.OrderCreated"
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

type OrderCanceled struct {
	CustomerID string
}

func NewOrderCanceled(customerID string) *OrderCanceled {
	return &OrderCanceled{
		CustomerID: customerID,
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
	InvoiceID string
}

func NewOrderCompleted(invoiceID string) *OrderCompleted {
	return &OrderCompleted{
		InvoiceID: invoiceID,
	}
}

func (OrderCompleted) Key() string { return OrderCompletedEvent }
