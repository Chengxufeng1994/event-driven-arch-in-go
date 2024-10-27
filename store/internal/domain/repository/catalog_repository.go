package repository

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
)

type CatalogRepository interface {
	AddProduct(ctx context.Context, productID, storeID, name, description, sku string, price float64) error
	Rebrand(ctx context.Context, productID, name, description string) error
	UpdatePrice(ctx context.Context, productID string, delta float64) error
	RemoveProduct(ctx context.Context, productID string) error
	Find(ctx context.Context, productID string) (*aggregate.CatalogProduct, error)
	GetCatalog(ctx context.Context, storeID string) ([]*aggregate.CatalogProduct, error)
}
