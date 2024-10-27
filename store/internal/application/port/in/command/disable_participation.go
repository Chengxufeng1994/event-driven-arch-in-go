package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/repository"
	"github.com/stackus/errors"
)

type DisableParticipation struct {
	ID string
}

type DisableParticipationHandler struct {
	storeRepository repository.StoreRepository
}

func NewDisableParticipationHandler(
	stores repository.StoreRepository,
) DisableParticipationHandler {
	return DisableParticipationHandler{
		storeRepository: stores,
	}
}

func (h DisableParticipationHandler) DisableParticipation(ctx context.Context, cmd DisableParticipation) error {
	store, err := h.storeRepository.Load(ctx, cmd.ID)
	if err != nil {
		return errors.Wrap(err, "finding store")
	}

	if err := store.DisableParticipation(); err != nil {
		return errors.Wrap(err, "disabling participation")
	}

	if err := h.storeRepository.Save(ctx, store); err != nil {
		return errors.Wrap(err, "saving store")
	}

	return nil
}
