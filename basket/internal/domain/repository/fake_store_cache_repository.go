package repository

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/entity"
	"github.com/stackus/errors"
)

type FakeStoreCacheRepository struct {
	stores map[string]*entity.Store
}

var _ StoreCacheRepository = (*FakeStoreCacheRepository)(nil)

func NewFakeStoreCacheRepository() *FakeStoreCacheRepository {
	return &FakeStoreCacheRepository{stores: map[string]*entity.Store{}}
}

func (r *FakeStoreCacheRepository) Add(ctx context.Context, storeID, name string) error {
	r.stores[storeID] = &entity.Store{
		ID:   storeID,
		Name: name,
	}

	return nil
}

func (r *FakeStoreCacheRepository) Rename(ctx context.Context, storeID, name string) error {
	if store, exists := r.stores[storeID]; exists {
		store.Name = name
	}

	return nil
}

func (r *FakeStoreCacheRepository) Find(ctx context.Context, storeID string) (*entity.Store, error) {
	if store, exists := r.stores[storeID]; exists {
		return store, nil
	}

	return nil, errors.ErrNotFound.Msgf("store with id: `%s` does not exist", storeID)
}

func (r *FakeStoreCacheRepository) Reset(stores ...*entity.Store) {
	r.stores = make(map[string]*entity.Store)

	for _, store := range stores {
		r.stores[store.ID] = store
	}
}
