package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/repository"
)

type RebrandProduct struct {
	ID          string
	Name        string
	Description string
}

func NewRebrandProduct(id, name, description string) RebrandProduct {
	return RebrandProduct{
		ID:          id,
		Name:        name,
		Description: description,
	}
}

type RebrandProductHandler struct {
	products  repository.ProductRepository
	publisher ddd.EventPublisher[ddd.Event]
}

func NewRebrandProductHandler(products repository.ProductRepository, publisher ddd.EventPublisher[ddd.Event]) RebrandProductHandler {
	return RebrandProductHandler{
		products:  products,
		publisher: publisher,
	}
}

func (h RebrandProductHandler) RebrandProduct(ctx context.Context, cmd RebrandProduct) error {
	product, err := h.products.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	event, err := product.Rebrand(cmd.Name, cmd.Description)
	if err != nil {
		return err
	}

	err = h.products.Save(ctx, product)
	if err != nil {
		return err
	}

	return h.publisher.Publish(ctx, event)
}
