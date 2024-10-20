package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/repository"
)

type EnableParticipation struct {
	ID string
}

type EnableParticipationHandler struct {
	stores repository.StoreRepository
}

func NewEnableParticipationHandler(stores repository.StoreRepository) EnableParticipationHandler {
	return EnableParticipationHandler{stores: stores}
}

func (h EnableParticipationHandler) EnableParticipation(ctx context.Context, cmd EnableParticipation) error {
	store, err := h.stores.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = store.EnableParticipation()
	if err != nil {
		return err
	}

	return h.stores.Update(ctx, store)
}
