package query

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
)

type Queries interface {
	GetStore(ctx context.Context, query GetStore) (*aggregate.StoreAgg, error)
	GetStores(ctx context.Context, query GetStores) ([]*aggregate.StoreAgg, error)
	GetParticipatingStores(ctx context.Context, query GetParticipatingStores) ([]*aggregate.StoreAgg, error)
	GetCatalog(ctx context.Context, query GetCatalog) ([]*aggregate.ProductAgg, error)
	GetProduct(ctx context.Context, query GetProduct) (*aggregate.ProductAgg, error)
}
