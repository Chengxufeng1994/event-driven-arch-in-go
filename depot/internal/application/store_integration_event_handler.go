package application

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
)

type StoreIntegrationEventHandler[T ddd.Event] struct {
	storeCacheRepository repository.StoreCacheRepository
}

var _ ddd.EventHandler[ddd.Event] = (*StoreIntegrationEventHandler[ddd.Event])(nil)

func NewStoreIntegrationEventHandler(storeCacheRepository repository.StoreCacheRepository) *StoreIntegrationEventHandler[ddd.Event] {
	return &StoreIntegrationEventHandler[ddd.Event]{
		storeCacheRepository: storeCacheRepository,
	}
}

func (h StoreIntegrationEventHandler[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case storev1.StoreCreatedEvent:
		return h.onStoreCreated(ctx, event)
	case storev1.StoreRebrandedEvent:
		return h.onStoreRebranded(ctx, event)
	}
	return nil
}

func (h StoreIntegrationEventHandler[T]) onStoreCreated(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storev1.StoreCreated)
	return h.storeCacheRepository.Add(ctx, payload.GetId(), payload.GetName(), payload.GetLocation())
}

func (h StoreIntegrationEventHandler[T]) onStoreRebranded(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storev1.StoreRebranded)
	return h.storeCacheRepository.Rename(ctx, payload.GetId(), payload.GetName())
}
