package application

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/repository"

	domainevent "github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/event"
)

type CatalogDomainEventHandler[T ddd.AggregateEvent] struct {
	catalog repository.CatalogRepository
}

var _ ddd.EventHandler[ddd.AggregateEvent] = (*CatalogDomainEventHandler[ddd.AggregateEvent])(nil)

func NewCatalogDomainEventHandler(catalog repository.CatalogRepository) *CatalogDomainEventHandler[ddd.AggregateEvent] {
	return &CatalogDomainEventHandler[ddd.AggregateEvent]{catalog: catalog}
}

func (h *CatalogDomainEventHandler[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case domainevent.ProductAddedEvent:
		return h.onProductAdded(ctx, event)
	case domainevent.ProductRebrandedEvent:
		return h.onProductRebranded(ctx, event)
	case domainevent.ProductPriceIncreasedEvent:
		return h.onProductPriceIncreased(ctx, event)
	case domainevent.ProductPriceDecreasedEvent:
		return h.onProductPriceDecreased(ctx, event)
	case domainevent.ProductRemovedEvent:
		return h.onProductRemoved(ctx, event)
	}
	return nil
}

func (h CatalogDomainEventHandler[T]) onProductAdded(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domainevent.ProductAdded)
	return h.catalog.AddProduct(ctx, event.AggregateID(), payload.StoreID, payload.Name, payload.Description, payload.SKU, payload.Price)
}

func (h CatalogDomainEventHandler[T]) onProductRebranded(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domainevent.ProductRebranded)
	return h.catalog.Rebrand(ctx, event.AggregateID(), payload.Name, payload.Description)
}

func (h CatalogDomainEventHandler[T]) onProductPriceIncreased(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domainevent.ProductPriceChanged)
	return h.catalog.UpdatePrice(ctx, event.AggregateID(), payload.Delta)
}

func (h CatalogDomainEventHandler[T]) onProductPriceDecreased(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domainevent.ProductPriceChanged)
	return h.catalog.UpdatePrice(ctx, event.AggregateID(), payload.Delta)
}

func (h CatalogDomainEventHandler[T]) onProductRemoved(ctx context.Context, event ddd.AggregateEvent) error {
	return h.catalog.RemoveProduct(ctx, event.AggregateID())
}
