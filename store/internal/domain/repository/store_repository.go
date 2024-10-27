package repository

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
)

type StoreRepository interface {
	Load(ctx context.Context, storeID string) (*aggregate.Store, error)
	Save(ctx context.Context, store *aggregate.Store) error
}
