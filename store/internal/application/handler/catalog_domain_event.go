package handler

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/di"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
	domainevent "github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/event"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/repository"
)

type CatalogDomainEventHandler[T ddd.Event] struct {
	catalog repository.CatalogRepository
}

var _ ddd.EventHandler[ddd.Event] = (*CatalogDomainEventHandler[ddd.Event])(nil)

func NewCatalogDomainEventHandler(catalog repository.CatalogRepository) *CatalogDomainEventHandler[ddd.Event] {
	return &CatalogDomainEventHandler[ddd.Event]{
		catalog: catalog,
	}
}

func RegisterCatalogDomainEventHandlers(subscriber ddd.EventSubscriber[ddd.Event], handlers ddd.EventHandler[ddd.Event]) {
	subscriber.Subscribe(handlers,
		domainevent.ProductAddedEvent,
		domainevent.ProductRebrandedEvent,
		domainevent.ProductPriceIncreasedEvent,
		domainevent.ProductPriceDecreasedEvent,
		domainevent.ProductRemovedEvent)
}

func RegisterCatalogDomainEventHandlersTx(container di.Container) {
	handlers := ddd.EventHandlerFunc[ddd.Event](func(ctx context.Context, event ddd.Event) error {
		catalogHandlers := di.Get(ctx, "catalogHandlers").(ddd.EventHandler[ddd.Event])

		return catalogHandlers.HandleEvent(ctx, event)
	})

	subscriber := container.Get("domainEventDispatcher").(ddd.EventDispatcher[ddd.Event])

	RegisterCatalogDomainEventHandlers(subscriber, handlers)
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

func (h CatalogDomainEventHandler[T]) onProductAdded(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*aggregate.Product)
	return h.catalog.AddProduct(ctx, payload.ID(), payload.StoreID, payload.Name, payload.Description, payload.SKU, payload.Price)
}

func (h CatalogDomainEventHandler[T]) onProductRebranded(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*aggregate.Product)
	return h.catalog.Rebrand(ctx, payload.ID(), payload.Name, payload.Description)
}

func (h CatalogDomainEventHandler[T]) onProductPriceIncreased(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*domainevent.ProductPriceDelta)
	return h.catalog.UpdatePrice(ctx, payload.ProductID, payload.Delta)
}

func (h CatalogDomainEventHandler[T]) onProductPriceDecreased(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*domainevent.ProductPriceDelta)
	return h.catalog.UpdatePrice(ctx, payload.ProductID, payload.Delta)
}

func (h CatalogDomainEventHandler[T]) onProductRemoved(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*aggregate.Product)
	return h.catalog.RemoveProduct(ctx, payload.ID())
}
