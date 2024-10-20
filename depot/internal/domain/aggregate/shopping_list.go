package aggregate

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/entity"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/valueobject"
	"github.com/stackus/errors"
)

var (
	ErrShoppingCannotBeCancelled = errors.Wrap(errors.ErrBadRequest, "the shopping list cannot be cancelled")
)

type ShoppingListAgg struct {
	ID            string
	OrderID       string
	AssignedBotID string
	Stops         entity.Stops
	Status        valueobject.ShoppingListStatus
}

func CreateShoppingList(id, orderID string) *ShoppingListAgg {
	return &ShoppingListAgg{
		ID:      id,
		OrderID: orderID,
		Status:  valueobject.ShoppingListAvailable,
		Stops:   entity.NewStops(),
	}
}

func (shoppingList *ShoppingListAgg) AddItem(store valueobject.Store, product valueobject.Product, quantity int) error {
	if _, exists := shoppingList.Stops[store.ID]; !exists {
		storeEnt := entity.NewStop(store.Name, store.Location)
		shoppingList.Stops[store.ID] = storeEnt
	}

	return shoppingList.Stops[store.ID].AddItem(product, quantity)
}

func (shoppingList *ShoppingListAgg) Cancel() error {
	// validate status

	shoppingList.Status = valueobject.ShoppingListCancelled

	return nil
}

func (shoppingList *ShoppingListAgg) Assign(id string) error {
	// validate status

	shoppingList.AssignedBotID = id
	shoppingList.Status = valueobject.ShoppingListAssigned

	return nil
}

func (shoppingList *ShoppingListAgg) Complete() error {
	// validate status

	shoppingList.Status = valueobject.ShoppingListCompleted

	return nil
}
