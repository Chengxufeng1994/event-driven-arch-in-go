package aggregate

import "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/valueobject"

type OrderV1 struct {
	CustomerID string
	PaymentID  string
	InvoiceID  string
	ShoppingID string
	Items      []valueobject.Item
	Status     valueobject.OrderStatus
}

func (OrderV1) SnapshotName() string { return "ordering.OrderV1" }
