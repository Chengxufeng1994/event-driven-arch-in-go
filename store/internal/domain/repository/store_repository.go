package repository

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
)

type StoreRepository interface {
	Save(ctx context.Context, store *aggregate.Store) error
	Update(ctx context.Context, store *aggregate.Store) error
	Delete(ctx context.Context, storeID string) error
	Find(ctx context.Context, storeID string) (*aggregate.Store, error)
	FindAll(ctx context.Context) ([]*aggregate.Store, error)
}
