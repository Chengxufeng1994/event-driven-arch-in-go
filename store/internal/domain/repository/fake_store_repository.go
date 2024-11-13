package repository

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
)

type FakeStoreRepository struct {
	stores map[string]*aggregate.Store
}

func NewFakeStoreRepository() *FakeStoreRepository {
	return &FakeStoreRepository{stores: map[string]*aggregate.Store{}}
}

var _ StoreRepository = (*FakeStoreRepository)(nil)

func (r *FakeStoreRepository) Load(ctx context.Context, storeID string) (*aggregate.Store, error) {
	if store, exists := r.stores[storeID]; exists {
		return store, nil
	}

	return aggregate.NewStore(storeID), nil
}

func (r *FakeStoreRepository) Save(ctx context.Context, store *aggregate.Store) error {
	for _, event := range store.Events() {
		if err := store.ApplyEvent(event); err != nil {
			return err
		}
	}

	r.stores[store.ID()] = store

	return nil
}

func (r *FakeStoreRepository) Reset(stores ...*aggregate.Store) {
	r.stores = make(map[string]*aggregate.Store)

	for _, store := range stores {
		r.stores[store.ID()] = store
	}
}
