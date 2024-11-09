package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/stackus/errors"
)

type StartBasket struct {
	ID         string
	CustomerID string
}

func NewStartBasket(id, customerID string) StartBasket {
	return StartBasket{
		ID:         id,
		CustomerID: customerID,
	}
}

type StartBasketHandler struct {
	basketRepository repository.BasketRepository
	publisher        ddd.EventPublisher[ddd.Event]
}

func NewStartBasketHandler(basketRepository repository.BasketRepository, publisher ddd.EventPublisher[ddd.Event]) StartBasketHandler {
	return StartBasketHandler{
		basketRepository: basketRepository,
		publisher:        publisher,
	}
}

func (h StartBasketHandler) StartBasket(ctx context.Context, cmd StartBasket) error {
	basket, err := h.basketRepository.Load(ctx, cmd.ID)
	if err != nil {
		return errors.Wrap(err, "start basket command")
	}

	// create basket
	event, err := basket.Start(cmd.CustomerID)
	if err != nil {
		return errors.Wrap(err, "start basket command")
	}

	// save basket
	if err := h.basketRepository.Save(ctx, basket); err != nil {
		return errors.Wrap(err, "start basket command")
	}

	return h.publisher.Publish(ctx, event)
}
