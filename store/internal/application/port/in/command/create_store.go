package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/repository"
	"github.com/stackus/errors"
)

type (
	CreateStore struct {
		ID       string
		Name     string
		Location string
	}

	CreateStoreHandler struct {
		stores    repository.StoreRepository
		publisher ddd.EventPublisher[ddd.Event]
	}
)

func NewCreateStoreHandler(stores repository.StoreRepository, publisher ddd.EventPublisher[ddd.Event]) CreateStoreHandler {
	return CreateStoreHandler{
		stores:    stores,
		publisher: publisher,
	}
}

func (h CreateStoreHandler) CreateStore(ctx context.Context, cmd CreateStore) error {
	store, err := h.stores.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	event, err := store.InitStore(cmd.Name, cmd.Location)
	if err != nil {
		return errors.Wrap(err, "initializing store")
	}

	if err = h.stores.Save(ctx, store); err != nil {
		return errors.Wrap(err, "saving store")
	}

	return h.publisher.Publish(ctx, event)
}
