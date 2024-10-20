package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/repository"
	"github.com/stackus/errors"
)

type AddProduct struct {
	ID          string
	StoreID     string
	Name        string
	Description string
	SKU         string
	Price       float64
}

type AddProductHandler struct {
	stores   repository.StoreRepository
	products repository.ProductRepository
}

func NewAddProductHandler(stores repository.StoreRepository, products repository.ProductRepository) AddProductHandler {
	return AddProductHandler{
		stores:   stores,
		products: products,
	}
}

func (h AddProductHandler) AddProduct(ctx context.Context, cmd AddProduct) error {
	_, err := h.stores.Find(ctx, cmd.StoreID)
	if err != nil {
		return errors.Wrap(err, "error adding product")
	}

	product, err := aggregate.CreateProduct(cmd.ID, cmd.StoreID, cmd.Name, cmd.Description, cmd.SKU, cmd.Price)
	if err != nil {
		return errors.Wrap(err, "error adding product")
	}

	return errors.Wrap(h.products.AddProduct(ctx, product), "error adding product")
}
