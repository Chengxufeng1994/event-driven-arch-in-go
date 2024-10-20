package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/port/out/client"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/repository"
)

type RemoveItem struct {
	ID        string
	ProductID string
	Quantity  int
}

func NewRemoveItem(id, productID string, quantity int) RemoveItem {
	return RemoveItem{
		ID:        id,
		ProductID: productID,
		Quantity:  quantity,
	}
}

type RemoveItemHandler struct {
	basketRepository repository.BasketRepository
	productClient    client.ProductClient
}

func NewRemoveItemHandler(basketRepository repository.BasketRepository, productClient client.ProductClient) RemoveItemHandler {
	return RemoveItemHandler{
		basketRepository: basketRepository,
		productClient:    productClient,
	}
}

func (h RemoveItemHandler) RemoveItem(ctx context.Context, cmd RemoveItem) error {
	product, err := h.productClient.Find(ctx, cmd.ProductID)
	if err != nil {
		return err
	}

	basketAgg, err := h.basketRepository.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = basketAgg.RemoveItem(product, cmd.Quantity)
	if err != nil {
		return err
	}

	return h.basketRepository.Update(ctx, basketAgg)
}
