package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/repository"
)

type DecreaseProductPrice struct {
	ID    string
	Price float64
}

func NewDecreaseProductPrice(id string, price float64) DecreaseProductPrice {
	return DecreaseProductPrice{
		ID:    id,
		Price: price,
	}
}

type DecreaseProductPriceHandler struct {
	products  repository.ProductRepository
	publisher ddd.EventPublisher[ddd.Event]
}

func NewDecreaseProductPriceHandler(products repository.ProductRepository, publisher ddd.EventPublisher[ddd.Event]) DecreaseProductPriceHandler {
	return DecreaseProductPriceHandler{
		products:  products,
		publisher: publisher,
	}
}

func (h DecreaseProductPriceHandler) DecreaseProductPrice(ctx context.Context, cmd DecreaseProductPrice) error {
	product, err := h.products.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	event, err := product.DecreasePrice(cmd.Price)
	if err != nil {
		return err
	}

	err = h.products.Save(ctx, product)
	if err != nil {
		return err
	}

	return h.publisher.Publish(ctx, event)
}
