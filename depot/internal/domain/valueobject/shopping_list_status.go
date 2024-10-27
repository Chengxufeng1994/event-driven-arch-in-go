package valueobject

type ShoppingListStatus string

const (
	ShoppingListUnknown     ShoppingListStatus = ""
	ShoppingListIsAvailable ShoppingListStatus = "available"
	ShoppingListIsAssigned  ShoppingListStatus = "assigned"
	ShoppingListIsActive    ShoppingListStatus = "active"
	ShoppingListIsCompleted ShoppingListStatus = "completed"
	ShoppingListIsCancelled ShoppingListStatus = "canceled"
)

func NewShoppingListStatus(status string) ShoppingListStatus {
	switch status {
	case ShoppingListIsAvailable.String():
		return ShoppingListIsAvailable
	case ShoppingListIsAssigned.String():
		return ShoppingListIsAssigned
	case ShoppingListIsActive.String():
		return ShoppingListIsActive
	case ShoppingListIsCompleted.String():
		return ShoppingListIsCompleted
	case ShoppingListIsCancelled.String():
		return ShoppingListIsCancelled
	default:
		return ShoppingListUnknown
	}
}

func (s ShoppingListStatus) String() string {
	switch s {
	case ShoppingListIsAvailable, ShoppingListIsAssigned, ShoppingListIsActive, ShoppingListIsCompleted, ShoppingListIsCancelled:
		return string(s)
	default:
		return ""
	}
}
