package repository

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/entity"
	"github.com/stackus/errors"
)

type FakeProductCacheRepository struct {
	products map[string]*entity.Product
}

var _ ProductCacheRepository = (*FakeProductCacheRepository)(nil)

func NewFakeProductCacheRepository() *FakeProductCacheRepository {
	return &FakeProductCacheRepository{products: map[string]*entity.Product{}}
}

func (r *FakeProductCacheRepository) Add(ctx context.Context, productID, storeID, name string, price float64) error {
	r.products[productID] = &entity.Product{
		ID:      productID,
		StoreID: storeID,
		Name:    name,
		Price:   price,
	}

	return nil
}

func (r *FakeProductCacheRepository) Rebrand(ctx context.Context, productID, name string) error {
	if product, exists := r.products[productID]; exists {
		product.Name = name
	}

	return nil
}

func (r *FakeProductCacheRepository) UpdatePrice(ctx context.Context, productID string, delta float64) error {
	if product, exists := r.products[productID]; exists {
		product.Price += delta
	}

	return nil
}

func (r *FakeProductCacheRepository) Remove(ctx context.Context, productID string) error {
	delete(r.products, productID)

	return nil
}

func (r *FakeProductCacheRepository) Find(ctx context.Context, productID string) (*entity.Product, error) {
	if product, exists := r.products[productID]; exists {
		return product, nil
	}

	return nil, errors.ErrNotFound.Msgf("product with id: `%s` does not exist", productID)
}

func (r *FakeProductCacheRepository) Reset(products ...*entity.Product) {
	r.products = make(map[string]*entity.Product)

	for _, product := range products {
		r.products[product.ID] = product
	}
}
