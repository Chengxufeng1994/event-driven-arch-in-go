package command

import (
	"context"

	"github.com/stackus/errors"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/valueobject"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
)

type CreateShoppingList struct {
	ID      string
	OrderID string
	Items   []valueobject.OrderItem
}

type CreateShoppingListHandler struct {
	shoppingRepository repository.ShoppingListRepository
	stores             repository.StoreCacheRepository
	products           repository.ProductCacheRepository
	publisher          ddd.EventPublisher[ddd.AggregateEvent]
}

func NewCreateShoppingListHandler(
	shoppingList repository.ShoppingListRepository,
	stores repository.StoreCacheRepository,
	products repository.ProductCacheRepository,
	publisher ddd.EventPublisher[ddd.AggregateEvent],
) CreateShoppingListHandler {
	return CreateShoppingListHandler{
		shoppingRepository: shoppingList,
		stores:             stores,
		products:           products,
		publisher:          publisher,
	}
}

func (h CreateShoppingListHandler) CreateShoppingList(ctx context.Context, cmd CreateShoppingList) error {
	shoppingList := aggregate.CreateShoppingList(cmd.ID, cmd.OrderID)

	for _, item := range cmd.Items {
		store, err := h.stores.Find(ctx, item.StoreID)
		if err != nil {
			return errors.Wrap(err, "finding store")
		}

		product, err := h.products.Find(ctx, item.ProductID)
		if err != nil {
			return errors.Wrap(err, "finding product")
		}

		err = shoppingList.AddItem(store, product, item.Quantity)
		if err != nil {
			return errors.Wrap(err, "adding item to shopping list")
		}
	}

	if err := h.shoppingRepository.Save(ctx, shoppingList); err != nil {
		return errors.Wrap(err, "saving shopping list")
	}

	// publish domain events
	return h.publisher.Publish(ctx, shoppingList.Events()...)
}
