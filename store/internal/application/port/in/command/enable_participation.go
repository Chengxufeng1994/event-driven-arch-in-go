package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/repository"
	"github.com/stackus/errors"
)

type EnableParticipation struct {
	ID string
}

type EnableParticipationHandler struct {
	storeRepository repository.StoreRepository
}

func NewEnableParticipationHandler(
	storeRepository repository.StoreRepository,
) EnableParticipationHandler {
	return EnableParticipationHandler{
		storeRepository: storeRepository,
	}
}

func (h EnableParticipationHandler) EnableParticipation(ctx context.Context, cmd EnableParticipation) error {
	store, err := h.storeRepository.Load(ctx, cmd.ID)
	if err != nil {
		return errors.Wrap(err, "finding store")
	}

	if err := store.EnableParticipation(); err != nil {
		return errors.Wrap(err, "enabling participation")
	}

	if err := h.storeRepository.Save(ctx, store); err != nil {
		return errors.Wrap(err, "saving store")
	}

	return nil
}
