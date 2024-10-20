package query

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/repository"
)

type GetCatalog struct {
	StoreID string
}

type GetCatalogHandler struct {
	products repository.ProductRepository
}

func NewGetCatalogHandler(products repository.ProductRepository) GetCatalogHandler {
	return GetCatalogHandler{products: products}
}

func (h GetCatalogHandler) GetCatalog(ctx context.Context, query GetCatalog) ([]*aggregate.ProductAgg, error) {
	return h.products.GetCatalog(ctx, query.StoreID)
}
