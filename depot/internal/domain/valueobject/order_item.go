package valueobject

type OrderItem struct {
	StoreID   string
	ProductID string
	Quantity  int
}

func NewOrderItem(storeID, productID string, quantity int) OrderItem {
	return OrderItem{
		StoreID:   storeID,
		ProductID: productID,
		Quantity:  quantity,
	}
}
