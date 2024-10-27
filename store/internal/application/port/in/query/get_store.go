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
	mall repository.MallRepository
}

func NewGetStoreHandler(mall repository.MallRepository) GetStoreHandler {
	return GetStoreHandler{mall: mall}
}

func (h GetStoreHandler) GetStore(ctx context.Context, query GetStore) (*aggregate.MallStore, error) {
	return h.mall.Find(ctx, query.ID)
}
