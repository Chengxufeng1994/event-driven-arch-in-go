package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/repository"
	"github.com/stackus/errors"
)

type RemoveProduct struct {
	ID string
}

type RemoveProductHandler struct {
	products  repository.ProductRepository
	publisher ddd.EventPublisher[ddd.Event]
}

func NewRemoveProductHandler(products repository.ProductRepository, publisher ddd.EventPublisher[ddd.Event]) RemoveProductHandler {
	return RemoveProductHandler{
		products:  products,
		publisher: publisher,
	}
}

func (h RemoveProductHandler) RemoveProduct(ctx context.Context, cmd RemoveProduct) error {
	product, err := h.products.Load(ctx, cmd.ID)
	if err != nil {
		return errors.Wrap(err, "loading product")
	}

	event, err := product.Remove()
	if err != nil {
		return errors.Wrap(err, "remove product")
	}

	if err := h.products.Save(ctx, product); err != nil {
		return errors.Wrap(err, "saving product")
	}

	return h.publisher.Publish(ctx, event)
}
