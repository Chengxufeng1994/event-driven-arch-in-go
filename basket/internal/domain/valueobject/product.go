package valueobject

type Product struct {
	ID      string
	StoreID string
	Name    string
	Price   float64
}

func NewProduct(id, storeID, name string, price float64) Product {
	return Product{
		ID:      id,
		StoreID: storeID,
		Name:    name,
		Price:   price,
	}
}
