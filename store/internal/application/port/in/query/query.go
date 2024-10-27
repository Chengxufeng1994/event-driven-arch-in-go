package query

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
)

type Queries interface {
	GetStore(ctx context.Context, query GetStore) (*aggregate.Store, error)
	GetStores(ctx context.Context, query GetStores) ([]*aggregate.Store, error)
	GetParticipatingStores(ctx context.Context, query GetParticipatingStores) ([]*aggregate.Store, error)
	GetCatalog(ctx context.Context, query GetCatalog) ([]*aggregate.Product, error)
	GetProduct(ctx context.Context, query GetProduct) (*aggregate.Product, error)
}
