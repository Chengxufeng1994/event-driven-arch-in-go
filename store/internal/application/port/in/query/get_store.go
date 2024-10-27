package query

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/repository"
)

type GetStore struct {
	ID string
}

type GetStoreHandler struct {
	stores repository.StoreRepository
}

func NewGetStoreHandler(stores repository.StoreRepository) GetStoreHandler {
	return GetStoreHandler{stores: stores}
}

func (h GetStoreHandler) GetStore(ctx context.Context, query GetStore) (*aggregate.Store, error) {
	return h.stores.Find(ctx, query.ID)
}
