package repository

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
)

type FakeMallRepository struct {
	stores map[string]*aggregate.MallStore
}

var _ MallRepository = (*FakeMallRepository)(nil)

func NewFakeMallRepository() *FakeMallRepository {
	return &FakeMallRepository{
		stores: map[string]*aggregate.MallStore{},
	}
}

func (r *FakeMallRepository) AddStore(ctx context.Context, storeID, name, location string) error {
	// TODO implement me
	panic("implement me")
}

func (r *FakeMallRepository) SetStoreParticipation(ctx context.Context, storeID string, participating bool) error {
	// TODO implement me
	panic("implement me")
}

func (r *FakeMallRepository) RenameStore(ctx context.Context, storeID, name string) error {
	// TODO implement me
	panic("implement me")
}

func (r *FakeMallRepository) Find(ctx context.Context, storeID string) (*aggregate.MallStore, error) {
	// TODO implement me
	panic("implement me")
}

func (r *FakeMallRepository) All(ctx context.Context) ([]*aggregate.MallStore, error) {
	// TODO implement me
	panic("implement me")
}

func (r *FakeMallRepository) AllParticipating(ctx context.Context) ([]*aggregate.MallStore, error) {
	// TODO implement me
	panic("implement me")
}
