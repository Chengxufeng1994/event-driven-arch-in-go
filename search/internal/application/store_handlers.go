package application

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/application/port/out"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
)

type StoreHandlers[T ddd.Event] struct {
	cache out.StoreCacheRepository
}

var _ ddd.EventHandler[ddd.Event] = (*StoreHandlers[ddd.Event])(nil)

func NewStoreHandlers(cache out.StoreCacheRepository) StoreHandlers[ddd.Event] {
	return StoreHandlers[ddd.Event]{
		cache: cache,
	}
}

func (h StoreHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case storev1.StoreCreatedEvent:
		return h.onStoreCreated(ctx, event)
	case storev1.StoreRebrandedEvent:
		return h.onStoreRebranded(ctx, event)
	}

	return nil
}

func (h StoreHandlers[T]) onStoreCreated(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storev1.StoreCreated)
	return h.cache.Add(ctx, payload.GetId(), payload.GetName())
}

func (h StoreHandlers[T]) onStoreRebranded(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storev1.StoreRebranded)
	return h.cache.Rename(ctx, payload.GetId(), payload.GetName())
}
