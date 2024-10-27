package query

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
)

type Queries interface {
	GetStore(ctx context.Context, query GetStore) (*aggregate.MallStore, error)
	GetStores(ctx context.Context, query GetStores) ([]*aggregate.MallStore, error)
	GetParticipatingStores(ctx context.Context, query GetParticipatingStores) ([]*aggregate.MallStore, error)
	GetCatalog(ctx context.Context, query GetCatalog) ([]*aggregate.CatalogProduct, error)
	GetProduct(ctx context.Context, query GetProduct) (*aggregate.CatalogProduct, error)
}
