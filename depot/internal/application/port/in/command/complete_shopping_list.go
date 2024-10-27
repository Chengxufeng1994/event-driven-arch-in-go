package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/application/port/out/client"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/stackus/errors"
)

type CompleteShoppingList struct {
	ID string
}

type CompleteShoppingListHandler struct {
	shoppingListRepository repository.ShoppingListRepository
	orderClient            client.OrderClient
	domainEventPublisher   ddd.EventPublisher[ddd.AggregateEvent]
}

func NewCompleteShoppingListHandler(
	shoppingListRepository repository.ShoppingListRepository,
	orderClient client.OrderClient,
	domainEventPublisher ddd.EventPublisher[ddd.AggregateEvent],
) CompleteShoppingListHandler {
	return CompleteShoppingListHandler{
		shoppingListRepository: shoppingListRepository,
		orderClient:            orderClient,
		domainEventPublisher:   domainEventPublisher,
	}
}

func (h CompleteShoppingListHandler) CompleteShoppingList(ctx context.Context, cmd CompleteShoppingList) error {
	list, err := h.shoppingListRepository.Find(ctx, cmd.ID)
	if err != nil {
		return errors.Wrap(err, "completing shopping list")
	}

	err = list.Complete()
	if err != nil {
		return errors.Wrap(err, "completing shopping list")
	}

	if err := h.shoppingListRepository.Update(ctx, list); err != nil {
		return errors.Wrap(err, "updating shopping list")
	}

	// publish domain events
	return h.domainEventPublisher.Publish(ctx, list.Events()...)
}
