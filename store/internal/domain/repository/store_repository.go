package repository

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
)

type StoreRepository interface {
	Save(ctx context.Context, store *aggregate.StoreAgg) error
	Update(ctx context.Context, store *aggregate.StoreAgg) error
	Delete(ctx context.Context, storeID string) error
	Find(ctx context.Context, storeID string) (*aggregate.StoreAgg, error)
	FindAll(ctx context.Context) ([]*aggregate.StoreAgg, error)
}
