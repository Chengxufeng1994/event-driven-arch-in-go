package query

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/repository"
)

type GetParticipatingStores struct{}

type GetParticipatingStoresHandler struct {
	mall repository.MallRepository
}

func NewGetParticipatingStoresHandler(mall repository.MallRepository) GetParticipatingStoresHandler {
	return GetParticipatingStoresHandler{mall: mall}
}

func (h GetParticipatingStoresHandler) GetParticipatingStores(ctx context.Context, _ GetParticipatingStores) ([]*aggregate.MallStore, error) {
	return h.mall.AllParticipating(ctx)
}
