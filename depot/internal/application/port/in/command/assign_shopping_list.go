package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/stackus/errors"
)

type AssignShoppingList struct {
	ID    string
	BotID string
}

type AssignShoppingListHandler struct {
	shoppingRepository   repository.ShoppingListRepository
	domainEventPublisher ddd.EventPublisher
}

func NewAssignShoppingListHandler(
	shoppingRepository repository.ShoppingListRepository,
	domainEventPublisher ddd.EventPublisher,
) AssignShoppingListHandler {
	return AssignShoppingListHandler{
		shoppingRepository:   shoppingRepository,
		domainEventPublisher: domainEventPublisher,
	}
}

func (h AssignShoppingListHandler) AssignShoppingList(ctx context.Context, cmd AssignShoppingList) error {
	list, err := h.shoppingRepository.Find(ctx, cmd.ID)
	if err != nil {
		return errors.Wrap(err, "assigning shopping list")
	}

	err = list.Assign(cmd.BotID)
	if err != nil {
		return errors.Wrap(err, "assigning shopping list")
	}

	if err := h.shoppingRepository.Update(ctx, list); err != nil {
		return errors.Wrap(err, "updating shopping list")
	}

	// publish domain events
	return h.domainEventPublisher.Publish(ctx, list.GetEvents()...)
}
