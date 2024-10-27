package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/repository"
)

type RebrandProduct struct {
	ID          string
	Name        string
	Description string
}

func NewRebrandProduct(id, name, description string) RebrandProduct {
	return RebrandProduct{
		ID:          id,
		Name:        name,
		Description: description,
	}
}

type RebrandProductHandler struct {
	productRepository repository.ProductRepository
}

func NewRebrandProductHandler(productRepository repository.ProductRepository) RebrandProductHandler {
	return RebrandProductHandler{
		productRepository: productRepository,
	}
}

func (h RebrandProductHandler) RebrandProduct(ctx context.Context, cmd RebrandProduct) error {
	product, err := h.productRepository.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if err := product.Rebrand(cmd.Name, cmd.Description); err != nil {
		return err
	}

	return h.productRepository.Save(ctx, product)
}
