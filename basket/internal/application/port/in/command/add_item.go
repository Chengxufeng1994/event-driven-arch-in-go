package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/port/out/client"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/repository"
	"github.com/stackus/errors"
)

type AddItem struct {
	ID        string
	ProductID string
	Quantity  int
}

func NewAddItem(id, productID string, quantity int) AddItem {
	return AddItem{
		ID:        id,
		ProductID: productID,
		Quantity:  quantity,
	}
}

type AddItemHandler struct {
	basketRepository repository.BasketRepository
	productClient    client.ProductClient
	storeClient      client.StoreClient
}

func NewAddItemHandler(
	basketRepository repository.BasketRepository,
	productClient client.ProductClient,
	storeClient client.StoreClient,
) AddItemHandler {
	return AddItemHandler{
		basketRepository: basketRepository,
		productClient:    productClient,
		storeClient:      storeClient,
	}
}

func (h AddItemHandler) AddItem(ctx context.Context, cmd AddItem) error {
	basket, err := h.basketRepository.Load(ctx, cmd.ID)
	if err != nil {
		return errors.Wrap(err, "add item command")
	}

	product, err := h.productClient.Find(ctx, cmd.ProductID)
	if err != nil {
		return errors.Wrap(err, "finding product")
	}

	store, err := h.storeClient.Find(ctx, product.StoreID)
	if err != nil {
		return errors.Wrap(err, "finding store")
	}

	err = basket.AddItem(store, product, cmd.Quantity)
	if err != nil {
		return errors.Wrap(err, "adding item to basket")
	}

	if err := h.basketRepository.Save(ctx, basket); err != nil {
		return errors.Wrap(err, "saving basket")
	}

	return nil
}
