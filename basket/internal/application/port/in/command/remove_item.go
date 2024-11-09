package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/repository"
	"github.com/stackus/errors"
)

type RemoveItem struct {
	ID        string
	ProductID string
	Quantity  int
}

func NewRemoveItem(id, productID string, quantity int) RemoveItem {
	return RemoveItem{
		ID:        id,
		ProductID: productID,
		Quantity:  quantity,
	}
}

type RemoveItemHandler struct {
	basketRepository  repository.BasketRepository
	productRepository repository.ProductRepository
}

func NewRemoveItemHandler(basketRepository repository.BasketRepository, productRepository repository.ProductRepository) RemoveItemHandler {
	return RemoveItemHandler{
		basketRepository:  basketRepository,
		productRepository: productRepository,
	}
}

func (h RemoveItemHandler) RemoveItem(ctx context.Context, cmd RemoveItem) error {
	product, err := h.productRepository.Find(ctx, cmd.ProductID)
	if err != nil {
		return errors.Wrap(err, "fetching product")
	}

	basket, err := h.basketRepository.Load(ctx, cmd.ID)
	if err != nil {
		return errors.Wrap(err, "finding basket")
	}

	err = basket.RemoveItem(product, cmd.Quantity)
	if err != nil {
		return errors.Wrap(err, "remove item command")
	}

	if err := h.basketRepository.Save(ctx, basket); err != nil {
		return errors.Wrap(err, "save basket")
	}

	return nil
}
