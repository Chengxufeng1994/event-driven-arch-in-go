package repository

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
)

type MallRepository interface {
	AddStore(ctx context.Context, storeID, name, location string) error
	SetStoreParticipation(ctx context.Context, storeID string, participating bool) error
	RenameStore(ctx context.Context, storeID, name string) error
	Find(ctx context.Context, storeID string) (*aggregate.MallStore, error)
	All(ctx context.Context) ([]*aggregate.MallStore, error)
	AllParticipating(ctx context.Context) ([]*aggregate.MallStore, error)
}
