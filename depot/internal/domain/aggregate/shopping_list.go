package aggregate

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/entity"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/event"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/valueobject"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/stackus/errors"
)

const ShoppingListAggregate = "depot.ShoppingList"

var (
	ErrShoppingCannotBeCanceled  = errors.Wrap(errors.ErrBadRequest, "the shopping list cannot be canceled")
	ErrShoppingCannotBeInitiated = errors.Wrap(errors.ErrBadRequest, "the shopping list cannot be initiated")
	ErrShoppingCannotBeAssigned  = errors.Wrap(errors.ErrBadRequest, "the shopping list cannot be assigned")
	ErrShoppingCannotBeCompleted = errors.Wrap(errors.ErrBadRequest, "the shopping list cannot be completed")
)

type ShoppingList struct {
	ddd.AggregateBase
	OrderID       string
	Stops         entity.Stops
	AssignedBotID string
	Status        valueobject.ShoppingListStatus
}

var _ ddd.Aggregate = (*ShoppingList)(nil)

func NewShoppingList(id string) *ShoppingList {
	return &ShoppingList{
		AggregateBase: ddd.NewAggregateBase(id, ShoppingListAggregate),
	}
}

func CreateShoppingList(id, orderID string) *ShoppingList {
	shoppingList := NewShoppingList(id)
	shoppingList.OrderID = orderID
	shoppingList.Status = valueobject.ShoppingListIsPending
	shoppingList.Stops = entity.NewStops()

	shoppingList.AddEvent(event.ShoppingListCreatedEvent, event.NewShoppingListCreated(
		shoppingList.ID(),
		shoppingList.OrderID,
		shoppingList.Stops,
	))

	return shoppingList
}

func (ShoppingList) Key() string { return ShoppingListAggregate }

func (shoppingList *ShoppingList) AddItem(store *valueobject.Store, product *valueobject.Product, quantity int) error {
	if _, exists := shoppingList.Stops[store.ID]; !exists {
		storeEnt := entity.NewStop(store.Name, store.Location)
		shoppingList.Stops[store.ID] = storeEnt
	}

	return shoppingList.Stops[store.ID].AddItem(product, quantity)
}

func (shoppingList *ShoppingList) isCancelable() bool {
	switch shoppingList.Status {
	case valueobject.ShoppingListIsPending,
		valueobject.ShoppingListIsAvailable,
		valueobject.ShoppingListIsAssigned,
		valueobject.ShoppingListIsActive:
		return true
	default:
		return false
	}
}

func (shoppingList *ShoppingList) Cancel() error {
	if !shoppingList.isCancelable() {
		return ErrShoppingCannotBeCanceled
	}

	shoppingList.Status = valueobject.ShoppingListIsCanceled

	shoppingList.AddEvent(event.ShoppingListCanceledEvent,
		event.NewShoppingListCanceled(shoppingList.ID()))

	return nil
}

func (shoppingList *ShoppingList) isPending() bool {
	return shoppingList.Status == valueobject.ShoppingListIsPending
}

func (shoppingList *ShoppingList) Initiate() error {
	if !shoppingList.isPending() {
		return ErrShoppingCannotBeInitiated
	}

	shoppingList.Status = valueobject.ShoppingListIsAvailable

	shoppingList.AddEvent(event.ShoppingListInitiatedEvent,
		event.NewShoppingListInitiated(shoppingList.ID()))

	return nil
}

func (shoppingList *ShoppingList) isAssignable() bool {
	return shoppingList.Status == valueobject.ShoppingListIsAvailable
}

func (shoppingList *ShoppingList) Assign(botID string) error {
	if !shoppingList.isAssignable() {
		return ErrShoppingCannotBeAssigned
	}

	shoppingList.AssignedBotID = botID
	shoppingList.Status = valueobject.ShoppingListIsAssigned

	shoppingList.AddEvent(event.ShoppingListAssignedEvent,
		event.NewShoppingListAssigned(
			shoppingList.ID(),
			botID,
		))

	return nil
}

func (shoppingList *ShoppingList) isCompletable() bool {
	return shoppingList.Status == valueobject.ShoppingListIsAssigned
}

func (shoppingList *ShoppingList) Complete() error {
	if !shoppingList.isCompletable() {
		return ErrShoppingCannotBeCompleted
	}

	shoppingList.Status = valueobject.ShoppingListIsCompleted

	shoppingList.AddEvent(event.ShoppingListCompletedEvent,
		event.NewShoppingListCompleted(
			shoppingList.ID(),
			shoppingList.OrderID,
		))

	return nil
}
