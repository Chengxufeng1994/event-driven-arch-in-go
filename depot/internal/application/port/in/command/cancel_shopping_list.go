package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/stackus/errors"
)

type CancelShoppingList struct {
	ID string
}

type CancelShoppingListHandler struct {
	shoppingRepository   repository.ShoppingListRepository
	domainEventPublisher ddd.EventPublisher
}

func NewCancelShoppingListHandler(
	shoppingRepository repository.ShoppingListRepository,
	domainEventPublisher ddd.EventPublisher,
) CancelShoppingListHandler {
	return CancelShoppingListHandler{
		shoppingRepository:   shoppingRepository,
		domainEventPublisher: domainEventPublisher,
	}
}

func (h CancelShoppingListHandler) CancelShoppingList(ctx context.Context, cmd CancelShoppingList) error {
	list, err := h.shoppingRepository.Find(ctx, cmd.ID)
	if err != nil {
		return errors.Wrap(err, "cancel shopping list")
	}

	err = list.Cancel()
	if err != nil {
		return errors.Wrap(err, "cancel shopping list")
	}

	if err := h.shoppingRepository.Update(ctx, list); err != nil {
		return errors.Wrap(err, "updating shopping list")
	}

	// publish domain events
	return h.domainEventPublisher.Publish(ctx, list.GetEvents()...)
}
