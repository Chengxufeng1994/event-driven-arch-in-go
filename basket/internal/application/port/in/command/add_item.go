package command

import (
	"context"

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
	basketRepository  repository.BasketRepository
	productRepository repository.ProductRepository
	storeRepository   repository.StoreRepository
}

func NewAddItemHandler(
	basketRepository repository.BasketRepository,
	productRepository repository.ProductRepository,
	storeRepository repository.StoreRepository,
) AddItemHandler {
	return AddItemHandler{
		basketRepository:  basketRepository,
		productRepository: productRepository,
		storeRepository:   storeRepository,
	}
}

func (h AddItemHandler) AddItem(ctx context.Context, cmd AddItem) error {
	basket, err := h.basketRepository.Load(ctx, cmd.ID)
	if err != nil {
		return errors.Wrap(err, "add item command")
	}

	product, err := h.productRepository.Find(ctx, cmd.ProductID)
	if err != nil {
		return errors.Wrap(err, "finding product")
	}

	store, err := h.storeRepository.Find(ctx, product.StoreID)
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
