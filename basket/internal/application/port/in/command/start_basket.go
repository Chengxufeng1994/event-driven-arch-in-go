package command

import (
	"context"
	"fmt"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/repository"
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
	fmt.Println("123")
	basketAgg, err := aggregate.StartBasket(cmd.ID, cmd.CustomerID)
	if err != nil {
		return err
	}

	return h.basketRepository.Save(ctx, basketAgg)
}
