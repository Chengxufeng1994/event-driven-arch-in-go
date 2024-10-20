package valueobject

type Item struct {
	ProductName string
	Quantity    int
}

func NewItems() Items {
	return make(map[string]*Item)
}

func NewItem(productName string, quantity int) Item {
	return Item{
		ProductName: productName,
		Quantity:    quantity,
	}
}

type Items map[string]*Item
