package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
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
	products  repository.ProductRepository
	publisher ddd.EventPublisher[ddd.Event]
}

func NewAddProductHandler(
	products repository.ProductRepository,
	publisher ddd.EventPublisher[ddd.Event],
) AddProductHandler {
	return AddProductHandler{
		products:  products,
		publisher: publisher,
	}
}

func (h AddProductHandler) AddProduct(ctx context.Context, cmd AddProduct) error {
	product, err := h.products.Load(ctx, cmd.ID)
	if err != nil {
		return errors.Wrap(err, "error adding product")
	}

	event, err := product.InitProduct(cmd.ID, cmd.StoreID, cmd.Name, cmd.Description, cmd.SKU, cmd.Price)
	if err != nil {
		return errors.Wrap(err, "initializing product")
	}

	err = h.products.Save(ctx, product)
	if err != nil {
		return errors.Wrap(err, "error adding product")
	}

	return errors.Wrap(h.publisher.Publish(ctx, event), "publishing domain event")
}
