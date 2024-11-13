package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/repository"
)

type IncreaseProductPrice struct {
	ID    string
	Price float64
}

func NewIncreaseProductPrice(id string, price float64) IncreaseProductPrice {
	return IncreaseProductPrice{
		ID:    id,
		Price: price,
	}
}

type IncreaseProductPriceHandler struct {
	products  repository.ProductRepository
	publisher ddd.EventPublisher[ddd.Event]
}

func NewIncreaseProductPriceHandler(products repository.ProductRepository, publisher ddd.EventPublisher[ddd.Event]) IncreaseProductPriceHandler {
	return IncreaseProductPriceHandler{
		products:  products,
		publisher: publisher,
	}
}

func (h IncreaseProductPriceHandler) IncreaseProductPrice(ctx context.Context, cmd IncreaseProductPrice) error {
	product, err := h.products.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	event, err := product.IncreasePrice(cmd.Price)
	if err != nil {
		return err
	}

	err = h.products.Save(ctx, product)
	if err != nil {
		return err
	}

	return h.publisher.Publish(ctx, event)
}
