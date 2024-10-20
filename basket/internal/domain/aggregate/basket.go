package aggregate

import (
	"sort"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/entity"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/valueobject"
	"github.com/stackus/errors"
)

var (
	ErrBasketHasNoItems         = errors.Wrap(errors.ErrBadRequest, "the basket has no items")
	ErrBasketCannotBeModified   = errors.Wrap(errors.ErrBadRequest, "the basket cannot be modified")
	ErrBasketCannotBeCancelled  = errors.Wrap(errors.ErrBadRequest, "the basket cannot be canceled")
	ErrQuantityCannotBeNegative = errors.Wrap(errors.ErrBadRequest, "the item quantity cannot be negative")
	ErrBasketIDCannotBeBlank    = errors.Wrap(errors.ErrBadRequest, "the basket id cannot be blank")
	ErrPaymentIDCannotBeBlank   = errors.Wrap(errors.ErrBadRequest, "the payment id cannot be blank")
	ErrCustomerIDCannotBeBlank  = errors.Wrap(errors.ErrBadRequest, "the customer id cannot be blank")
)

type BasketAgg struct {
	ID         string
	CustomerID string
	PaymentID  string
	Items      []*entity.Item
	Status     valueobject.BasketStatus
}

func StartBasket(id, customerID string) (*BasketAgg, error) {
	if id == "" {
		return nil, ErrBasketIDCannotBeBlank
	}

	if customerID == "" {
		return nil, ErrCustomerIDCannotBeBlank
	}

	basket := &BasketAgg{
		ID:         id,
		CustomerID: customerID,
		Status:     valueobject.BasketOpen,
		Items:      []*entity.Item{},
	}

	return basket, nil
}

func (b *BasketAgg) IsCancellable() bool {
	return b.Status == valueobject.BasketOpen
}

func (b *BasketAgg) IsOpen() bool {
	return b.Status == valueobject.BasketOpen
}

func (b *BasketAgg) Cancel() error {
	if !b.IsCancellable() {
		return ErrBasketCannotBeCancelled
	}

	b.Status = valueobject.BasketCancelled
	b.Items = []*entity.Item{}

	return nil
}

func (b *BasketAgg) Checkout(paymentID string) error {
	if !b.IsOpen() {
		return ErrBasketCannotBeModified
	}

	if len(b.Items) == 0 {
		return ErrBasketHasNoItems
	}

	if paymentID == "" {
		return ErrPaymentIDCannotBeBlank
	}

	b.PaymentID = paymentID
	b.Status = valueobject.BasketCheckedOut

	return nil
}

func (b *BasketAgg) AddItem(store valueobject.Store, product valueobject.Product, quantity int) error {
	if !b.IsOpen() {
		return ErrBasketCannotBeModified
	}

	if quantity < 0 {
		return ErrQuantityCannotBeNegative
	}

	for i, item := range b.Items {
		if item.ProductID == product.ID && item.StoreID == product.StoreID {
			b.Items[i].Quantity += quantity
			return nil
		}
	}

	b.Items = append(b.Items, entity.NewItem(store.ID, product.ID, store.Name, product.Name, product.Price, quantity))

	sort.Slice(b.Items, func(i, j int) bool {
		return b.Items[i].StoreName <= b.Items[j].StoreName && b.Items[i].ProductName < b.Items[j].ProductName
	})

	return nil
}

func (b *BasketAgg) RemoveItem(product valueobject.Product, quantity int) error {
	if !b.IsOpen() {
		return ErrBasketCannotBeModified
	}

	if quantity < 0 {
		return ErrQuantityCannotBeNegative
	}

	for i, item := range b.Items {
		if item.ProductID == product.ID && item.StoreID == product.StoreID {
			b.Items[i].Quantity -= quantity

			if b.Items[i].Quantity < 1 {
				b.Items = append(b.Items[:i], b.Items[i+1:]...)
			}
			return nil
		}
	}

	return nil
}
