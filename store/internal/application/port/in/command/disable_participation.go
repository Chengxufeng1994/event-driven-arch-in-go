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
	stores    repository.StoreRepository
	publisher ddd.EventPublisher[ddd.Event]
}

func NewDisableParticipationHandler(stores repository.StoreRepository, publisher ddd.EventPublisher[ddd.Event]) DisableParticipationHandler {
	return DisableParticipationHandler{
		stores:    stores,
		publisher: publisher,
	}
}

func (h DisableParticipationHandler) DisableParticipation(ctx context.Context, cmd DisableParticipation) error {
	store, err := h.stores.Load(ctx, cmd.ID)
	if err != nil {
		return errors.Wrap(err, "finding store")
	}

	event, err := store.DisableParticipation()
	if err != nil {
		return errors.Wrap(err, "disabling participation")
	}

	if err := h.stores.Save(ctx, store); err != nil {
		return errors.Wrap(err, "saving store")
	}

	return h.publisher.Publish(ctx, event)
}
