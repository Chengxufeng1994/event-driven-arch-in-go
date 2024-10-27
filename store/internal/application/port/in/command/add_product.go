package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
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
	stores               repository.StoreRepository
	products             repository.ProductRepository
	domainEventPublisher ddd.EventPublisher
}

func NewAddProductHandler(
	stores repository.StoreRepository,
	products repository.ProductRepository,
	domainEventPublisher ddd.EventPublisher,
) AddProductHandler {
	return AddProductHandler{
		stores:               stores,
		products:             products,
		domainEventPublisher: domainEventPublisher,
	}
}

func (h AddProductHandler) AddProduct(ctx context.Context, cmd AddProduct) error {
	_, err := h.stores.Find(ctx, cmd.StoreID)
	if err != nil {
		return errors.Wrap(err, "fetching store")
	}

	product, err := aggregate.CreateProduct(cmd.ID, cmd.StoreID, cmd.Name, cmd.Description, cmd.SKU, cmd.Price)
	if err != nil {
		return errors.Wrap(err, "creating product")
	}

	if err := h.domainEventPublisher.Publish(ctx, product.GetEvents()...); err != nil {
		return errors.Wrap(err, "publishing events")
	}

	return errors.Wrap(h.products.Save(ctx, product), "error adding product")
}
