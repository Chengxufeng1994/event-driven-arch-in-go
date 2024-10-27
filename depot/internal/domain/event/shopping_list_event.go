package event

const (
	ShoppingListCreatedEvent   = "depot.ShoppingListCreated"
	ShoppingListCanceledEvent  = "depot.ShoppingListCanceled"
	ShoppingListAssignedEvent  = "depot.ShoppingListAssigned"
	ShoppingListCompletedEvent = "depot.ShoppingListCompleted"
)

type ShoppingListCreated struct{}

func NewShoppingListCreated() *ShoppingListCreated {
	return &ShoppingListCreated{}
}

func (ShoppingListCreated) EventName() string { return ShoppingListCreatedEvent }

type ShoppingListCanceled struct{}

func NewShoppingListCanceled() *ShoppingListCanceled {
	return &ShoppingListCanceled{}
}

func (ShoppingListCanceled) EventName() string { return ShoppingListCanceledEvent }

type ShoppingListAssigned struct {
	BotID string
}

func NewShoppingListAssigned(botID string) *ShoppingListAssigned {
	return &ShoppingListAssigned{
		BotID: botID,
	}
}

func (ShoppingListAssigned) EventName() string { return ShoppingListAssignedEvent }

type ShoppingListCompleted struct {
	OrderID string
}

func NewShoppingListCompleted(orderID string) *ShoppingListCompleted {
	return &ShoppingListCompleted{
		OrderID: orderID,
	}
}

func (ShoppingListCompleted) EventName() string { return ShoppingListCompletedEvent }
