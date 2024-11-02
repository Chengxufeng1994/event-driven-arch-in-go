package application

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
)

type ProductIntegrationEventHandler[T ddd.Event] struct {
	productCacheRepository repository.ProductCacheRepository
}

var _ ddd.EventHandler[ddd.Event] = (*ProductIntegrationEventHandler[ddd.Event])(nil)

func NewProductIntegrationEventHandler(productCacheRepository repository.ProductCacheRepository) *ProductIntegrationEventHandler[ddd.Event] {
	return &ProductIntegrationEventHandler[ddd.Event]{
		productCacheRepository: productCacheRepository,
	}
}

func (h ProductIntegrationEventHandler[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case storev1.ProductAddedEvent:
		return h.onProductAdded(ctx, event)

	case storev1.ProductRebrandedEvent:
		return h.onProductRebranded(ctx, event)

	case storev1.ProductPriceIncreasedEvent:
	case storev1.ProductPriceDecreasedEvent:
		return h.onProductPriceChanged(ctx, event)

	case storev1.ProductRemovedEvent:
		return h.onProductRemoved(ctx, event)
	}
	return nil
}

func (h ProductIntegrationEventHandler[T]) onProductAdded(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storev1.ProductAdded)
	return h.productCacheRepository.Add(ctx, payload.GetId(), payload.GetStoreId(), payload.GetName(), payload.GetPrice())
}

func (h ProductIntegrationEventHandler[T]) onProductRebranded(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storev1.ProductRebranded)
	return h.productCacheRepository.Rebrand(ctx, payload.GetId(), payload.GetName())
}

func (h ProductIntegrationEventHandler[T]) onProductPriceChanged(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storev1.ProductPriceChanged)
	return h.productCacheRepository.UpdatePrice(ctx, payload.GetId(), payload.GetDelta())
}

func (h ProductIntegrationEventHandler[T]) onProductRemoved(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storev1.ProductRemoved)
	return h.productCacheRepository.Remove(ctx, payload.GetId())
}
