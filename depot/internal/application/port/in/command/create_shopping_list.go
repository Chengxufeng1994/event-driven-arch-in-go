package command

import (
	"context"

	"github.com/stackus/errors"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/application/port/out/client"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/valueobject"
)

type CreateShoppingList struct {
	ID      string
	OrderID string
	Items   []valueobject.OrderItem
}

type CreateShoppingListHandler struct {
	shoppingRepository repository.ShoppingListRepository
	storeClient        client.StoreClient
	productClient      client.ProductClient
}

func NewCreateShoppingListHandler(
	shoppingRepository repository.ShoppingListRepository,
	storeClient client.StoreClient,
	productClient client.ProductClient,
) CreateShoppingListHandler {
	return CreateShoppingListHandler{
		shoppingRepository: shoppingRepository,
		storeClient:        storeClient,
		productClient:      productClient,
	}
}

func (h CreateShoppingListHandler) CreateShoppingList(ctx context.Context, cmd CreateShoppingList) error {
	shoppingList := aggregate.CreateShoppingList(cmd.ID, cmd.OrderID)
	for _, item := range cmd.Items {
		store, err := h.storeClient.Find(ctx, item.StoreID)
		if err != nil {
			return errors.Wrap(err, "building shopping list")
		}

		product, err := h.productClient.Find(ctx, item.ProductID)
		if err != nil {
			return errors.Wrap(err, "building shopping list")
		}

		err = shoppingList.AddItem(store, product, item.Quantity)
		if err != nil {
			return errors.Wrap(err, "building shopping list")
		}
	}

	return h.shoppingRepository.Save(ctx, shoppingList)
}
