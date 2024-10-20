package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/repository"
)

type AssignShoppingList struct {
	ID    string
	BotID string
}

type AssignShoppingListHandler struct {
	shoppingRepository repository.ShoppingListRepository
}

func NewAssignShoppingListHandler(shoppingRepository repository.ShoppingListRepository) AssignShoppingListHandler {
	return AssignShoppingListHandler{
		shoppingRepository: shoppingRepository,
	}
}

func (h AssignShoppingListHandler) AssignShoppingList(ctx context.Context, cmd AssignShoppingList) error {
	list, err := h.shoppingRepository.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = list.Assign(cmd.BotID)
	if err != nil {
		return err
	}

	return h.shoppingRepository.Update(ctx, list)
}
