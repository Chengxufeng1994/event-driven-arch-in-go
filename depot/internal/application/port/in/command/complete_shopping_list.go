package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/application/port/out/client"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/repository"
)

type CompleteShoppingList struct {
	ID string
}

type CompleteShoppingListHandler struct {
	shoppingListRepository repository.ShoppingListRepository
	orderClient            client.OrderClient
}

func NewCompleteShoppingListHandler(shoppingListRepository repository.ShoppingListRepository, orderClient client.OrderClient) CompleteShoppingListHandler {
	return CompleteShoppingListHandler{
		shoppingListRepository: shoppingListRepository,
		orderClient:            orderClient,
	}
}

func (h CompleteShoppingListHandler) CompleteShoppingList(ctx context.Context, cmd CompleteShoppingList) error {
	list, err := h.shoppingListRepository.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = list.Complete()
	if err != nil {
		return err
	}

	err = h.orderClient.Ready(ctx, list.OrderID)
	if err != nil {
		return err
	}

	return h.shoppingListRepository.Update(ctx, list)
}
