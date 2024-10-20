package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/repository"
)

type CancelShoppingList struct {
	ID string
}

type CancelShoppingListHandler struct {
	shoppingRepository repository.ShoppingListRepository
}

func NewCancelShoppingListHandler(shoppingRepository repository.ShoppingListRepository) CancelShoppingListHandler {
	return CancelShoppingListHandler{
		shoppingRepository: shoppingRepository,
	}
}

func (h CancelShoppingListHandler) CancelShoppingList(ctx context.Context, cmd CancelShoppingList) error {
	list, err := h.shoppingRepository.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = list.Cancel()
	if err != nil {
		return err
	}

	return h.shoppingRepository.Update(ctx, list)
}
