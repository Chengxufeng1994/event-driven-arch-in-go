package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/repository"
)

type DecreaseProductPrice struct {
	ID    string
	Price float64
}

func NewDecreaseProductPrice(id string, price float64) DecreaseProductPrice {
	return DecreaseProductPrice{
		ID:    id,
		Price: price,
	}
}

type DecreaseProductPriceHandler struct {
	productRepository repository.ProductRepository
}

func NewDecreaseProductPriceHandler(productRepository repository.ProductRepository) DecreaseProductPriceHandler {
	return DecreaseProductPriceHandler{
		productRepository: productRepository,
	}
}

func (h DecreaseProductPriceHandler) DecreaseProductPrice(ctx context.Context, cmd DecreaseProductPrice) error {
	product, err := h.productRepository.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if err = product.DecreasePrice(cmd.Price); err != nil {
		return err
	}

	return h.productRepository.Save(ctx, product)
}
