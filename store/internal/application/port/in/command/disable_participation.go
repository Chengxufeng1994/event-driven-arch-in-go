package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/repository"
)

type DisableParticipation struct {
	ID string
}

type DisableParticipationHandler struct {
	stores repository.StoreRepository
}

func NewDisableParticipationHandler(stores repository.StoreRepository) DisableParticipationHandler {
	return DisableParticipationHandler{stores: stores}
}

func (h DisableParticipationHandler) DisableParticipation(ctx context.Context, cmd DisableParticipation) error {
	store, err := h.stores.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = store.DisableParticipation()
	if err != nil {
		return err
	}

	return h.stores.Update(ctx, store)
}
