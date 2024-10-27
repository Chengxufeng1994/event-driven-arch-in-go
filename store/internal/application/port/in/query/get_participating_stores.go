package query

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/repository"
)

type GetParticipatingStores struct{}

type GetParticipatingStoresHandler struct {
	participatingStores repository.ParticipatingStoreRepository
}

func NewGetParticipatingStoresHandler(participatingStores repository.ParticipatingStoreRepository) GetParticipatingStoresHandler {
	return GetParticipatingStoresHandler{participatingStores: participatingStores}
}

func (h GetParticipatingStoresHandler) GetParticipatingStores(ctx context.Context, _ GetParticipatingStores) ([]*aggregate.Store, error) {
	return h.participatingStores.FindAll(ctx)
}
