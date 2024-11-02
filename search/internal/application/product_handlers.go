package application

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/application/port/out"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
)

type ProductHandlers[T ddd.Event] struct {
	cache out.ProductCacheRepository
}

var _ ddd.EventHandler[ddd.Event] = (*ProductHandlers[ddd.Event])(nil)

func NewProductHandlers(cache out.ProductCacheRepository) ProductHandlers[ddd.Event] {
	return ProductHandlers[ddd.Event]{
		cache: cache,
	}
}

func (h ProductHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case storev1.ProductAddedEvent:
		return h.onProductAdded(ctx, event)
	case storev1.ProductRebrandedEvent:
		return h.onProductRebranded(ctx, event)
	case storev1.ProductRemovedEvent:
		return h.onProductRemoved(ctx, event)
	}

	return nil
}

func (h ProductHandlers[T]) onProductAdded(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storev1.ProductAdded)
	return h.cache.Add(ctx, payload.GetId(), payload.GetStoreId(), payload.GetName())
}

func (h ProductHandlers[T]) onProductRebranded(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storev1.ProductRebranded)
	return h.cache.Rebrand(ctx, payload.GetId(), payload.GetName())
}

func (h ProductHandlers[T]) onProductRemoved(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storev1.ProductRemoved)
	return h.cache.Remove(ctx, payload.GetId())
}
