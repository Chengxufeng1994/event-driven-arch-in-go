package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/repository"
)

type IncreaseProductPrice struct {
	ID    string
	Price float64
}

func NewIncreaseProductPrice(id string, price float64) IncreaseProductPrice {
	return IncreaseProductPrice{
		ID:    id,
		Price: price,
	}
}

type IncreaseProductPriceHandler struct {
	productRepository repository.ProductRepository
}

func NewIncreaseProductPriceHandler(productRepository repository.ProductRepository) IncreaseProductPriceHandler {
	return IncreaseProductPriceHandler{
		productRepository: productRepository,
	}
}

func (h IncreaseProductPriceHandler) IncreaseProductPrice(ctx context.Context, cmd IncreaseProductPrice) error {
	product, err := h.productRepository.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if err = product.IncreasePrice(cmd.Price); err != nil {
		return err
	}

	return h.productRepository.Save(ctx, product)
}
