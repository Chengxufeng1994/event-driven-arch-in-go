package query

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/repository"
)

type GetProduct struct {
	ID string
}

type GetProductHandler struct {
	products repository.ProductRepository
}

func NewGetProductHandler(products repository.ProductRepository) GetProductHandler {
	return GetProductHandler{products: products}
}

func (h GetProductHandler) GetProduct(ctx context.Context, query GetProduct) (*aggregate.Product, error) {
	return h.products.Find(ctx, query.ID)
}
