package aggregate

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/es"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/event"
	"github.com/stackus/errors"
)

const ProductAggregate = "stores.Product"

var (
	ErrProductNameIsBlank     = errors.Wrap(errors.ErrBadRequest, "the product name cannot be blank")
	ErrProductPriceIsNegative = errors.Wrap(errors.ErrBadRequest, "the product price cannot be negative")
	ErrNotAPriceIncrease      = errors.Wrap(errors.ErrBadRequest, "the price change would be a decrease")
	ErrNotAPriceDecrease      = errors.Wrap(errors.ErrBadRequest, "the price change would be an increase")
)

type Product struct {
	es.Aggregate
	StoreID     string
	Name        string
	Description string
	SKU         string
	Price       float64
}

var _ interface {
	es.EventApplier
	es.Snapshotter
} = (*Product)(nil)

func NewProduct(id string) *Product {
	return &Product{
		Aggregate: es.NewAggregate(id, ProductAggregate),
	}
}

func (p *Product) InitProduct(id, storeID, name, description, sku string, price float64) (ddd.Event, error) {
	if name == "" {
		return nil, ErrProductNameIsBlank
	}

	if price < 0 {
		return nil, ErrProductPriceIsNegative
	}

	p.AddEvent(event.ProductAddedEvent, &event.ProductAdded{
		StoreID:     storeID,
		Name:        name,
		Description: description,
		SKU:         sku,
		Price:       price,
	})

	return ddd.NewEvent(event.ProductAddedEvent, p), nil
}

// Key implements registry.Registerable
func (Product) Key() string { return ProductAggregate }

func (p *Product) Rebrand(name, description string) (ddd.Event, error) {
	p.AddEvent(event.ProductRebrandedEvent, &event.ProductRebranded{
		Name:        name,
		Description: description,
	})

	return ddd.NewEvent(event.ProductRebrandedEvent, p), nil
}

func (p *Product) IncreasePrice(price float64) (ddd.Event, error) {
	if price < p.Price {
		return nil, ErrNotAPriceIncrease
	}

	delta := price - p.Price
	p.AddEvent(event.ProductPriceIncreasedEvent, &event.ProductPriceChanged{
		Delta: delta,
	})

	// FIXME: ProductID to Product
	return ddd.NewEvent(event.ProductPriceIncreasedEvent, event.ProductPriceDelta{
		ProductID: p.ID(),
		Delta:     delta,
	}), nil
}

func (p *Product) DecreasePrice(price float64) (ddd.Event, error) {
	if price > p.Price {
		return nil, ErrNotAPriceDecrease
	}

	delta := price - p.Price
	p.AddEvent(event.ProductPriceDecreasedEvent, &event.ProductPriceChanged{
		Delta: delta,
	})

	// FIXME: ProductID to Product
	return ddd.NewEvent(event.ProductPriceDecreasedEvent, event.ProductPriceDelta{
		ProductID: p.ID(),
		Delta:     delta,
	}), nil
}

func (p *Product) Remove() (ddd.Event, error) {
	p.AddEvent(event.ProductRemovedEvent, &event.ProductRemoved{})

	return ddd.NewEvent(event.ProductRemovedEvent, p), nil
}

func (p *Product) ApplyEvent(e ddd.Event) error {
	switch payload := e.Payload().(type) {
	case *event.ProductAdded:
		p.StoreID = payload.StoreID
		p.Name = payload.Name
		p.Description = payload.Description
		p.SKU = payload.SKU
		p.Price = payload.Price

	case *event.ProductRebranded:
		p.Name = payload.Name
		p.Description = payload.Description

	case *event.ProductPriceChanged:
		p.Price += payload.Delta

	case *event.ProductRemoved:
		// noop

	default:
		return errors.ErrInternal.Msgf("%T received the event %s with unexpected payload %T", p, e.EventName(), payload)
	}

	return nil
}

func (p *Product) ApplySnapshot(snapshot es.Snapshot) error {
	switch ss := snapshot.(type) {
	case *ProductV1:
		p.StoreID = ss.StoreID
		p.Name = ss.Name
		p.Description = ss.Description
		p.SKU = ss.SKU
		p.Price = ss.Price

	default:
		return errors.ErrInternal.Msgf("%T received the unexpected snapshot %T", p, snapshot)
	}

	return nil
}

func (p Product) ToSnapshot() es.Snapshot {
	return ProductV1{
		StoreID:     p.StoreID,
		Name:        p.Name,
		Description: p.Description,
		SKU:         p.SKU,
		Price:       p.Price,
	}
}
