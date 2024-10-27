package query

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/repository"
)

type GetStores struct{}

type GetStoresHandler struct {
	mall repository.MallRepository
}

func NewGetStoresHandler(mall repository.MallRepository) GetStoresHandler {
	return GetStoresHandler{mall: mall}
}

func (h GetStoresHandler) GetStores(ctx context.Context, _ GetStores) ([]*aggregate.MallStore, error) {
	return h.mall.All(ctx)
}
