package aggregate

import (
	"sort"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/entity"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/event"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/valueobject"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
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

type Basket struct {
	ddd.AggregateBase
	CustomerID string
	PaymentID  string
	Items      []*entity.Item
	Status     valueobject.BasketStatus
}

func StartBasket(id, customerID string) (*Basket, error) {
	if id == "" {
		return nil, ErrBasketIDCannotBeBlank
	}

	if customerID == "" {
		return nil, ErrCustomerIDCannotBeBlank
	}

	basket := &Basket{
		AggregateBase: ddd.NewAggregateBase(id),
		CustomerID:    customerID,
		Status:        valueobject.BasketOpen,
		Items:         []*entity.Item{},
	}

	basket.AddEvent(event.NewBasketStarted(customerID))

	return basket, nil
}

func (b *Basket) IsCancellable() bool {
	return b.Status == valueobject.BasketOpen
}

func (b *Basket) IsOpen() bool {
	return b.Status == valueobject.BasketOpen
}

func (b *Basket) Cancel() error {
	if !b.IsCancellable() {
		return ErrBasketCannotBeCancelled
	}

	b.Status = valueobject.BasketCancelled
	b.Items = []*entity.Item{}

	b.AddEvent(event.NewBasketCanceled())

	return nil
}

func (b *Basket) Checkout(paymentID string) error {
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

	b.AddEvent(event.NewBasketCheckedOut(paymentID, b.CustomerID, b.Items))

	return nil
}

func (b *Basket) AddItem(store valueobject.Store, product valueobject.Product, quantity int) error {
	if !b.IsOpen() {
		return ErrBasketCannotBeModified
	}

	if quantity < 0 {
		return ErrQuantityCannotBeNegative
	}

	item := entity.NewItem(store.ID, product.ID, store.Name, product.Name, product.Price, quantity)
	if i, exists := b.hasProduct(product); exists {
		b.Items[i].Quantity += quantity
	} else {
		b.Items = append(b.Items, item)

		sort.Slice(b.Items, func(i, j int) bool {
			return b.Items[i].StoreName <= b.Items[j].StoreName && b.Items[i].ProductName < b.Items[j].ProductName
		})

	}

	b.AddEvent(event.NewBasketItemAdded(item))

	return nil
}

func (b *Basket) RemoveItem(product valueobject.Product, quantity int) error {
	if !b.IsOpen() {
		return ErrBasketCannotBeModified
	}

	if quantity < 0 {
		return ErrQuantityCannotBeNegative
	}

	if i, exists := b.hasProduct(product); exists {
		b.Items[i].Quantity -= quantity

		if b.Items[i].Quantity < 1 {
			b.Items = append(b.Items[:i], b.Items[i+1:]...)
		}

		b.AddEvent(event.NewBasketItemRemoved(product.ID, quantity))
	}

	return nil
}

func (b *Basket) hasProduct(product valueobject.Product) (int, bool) {
	for i, item := range b.Items {
		if item.ProductID == product.ID && item.StoreID == product.StoreID {
			return i, true
		}
	}

	return 0, false
}
