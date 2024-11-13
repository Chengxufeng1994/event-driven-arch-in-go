package repository

import (
	"context"

	"github.com/stackus/errors"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/entity"
)

type FakeStoreCacheRepository struct {
	stores map[string]*entity.Store
}

var _ StoreCacheRepository = (*FakeStoreCacheRepository)(nil)

func NewFakeStoreCacheRepository() *FakeStoreCacheRepository {
	return &FakeStoreCacheRepository{stores: make(map[string]*entity.Store)}
}

// Add implements StoreCacheRepository.
func (r *FakeStoreCacheRepository) Add(ctx context.Context, storeID string, name string, location string) error {
	r.stores[storeID] = &entity.Store{
		ID:       storeID,
		Name:     name,
		Location: location,
	}

	return nil
}

// Find implements StoreCacheRepository.
func (f *FakeStoreCacheRepository) Find(ctx context.Context, storeID string) (*entity.Store, error) {
	if store, exists := f.stores[storeID]; exists {
		return store, nil
	}

	return nil, errors.ErrNotFound.Msgf("store with id: `%s` does not exist", storeID)
}

// Rename implements StoreCacheRepository.
func (f *FakeStoreCacheRepository) Rename(ctx context.Context, storeID string, name string) error {
	if store, exists := f.stores[storeID]; exists {
		store.Name = name
	}

	return nil
}

func (r *FakeStoreCacheRepository) Reset(stores ...*entity.Store) {
	r.stores = make(map[string]*entity.Store)

	for _, store := range stores {
		r.stores[store.ID] = store
	}
}
