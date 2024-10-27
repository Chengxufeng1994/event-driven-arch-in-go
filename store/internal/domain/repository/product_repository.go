package repository

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
)

type ProductRepository interface {
	Save(ctx context.Context, product *aggregate.Product) error
	Delete(ctx context.Context, id string) error
	Find(ctx context.Context, id string) (*aggregate.Product, error)
	FindCatalog(ctx context.Context, storeID string) ([]*aggregate.Product, error)
}
