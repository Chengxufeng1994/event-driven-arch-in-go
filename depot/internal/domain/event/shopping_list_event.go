package event

import "github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/entity"

const (
	ShoppingListCreatedEvent   = "depot.ShoppingListCreated"
	ShoppingListCanceledEvent  = "depot.ShoppingListCanceled"
	ShoppingListAssignedEvent  = "depot.ShoppingListAssigned"
	ShoppingListCompletedEvent = "depot.ShoppingListCompleted"
)

type ShoppingListCreated struct {
	ShoppingListID string
	OrderID        string
	Stops          entity.Stops
}

func NewShoppingListCreated(shoppingListID, orderID string, stops entity.Stops) *ShoppingListCreated {
	return &ShoppingListCreated{
		ShoppingListID: shoppingListID,
		OrderID:        orderID,
		Stops:          stops,
	}
}

func (ShoppingListCreated) Key() string { return ShoppingListCreatedEvent }

type ShoppingListCanceled struct {
	ShoppingListID string
}

func NewShoppingListCanceled(shoppingListID string) *ShoppingListCanceled {
	return &ShoppingListCanceled{
		ShoppingListID: shoppingListID,
	}
}

func (ShoppingListCanceled) Key() string { return ShoppingListCanceledEvent }

type ShoppingListAssigned struct {
	ShoppingListID string
	BotID          string
}

func NewShoppingListAssigned(shoppingListID, botID string) *ShoppingListAssigned {
	return &ShoppingListAssigned{
		ShoppingListID: shoppingListID,
		BotID:          botID,
	}
}

func (ShoppingListAssigned) Key() string { return ShoppingListAssignedEvent }

type ShoppingListCompleted struct {
	ShoppingListID string
	OrderID        string
}

func NewShoppingListCompleted(shoppingListID, orderID string) *ShoppingListCompleted {
	return &ShoppingListCompleted{
		ShoppingListID: shoppingListID,
		OrderID:        orderID,
	}
}

func (ShoppingListCompleted) Key() string { return ShoppingListCompletedEvent }
