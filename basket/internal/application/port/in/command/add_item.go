package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/port/out/client"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/repository"
)

type AddItem struct {
	ID        string
	ProductID string
	Quantity  int
}

func NewAddItem(id, productID string, quantity int) AddItem {
	return AddItem{
		ID:        id,
		ProductID: productID,
		Quantity:  quantity,
	}
}

type AddItemHandler struct {
	basketRepository repository.BasketRepository
	productClient    client.ProductClient
	storeClient      client.StoreClient
}

func NewAddItemHandler(basketRepository repository.BasketRepository, productClient client.ProductClient, storeClient client.StoreClient) AddItemHandler {
	return AddItemHandler{
		basketRepository: basketRepository,
		productClient:    productClient,
		storeClient:      storeClient,
	}
}

func (h AddItemHandler) AddItem(ctx context.Context, cmd AddItem) error {
	basketAgg, err := h.basketRepository.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	product, err := h.productClient.Find(ctx, cmd.ProductID)
	if err != nil {
		return err
	}

	store, err := h.storeClient.Find(ctx, product.StoreID)
	if err != nil {
		return err
	}

	err = basketAgg.AddItem(store, product, cmd.Quantity)
	if err != nil {
		return err
	}

	return h.basketRepository.Update(ctx, basketAgg)
}
