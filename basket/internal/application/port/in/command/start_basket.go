package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/repository"
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
}

func NewStartBasketHandler(basketRepository repository.BasketRepository) StartBasketHandler {
	return StartBasketHandler{
		basketRepository: basketRepository,
	}
}

func (h StartBasketHandler) StartBasket(ctx context.Context, cmd StartBasket) error {
	// create basket
	basket, err := aggregate.StartBasket(cmd.ID, cmd.CustomerID)
	if err != nil {
		return errors.Wrap(err, "start basket command")
	}

	// save basket
	if err := h.basketRepository.Save(ctx, basket); err != nil {
		return errors.Wrap(err, "start basket command")
	}

	return nil
}
