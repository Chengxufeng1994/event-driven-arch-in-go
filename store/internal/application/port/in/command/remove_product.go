package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/repository"
)

type RemoveProduct struct {
	ID string
}

type RemoveProductHandler struct {
	stores   repository.StoreRepository
	products repository.ProductRepository
}

func NewRemoveProductHandler(stores repository.StoreRepository, products repository.ProductRepository) RemoveProductHandler {
	return RemoveProductHandler{
		stores:   stores,
		products: products,
	}
}

func (h RemoveProductHandler) RemoveProduct(ctx context.Context, cmd RemoveProduct) error {
	return h.products.RemoveProduct(ctx, cmd.ID)
}
