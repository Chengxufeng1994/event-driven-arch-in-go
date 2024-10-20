package repository

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
)

type ProductRepository interface {
	FindProduct(ctx context.Context, id string) (*aggregate.ProductAgg, error)
	AddProduct(ctx context.Context, product *aggregate.ProductAgg) error
	RemoveProduct(ctx context.Context, id string) error
	GetCatalog(ctx context.Context, storeID string) ([]*aggregate.ProductAgg, error)
}
