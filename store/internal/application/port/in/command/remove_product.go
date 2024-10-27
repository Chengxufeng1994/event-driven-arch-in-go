package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/repository"
	"github.com/stackus/errors"
)

type RemoveProduct struct {
	ID string
}

type RemoveProductHandler struct {
	productRepository repository.ProductRepository
}

func NewRemoveProductHandler(
	productRepository repository.ProductRepository,
) RemoveProductHandler {
	return RemoveProductHandler{
		productRepository: productRepository,
	}
}

func (h RemoveProductHandler) RemoveProduct(ctx context.Context, cmd RemoveProduct) error {
	product, err := h.productRepository.Load(ctx, cmd.ID)
	if err != nil {
		return errors.Wrap(err, "remove product command")
	}

	if err := product.Remove(); err != nil {
		return errors.Wrap(err, "remove product command")
	}

	if err := h.productRepository.Save(ctx, product); err != nil {
		return errors.Wrap(err, "saving product")
	}

	return nil
}
