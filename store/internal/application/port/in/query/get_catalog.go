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
	catalog repository.CatalogRepository
}

func NewGetCatalogHandler(catalog repository.CatalogRepository) GetCatalogHandler {
	return GetCatalogHandler{catalog: catalog}
}

func (h GetCatalogHandler) GetCatalog(ctx context.Context, query GetCatalog) ([]*aggregate.CatalogProduct, error) {
	return h.catalog.GetCatalog(ctx, query.StoreID)
}
