package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/repository"
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
	return CancelBasketHandler{
		basketRepository: basketRepository,
	}
}

func (h CancelBasketHandler) CancelBasket(ctx context.Context, cmd CancelBasket) error {
	basketAgg, err := h.basketRepository.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = basketAgg.Cancel()
	if err != nil {
		return err
	}

	return h.basketRepository.Update(ctx, basketAgg)
}
