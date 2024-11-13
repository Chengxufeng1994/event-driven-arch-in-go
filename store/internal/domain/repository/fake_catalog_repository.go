package repository

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
)

type FakeCatalogRepository struct {
	products map[string]*aggregate.CatalogProduct
}

var _ CatalogRepository = (*FakeCatalogRepository)(nil)

func NewFakeCatalogRepository() *FakeCatalogRepository {
	return &FakeCatalogRepository{
		products: map[string]*aggregate.CatalogProduct{},
	}
}

func (r *FakeCatalogRepository) AddProduct(ctx context.Context, productID, storeID, name, description, sku string, price float64) error {
	// TODO implement me
	panic("implement me")
}

func (r *FakeCatalogRepository) Rebrand(ctx context.Context, productID, name, description string) error {
	// TODO implement me
	panic("implement me")
}

func (r *FakeCatalogRepository) UpdatePrice(ctx context.Context, productID string, delta float64) error {
	// TODO implement me
	panic("implement me")
}

func (r *FakeCatalogRepository) RemoveProduct(ctx context.Context, productID string) error {
	// TODO implement me
	panic("implement me")
}

func (r *FakeCatalogRepository) Find(ctx context.Context, productID string) (*aggregate.CatalogProduct, error) {
	// TODO implement me
	panic("implement me")
}

func (r *FakeCatalogRepository) GetCatalog(ctx context.Context, storeID string) ([]*aggregate.CatalogProduct, error) {
	// TODO implement me
	panic("implement me")
}
