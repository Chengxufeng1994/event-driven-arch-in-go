package command

import (
	"context"

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
	storeRepository repository.StoreRepository
}

func NewRebrandStoreHandler(storeRepository repository.StoreRepository) RebrandStoreHandler {
	return RebrandStoreHandler{
		storeRepository: storeRepository,
	}
}

func (h RebrandStoreHandler) RebrandStore(ctx context.Context, cmd RebrandStore) error {
	store, err := h.storeRepository.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if err := store.Rebrand(cmd.Name); err != nil {
		return err
	}

	return h.storeRepository.Save(ctx, store)
}
