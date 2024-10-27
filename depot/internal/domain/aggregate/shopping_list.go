package aggregate

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/entity"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/event"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/valueobject"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/stackus/errors"
)

var (
	ErrShoppingCannotBeCanceled  = errors.Wrap(errors.ErrBadRequest, "the shopping list cannot be canceled")
	ErrShoppingCannotBeAssigned  = errors.Wrap(errors.ErrBadRequest, "the shopping list cannot be assigned")
	ErrShoppingCannotBeCompleted = errors.Wrap(errors.ErrBadRequest, "the shopping list cannot be completed")
)

type ShoppingList struct {
	ddd.AggregateBase
	OrderID       string
	AssignedBotID string
	Stops         entity.Stops
	Status        valueobject.ShoppingListStatus
}

func CreateShoppingList(id, orderID string) *ShoppingList {
	shoppingList := &ShoppingList{
		AggregateBase: ddd.NewAggregateBase(id),
		OrderID:       orderID,
		Status:        valueobject.ShoppingListIsAvailable,
		Stops:         entity.NewStops(),
	}

	shoppingList.AddEvent(event.NewShoppingListCreated())

	return shoppingList
}

func (shoppingList *ShoppingList) AddItem(store valueobject.Store, product valueobject.Product, quantity int) error {
	if _, exists := shoppingList.Stops[store.ID]; !exists {
		storeEnt := entity.NewStop(store.Name, store.Location)
		shoppingList.Stops[store.ID] = storeEnt
	}

	return shoppingList.Stops[store.ID].AddItem(product, quantity)
}

func (shoppingList *ShoppingList) isCancelable() bool {
	switch shoppingList.Status {
	case valueobject.ShoppingListIsAvailable, valueobject.ShoppingListIsAssigned, valueobject.ShoppingListIsActive:
		return true
	default:
		return false
	}
}

func (shoppingList *ShoppingList) Cancel() error {
	if !shoppingList.isCancelable() {
		return ErrShoppingCannotBeCanceled
	}

	shoppingList.Status = valueobject.ShoppingListIsCancelled

	shoppingList.AddEvent(event.NewShoppingListCanceled())

	return nil
}

func (shoppingList *ShoppingList) isAssignable() bool {
	return shoppingList.Status == valueobject.ShoppingListIsAvailable
}

func (shoppingList *ShoppingList) Assign(BotID string) error {
	if !shoppingList.isAssignable() {
		return ErrShoppingCannotBeAssigned
	}

	shoppingList.AssignedBotID = BotID
	shoppingList.Status = valueobject.ShoppingListIsAssigned

	shoppingList.AddEvent(event.NewShoppingListAssigned(BotID))

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

	shoppingList.AddEvent(event.NewShoppingListCompleted(shoppingList.OrderID))

	return nil
}
