package aggregate

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/entity"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/event"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/valueobject"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/es"
	"github.com/stackus/errors"
)

const BasketAggregate = "baskets.Basket"

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
	es.Aggregate
	CustomerID string
	PaymentID  string
	Items      map[string]*entity.Item
	Status     valueobject.BasketStatus
}

var _ interface {
	es.EventApplier
	es.Snapshotter
} = (*Basket)(nil)

func NewBasket(id string) *Basket {
	return &Basket{
		Aggregate: es.NewAggregate(id, BasketAggregate),
		Items:     make(map[string]*entity.Item),
	}
}

func (Basket) Key() string { return BasketAggregate }

func (b *Basket) Start(customerID string) (ddd.Event, error) {
	if b.Status != valueobject.BasketUnknown {
		return nil, ErrBasketCannotBeModified
	}

	if customerID == "" {
		return nil, ErrCustomerIDCannotBeBlank
	}

	b.AddEvent(event.BasketStartedEvent, event.NewBasketStarted(customerID))

	return ddd.NewEvent(event.BasketStartedEvent, b), nil
}

func (b *Basket) IsCancellable() bool {
	return b.Status == valueobject.BasketIsOpen
}

func (b *Basket) IsOpen() bool {
	return b.Status == valueobject.BasketIsOpen
}

func (b *Basket) Cancel() (ddd.Event, error) {
	if !b.IsCancellable() {
		return nil, ErrBasketCannotBeCancelled
	}

	b.AddEvent(event.BasketCanceledEvent, event.NewBasketCanceled())

	return ddd.NewEvent(event.BasketCanceledEvent, b), nil
}

func (b *Basket) Checkout(paymentID string) (ddd.Event, error) {
	if !b.IsOpen() {
		return nil, ErrBasketCannotBeModified
	}

	if len(b.Items) == 0 {
		return nil, ErrBasketHasNoItems
	}

	if paymentID == "" {
		return nil, ErrPaymentIDCannotBeBlank
	}

	b.AddEvent(event.BasketCheckedOutEvent, event.NewBasketCheckedOut(paymentID, b.CustomerID, b.Items))

	return ddd.NewEvent(event.BasketCheckedOutEvent, b), nil
}

func (b *Basket) AddItem(store *entity.Store, product *entity.Product, quantity int) error {
	if !b.IsOpen() {
		return ErrBasketCannotBeModified
	}

	if quantity < 0 {
		return ErrQuantityCannotBeNegative
	}

	b.AddEvent(event.BasketItemAddedEvent, event.NewBasketItemAdded(entity.NewItem(
		store.ID,
		product.ID,
		store.Name,
		product.Name,
		product.Price,
		quantity,
	)))

	return nil
}

func (b *Basket) RemoveItem(product *entity.Product, quantity int) error {
	if !b.IsOpen() {
		return ErrBasketCannotBeModified
	}

	if quantity < 0 {
		return ErrQuantityCannotBeNegative
	}

	if exists := b.hasProduct(product); exists {
		b.AddEvent(event.BasketItemRemovedEvent, event.NewBasketItemRemoved(product.ID, quantity))
	}

	return nil
}

func (b *Basket) hasProduct(product *entity.Product) bool {
	for _, item := range b.Items {
		if item.ProductID == product.ID && item.StoreID == product.StoreID {
			return true
		}
	}

	return false
}

func (b *Basket) ApplyEvent(evt ddd.Event) error {
	switch payload := evt.Payload().(type) {
	case *event.BasketStarted:
		b.CustomerID = payload.CustomerID
		b.Status = valueobject.BasketIsOpen

	case *event.BasketItemAdded:
		if item, exists := b.Items[payload.Item.ProductID]; exists {
			item.Quantity += payload.Item.Quantity
			b.Items[payload.Item.ProductID] = item
		} else {
			b.Items[payload.Item.ProductID] = payload.Item
		}

	case *event.BasketItemRemoved:
		if item, exists := b.Items[payload.ProductID]; exists {
			if item.Quantity-payload.Quantity <= 1 {
				delete(b.Items, payload.ProductID)
			} else {
				item.Quantity -= payload.Quantity
				b.Items[payload.ProductID] = item
			}
		}

	case *event.BasketCanceled:
		b.Items = make(map[string]*entity.Item)
		b.Status = valueobject.BasketIsCanceled

	case *event.BasketCheckedOut:
		b.PaymentID = payload.PaymentID
		b.Status = valueobject.BasketIsCheckedOut

	default:
		return errors.ErrInternal.Msgf("%T received the event %s with unexpected payload %T", b, evt.EventName(), payload)
	}
	return nil
}

func (b *Basket) ApplySnapshot(snapshot es.Snapshot) error {
	switch ss := snapshot.(type) {
	case *BasketV1:
		b.CustomerID = ss.CustomerID
		b.PaymentID = ss.PaymentID
		b.Items = ss.Items
		b.Status = ss.Status

	default:
		return errors.ErrInternal.Msgf("%T received the unexpected snapshot %T", b, snapshot)
	}

	return nil
}

func (b *Basket) ToSnapshot() es.Snapshot {
	return &BasketV1{
		CustomerID: b.CustomerID,
		PaymentID:  b.PaymentID,
		Items:      b.Items,
		Status:     b.Status,
	}
}
