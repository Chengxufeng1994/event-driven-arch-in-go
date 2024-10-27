package event

const (
	ProductAddedEvent          = "stores.ProductAdded"
	ProductReBrandedEvent      = "stores.ProductReBranded"
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

func (ProductAdded) EventName() string { return ProductAddedEvent }

type ProductRemoved struct{}

func NewProductRemoved() *ProductRemoved {
	return &ProductRemoved{}
}

func (ProductRemoved) EventName() string { return ProductRemovedEvent }
