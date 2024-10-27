package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/repository"
	"github.com/stackus/errors"
)

type DisableParticipation struct {
	ID string
}

type DisableParticipationHandler struct {
	storeRepository      repository.StoreRepository
	domainEventPublisher ddd.EventPublisher
}

func NewDisableParticipationHandler(
	stores repository.StoreRepository,
	domainEventPublisher ddd.EventPublisher,
) DisableParticipationHandler {
	return DisableParticipationHandler{
		storeRepository:      stores,
		domainEventPublisher: domainEventPublisher,
	}
}

func (h DisableParticipationHandler) DisableParticipation(ctx context.Context, cmd DisableParticipation) error {
	store, err := h.storeRepository.Find(ctx, cmd.ID)
	if err != nil {
		return errors.Wrap(err, "finding store")
	}

	if err := store.DisableParticipation(); err != nil {
		return errors.Wrap(err, "disabling participation")
	}

	if err := h.storeRepository.Update(ctx, store); err != nil {
		return errors.Wrap(err, "updating store")
	}

	if err := h.domainEventPublisher.Publish(ctx, store.GetEvents()...); err != nil {
		return errors.Wrap(err, "publishing events")
	}

	return nil
}
