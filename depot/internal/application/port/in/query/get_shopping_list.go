package query

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/repository"
)

type GetShoppingList struct {
	ID string
}

type GetShoppingListHandler struct {
	shoppingListRepository repository.ShoppingListRepository
}

func NewGetShoppingListHandler(shoppingListRepository repository.ShoppingListRepository) GetShoppingListHandler {
	return GetShoppingListHandler{shoppingListRepository: shoppingListRepository}
}

func (h GetShoppingListHandler) GetShoppingList(ctx context.Context, query GetShoppingList) (*aggregate.ShoppingList, error) {
	return h.shoppingListRepository.Find(ctx, query.ID)
}
