package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/repository"
	"github.com/stackus/errors"
)

type EnableParticipation struct {
	ID string
}

type EnableParticipationHandler struct {
	storeRepository      repository.StoreRepository
	domainEventPublisher ddd.EventPublisher
}

func NewEnableParticipationHandler(
	storeRepository repository.StoreRepository,
	domainEventPublisher ddd.EventPublisher,
) EnableParticipationHandler {
	return EnableParticipationHandler{
		storeRepository:      storeRepository,
		domainEventPublisher: domainEventPublisher,
	}
}

func (h EnableParticipationHandler) EnableParticipation(ctx context.Context, cmd EnableParticipation) error {
	store, err := h.storeRepository.Find(ctx, cmd.ID)
	if err != nil {
		return errors.Wrap(err, "finding store")
	}

	if err := store.EnableParticipation(); err != nil {
		return errors.Wrap(err, "enabling participation")
	}

	if err := h.storeRepository.Update(ctx, store); err != nil {
		return errors.Wrap(err, "updating store")
	}

	if err := h.domainEventPublisher.Publish(ctx, store.GetEvents()...); err != nil {
		return errors.Wrap(err, "publishing events")
	}

	return nil
}
