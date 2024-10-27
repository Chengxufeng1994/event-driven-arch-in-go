package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/repository"
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
	basketRepository repository.BasketRepository
}

func NewCancelBasketHandler(basketRepository repository.BasketRepository) CancelBasketHandler {
	return CancelBasketHandler{basketRepository: basketRepository}
}

func (h CancelBasketHandler) CancelBasket(ctx context.Context, cmd CancelBasket) error {
	// find a basket
	basketAgg, err := h.basketRepository.Load(ctx, cmd.ID)
	if err != nil {
		return errors.Wrap(err, "finding basket")
	}

	// cancel the basket
	if err := basketAgg.Cancel(); err != nil {
		return errors.Wrap(err, "canceling the basket")
	}

	// save the basket
	if err := h.basketRepository.Save(ctx, basketAgg); err != nil {
		return errors.Wrap(err, "saving basket")
	}

	return nil
}
