package query

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/repository"
)

type GetStores struct{}

type GetStoresHandler struct {
	stores repository.StoreRepository
}

func NewGetStoresHandler(stores repository.StoreRepository) GetStoresHandler {
	return GetStoresHandler{stores: stores}
}

func (h GetStoresHandler) GetStores(ctx context.Context, _ GetStores) ([]*aggregate.Store, error) {
	return h.stores.FindAll(ctx)
}
