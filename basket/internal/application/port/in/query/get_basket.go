package query

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/repository"
)

type GetBasket struct {
	ID string
}

func NewGetBasket(id string) GetBasket {
	return GetBasket{
		ID: id,
	}
}

type GetBasketHandler struct {
	basketRepository repository.BasketRepository
}

func NewGetBasketHandler(baskets repository.BasketRepository) GetBasketHandler {
	return GetBasketHandler{
		basketRepository: baskets,
	}
}

func (h GetBasketHandler) GetBasket(ctx context.Context, query GetBasket) (*aggregate.Basket, error) {
	return h.basketRepository.Find(ctx, query.ID)
}
