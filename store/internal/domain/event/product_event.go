package event

const (
	ProductAddedEvent          = "stores.ProductAdded"
	ProductRebrandedEvent      = "stores.ProductRebranded"
	ProductPriceIncreasedEvent = "stores.ProductPriceIncreased"
	ProductPriceDecreasedEvent = "stores.ProductPriceDecreased"
	ProductRemovedEvent        = "stores.ProductRemoved"
)

type ProductAdded struct {
	StoreID     string
	Name        string
	Description string
	SKU         string
	Price       float64
}

func NewProductAdded(storeID, name, description, sku string, price float64) *ProductAdded {
	return &ProductAdded{
		StoreID:     storeID,
		Name:        name,
		Description: description,
		SKU:         sku,
		Price:       price,
	}
}

func (ProductAdded) Key() string { return ProductAddedEvent }

type ProductRemoved struct{}

func NewProductRemoved() *ProductRemoved {
	return &ProductRemoved{}
}

func (ProductRemoved) Key() string { return ProductRemovedEvent }

type ProductRebranded struct {
	Name        string
	Description string
}

func NewProductRebranded(name, description string) *ProductRebranded {
	return &ProductRebranded{
		Name:        name,
		Description: description,
	}
}

func (ProductRebranded) Key() string { return ProductRebrandedEvent }

type ProductPriceChanged struct {
	Delta float64
}

func NewProductPriceChanged(delta float64) *ProductPriceChanged {
	return &ProductPriceChanged{
		Delta: delta,
	}
}
