package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/stackus/errors"
)

type CancelBasket struct {
	ID string
}

func NewCancelBasket(id string) CancelBasket {
	return CancelBasket{
		ID: id,
	}
}

type CancelBasketHandler struct {
	basketRepository     repository.BasketRepository
	domainEventPublisher ddd.EventPublisher
}

func NewCancelBasketHandler(basketRepository repository.BasketRepository, domainEventPublisher ddd.EventPublisher) CancelBasketHandler {
	return CancelBasketHandler{
		basketRepository:     basketRepository,
		domainEventPublisher: domainEventPublisher,
	}
}

func (h CancelBasketHandler) CancelBasket(ctx context.Context, cmd CancelBasket) error {
	// find a basket
	basketAgg, err := h.basketRepository.Find(ctx, cmd.ID)
	if err != nil {
		return errors.Wrap(err, "finding")
	}

	// cancel the basket
	if err := basketAgg.Cancel(); err != nil {
		return errors.Wrap(err, "canceling the basket")
	}

	// update the basket
	if err := h.basketRepository.Update(ctx, basketAgg); err != nil {
		return errors.Wrap(err, "updating basket")
	}

	// publish domain events
	return h.domainEventPublisher.Publish(ctx, basketAgg.GetEvents()...)
}
