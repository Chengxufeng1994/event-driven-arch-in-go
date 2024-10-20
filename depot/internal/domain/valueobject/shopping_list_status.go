package valueobject

type ShoppingListStatus string

const (
	ShoppingListUnknown   ShoppingListStatus = ""
	ShoppingListAvailable ShoppingListStatus = "available"
	ShoppingListAssigned  ShoppingListStatus = "assigned"
	ShoppingListActive    ShoppingListStatus = "active"
	ShoppingListCompleted ShoppingListStatus = "completed"
	ShoppingListCancelled ShoppingListStatus = "canceled"
)

func NewShoppingListStatus(status string) ShoppingListStatus {
	switch status {
	case ShoppingListAvailable.String():
		return ShoppingListAvailable
	case ShoppingListAssigned.String():
		return ShoppingListAssigned
	case ShoppingListActive.String():
		return ShoppingListActive
	case ShoppingListCompleted.String():
		return ShoppingListCompleted
	case ShoppingListCancelled.String():
		return ShoppingListCancelled
	default:
		return ShoppingListUnknown
	}
}

func (s ShoppingListStatus) String() string {
	switch s {
	case ShoppingListAvailable, ShoppingListAssigned, ShoppingListActive, ShoppingListCompleted, ShoppingListCancelled:
		return string(s)
	default:
		return ""
	}
}
