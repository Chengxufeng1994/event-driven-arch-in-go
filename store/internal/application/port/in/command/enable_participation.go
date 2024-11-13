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
	stores    repository.StoreRepository
	publisher ddd.EventPublisher[ddd.Event]
}

func NewEnableParticipationHandler(stores repository.StoreRepository, publisher ddd.EventPublisher[ddd.Event]) EnableParticipationHandler {
	return EnableParticipationHandler{
		stores:    stores,
		publisher: publisher,
	}
}

func (h EnableParticipationHandler) EnableParticipation(ctx context.Context, cmd EnableParticipation) error {
	store, err := h.stores.Load(ctx, cmd.ID)
	if err != nil {
		return errors.Wrap(err, "finding store")
	}

	event, err := store.EnableParticipation()
	if err != nil {
		return errors.Wrap(err, "enabling participation")
	}

	if err := h.stores.Save(ctx, store); err != nil {
		return errors.Wrap(err, "saving store")
	}

	return h.publisher.Publish(ctx, event)
}
