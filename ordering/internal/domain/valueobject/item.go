package valueobject

type Item struct {
	ProductID   string
	StoreID     string
	StoreName   string
	ProductName string
	Price       float64
	Quantity    int
}

func NewItem(productID, storeID, storeName, productName string, price float64, quantity int) Item {
	return Item{
		ProductID:   productID,
		StoreID:     storeID,
		StoreName:   storeName,
		ProductName: productName,
		Price:       price,
		Quantity:    quantity,
	}
}
