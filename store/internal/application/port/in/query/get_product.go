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
	catalog repository.CatalogRepository
}

func NewGetProductHandler(catalog repository.CatalogRepository) GetProductHandler {
	return GetProductHandler{catalog: catalog}
}

func (h GetProductHandler) GetProduct(ctx context.Context, query GetProduct) (*aggregate.CatalogProduct, error) {
	return h.catalog.Find(ctx, query.ID)
}
