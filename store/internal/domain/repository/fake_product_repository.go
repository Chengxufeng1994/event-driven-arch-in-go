package repository

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
)

type FakeProductRepository struct {
	products map[string]*aggregate.Product
}

func NewFakeProductRepository() *FakeProductRepository {
	return &FakeProductRepository{products: map[string]*aggregate.Product{}}
}

var _ ProductRepository = (*FakeProductRepository)(nil)

func (r *FakeProductRepository) Load(ctx context.Context, productID string) (*aggregate.Product, error) {
	if product, exists := r.products[productID]; exists {
		return product, nil
	}

	return aggregate.NewProduct(productID), nil
}

func (r *FakeProductRepository) Save(ctx context.Context, product *aggregate.Product) error {
	for _, event := range product.Events() {
		if err := product.ApplyEvent(event); err != nil {
			return err
		}
	}

	r.products[product.ID()] = product

	return nil
}

func (r *FakeProductRepository) Reset(products ...*aggregate.Product) {
	r.products = make(map[string]*aggregate.Product)

	for _, product := range products {
		r.products[product.ID()] = product
	}
}
