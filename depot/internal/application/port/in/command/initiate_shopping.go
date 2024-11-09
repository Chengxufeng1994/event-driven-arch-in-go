package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
)

type InitiateShopping struct {
	ID string
}

type InitiateShoppingHandler struct {
	shoppingRepository repository.ShoppingListRepository
	publisher          ddd.EventPublisher[ddd.AggregateEvent]
}

func NewInitiateShoppingHandler(
	shoppingRepository repository.ShoppingListRepository,
	publisher ddd.EventPublisher[ddd.AggregateEvent],
) InitiateShoppingHandler {
	return InitiateShoppingHandler{
		shoppingRepository: shoppingRepository,
		publisher:          publisher,
	}
}

func (h InitiateShoppingHandler) InitiateShopping(ctx context.Context, cmd InitiateShopping) error {
	shoppingList, err := h.shoppingRepository.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = shoppingList.Initiate()
	if err != nil {
		return err
	}

	if err = h.shoppingRepository.Update(ctx, shoppingList); err != nil {
		return err
	}

	// publish domain events
	if err = h.publisher.Publish(ctx, shoppingList.Events()...); err != nil {
		return err
	}

	return nil
}
