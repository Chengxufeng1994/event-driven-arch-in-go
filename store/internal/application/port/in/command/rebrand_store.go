package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/repository"
)

type RebrandStore struct {
	ID   string
	Name string
}

func NewRebrandStore(id, name string) RebrandStore {
	return RebrandStore{
		ID:   id,
		Name: name,
	}
}

type RebrandStoreHandler struct {
	stores    repository.StoreRepository
	publisher ddd.EventPublisher[ddd.Event]
}

func NewRebrandStoreHandler(stores repository.StoreRepository, publisher ddd.EventPublisher[ddd.Event]) RebrandStoreHandler {
	return RebrandStoreHandler{
		stores:    stores,
		publisher: publisher,
	}
}

func (h RebrandStoreHandler) RebrandStore(ctx context.Context, cmd RebrandStore) error {
	store, err := h.stores.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	evt, err := store.Rebrand(cmd.Name)
	if err != nil {
		return err
	}

	err = h.stores.Save(ctx, store)
	if err != nil {
		return err
	}

	return h.publisher.Publish(ctx, evt)
}
