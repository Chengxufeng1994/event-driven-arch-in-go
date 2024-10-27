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
	productRepository    repository.ProductRepository
	domainEventPublisher ddd.EventPublisher
}

func NewRemoveProductHandler(
	productRepository repository.ProductRepository,
	domainEventPublisher ddd.EventPublisher,
) RemoveProductHandler {
	return RemoveProductHandler{
		productRepository:    productRepository,
		domainEventPublisher: domainEventPublisher,
	}
}

func (h RemoveProductHandler) RemoveProduct(ctx context.Context, cmd RemoveProduct) error {
	product, err := h.productRepository.Find(ctx, cmd.ID)
	if err != nil {
		return errors.Wrap(err, "remove product command")
	}

	if err := product.Remove(); err != nil {
		return errors.Wrap(err, "remove product command")
	}

	if err := h.productRepository.Delete(ctx, cmd.ID); err != nil {
		return errors.Wrap(err, "remove product command")
	}

	return h.domainEventPublisher.Publish(ctx, product.GetEvents()...)
}
