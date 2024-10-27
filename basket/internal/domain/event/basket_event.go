package event

import "github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/entity"

const (
	BasketStartedEvent     = "baskets.BasketStarted"
	BasketItemAddedEvent   = "baskets.BasketItemAdded"
	BasketItemRemovedEvent = "baskets.BasketItemRemoved"
	BasketCanceledEvent    = "baskets.BasketCanceled"
	BasketCheckedOutEvent  = "baskets.BasketCheckedOut"
)

type BasketStarted struct {
	CustomerID string
}

func NewBasketStarted(customerID string) *BasketStarted {
	return &BasketStarted{
		CustomerID: customerID,
	}
}

func (BasketStarted) EventName() string { return BasketStartedEvent }

type BasketItemAdded struct {
	Item *entity.Item
}

func NewBasketItemAdded(item *entity.Item) *BasketItemAdded {
	return &BasketItemAdded{Item: item}
}

func (BasketItemAdded) EventName() string { return BasketItemAddedEvent }

type BasketItemRemoved struct {
	ProductID string
	Quantity  int
}

func NewBasketItemRemoved(productID string, quantity int) *BasketItemRemoved {
	return &BasketItemRemoved{
		ProductID: productID,
		Quantity:  quantity,
	}
}

func (BasketItemRemoved) EventName() string { return BasketItemRemovedEvent }

type BasketCanceled struct{}

func NewBasketCanceled() *BasketCanceled {
	return &BasketCanceled{}
}

func (BasketCanceled) EventName() string { return BasketCanceledEvent }

type BasketCheckedOut struct {
	PaymentID  string
	CustomerID string
	Items      []*entity.Item
}

func NewBasketCheckedOut(paymentID, customerID string, items []*entity.Item) *BasketCheckedOut {
	return &BasketCheckedOut{
		PaymentID:  paymentID,
		CustomerID: customerID,
		Items:      items,
	}
}

func (BasketCheckedOut) EventName() string { return BasketCheckedOutEvent }
