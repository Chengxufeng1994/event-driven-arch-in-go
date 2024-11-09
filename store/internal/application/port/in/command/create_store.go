package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
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
		stores repository.StoreRepository
	}
)

func NewCreateStoreHandler(stores repository.StoreRepository) CreateStoreHandler {
	return CreateStoreHandler{
		stores: stores,
	}
}

func (h CreateStoreHandler) CreateStore(ctx context.Context, cmd CreateStore) error {
	store, err := aggregate.CreateStore(cmd.ID, cmd.Name, cmd.Location)
	if err != nil {
		return errors.Wrap(err, "creating store")
	}

	if err = h.stores.Save(ctx, store); err != nil {
		return errors.Wrap(err, "saving store")
	}

	return nil
}
