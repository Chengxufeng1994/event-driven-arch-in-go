package entity

type Item struct {
	StoreID      string
	ProductID    string
	StoreName    string
	ProductName  string
	ProductPrice float64
	Quantity     int
}

func NewItem(storeID, productID, storeName, productName string, productPrice float64, quantity int) *Item {
	return &Item{
		StoreID:      storeID,
		ProductID:    productID,
		StoreName:    storeName,
		ProductName:  productName,
		ProductPrice: productPrice,
		Quantity:     quantity,
	}
}
