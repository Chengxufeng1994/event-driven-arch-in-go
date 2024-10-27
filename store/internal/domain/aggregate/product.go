package aggregate

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/event"
	"github.com/stackus/errors"
)

var (
	ErrProductNameIsBlank     = errors.Wrap(errors.ErrBadRequest, "the product name cannot be blank")
	ErrProductPriceIsNegative = errors.Wrap(errors.ErrBadRequest, "the product price cannot be negative")
)

type Product struct {
	ddd.AggregateBase
	StoreID     string
	Name        string
	Description string
	SKU         string
	Price       float64
}

func CreateProduct(id, storeID, name, description, sku string, price float64) (*Product, error) {
	if name == "" {
		return nil, ErrProductNameIsBlank
	}

	if price < 0 {
		return nil, ErrProductPriceIsNegative
	}

	product := &Product{
		AggregateBase: ddd.NewAggregateBase(id),
		StoreID:       storeID,
		Name:          name,
		Description:   description,
		SKU:           sku,
		Price:         price,
	}

	product.AddEvent(event.NewProductAdded(
		storeID,
		name,
		description,
		sku,
		price,
	))

	return product, nil
}

func (product *Product) Remove() error {
	product.AddEvent(event.NewProductRemoved())
	return nil
}
