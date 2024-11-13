package handler

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
)

type IntegrationEventHandler[T ddd.Event] struct {
	storeCacheRepository   repository.StoreCacheRepository
	productCacheRepository repository.ProductCacheRepository
}

var _ ddd.EventHandler[ddd.Event] = (*IntegrationEventHandler[ddd.Event])(nil)

func NewIntegrationEventHandler(storeCacheRepository repository.StoreCacheRepository, productCacheRepository repository.ProductCacheRepository) ddd.EventHandler[ddd.Event] {
	return &IntegrationEventHandler[ddd.Event]{
		storeCacheRepository:   storeCacheRepository,
		productCacheRepository: productCacheRepository,
	}
}

func RegisterIntegrationEventHandlers(subscriber am.EventSubscriber, handler ddd.EventHandler[ddd.Event]) error {
	evtMsgHandler := am.MessageHandlerFunc[am.IncomingEventMessage](func(ctx context.Context, eventMsg am.IncomingEventMessage) error {
		return handler.HandleEvent(ctx, eventMsg)
	})
	_, err := subscriber.Subscribe(storev1.StoreAggregateChannel, evtMsgHandler, am.MessageFilter{
		storev1.StoreCreatedEvent,
		storev1.StoreRebrandedEvent,
	}, am.GroupName("baskets-stores"))
	if err != nil {
		return err
	}

	_, err = subscriber.Subscribe(storev1.ProductAggregateChannel, evtMsgHandler, am.MessageFilter{
		storev1.ProductAddedEvent,
		storev1.ProductRebrandedEvent,
		storev1.ProductPriceIncreasedEvent,
		storev1.ProductPriceDecreasedEvent,
		storev1.ProductRemovedEvent,
	}, am.GroupName("baskets-products"))

	return err
}

func (h IntegrationEventHandler[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case storev1.StoreCreatedEvent:
		return h.onStoreCreated(ctx, event)
	case storev1.StoreRebrandedEvent:
		return h.onStoreRebranded(ctx, event)
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
func (h IntegrationEventHandler[T]) onStoreCreated(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storev1.StoreCreated)
	return h.storeCacheRepository.Add(ctx, payload.GetId(), payload.GetName())
}

func (h IntegrationEventHandler[T]) onStoreRebranded(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storev1.StoreRebranded)
	return h.storeCacheRepository.Rename(ctx, payload.GetId(), payload.GetName())
}
func (h IntegrationEventHandler[T]) onProductAdded(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storev1.ProductAdded)
	return h.productCacheRepository.Add(ctx, payload.GetId(), payload.GetStoreId(), payload.GetName(), payload.GetPrice())
}

func (h IntegrationEventHandler[T]) onProductRebranded(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storev1.ProductRebranded)
	return h.productCacheRepository.Rebrand(ctx, payload.GetId(), payload.GetName())
}

func (h IntegrationEventHandler[T]) onProductPriceChanged(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storev1.ProductPriceChanged)
	return h.productCacheRepository.UpdatePrice(ctx, payload.GetId(), payload.GetDelta())
}

func (h IntegrationEventHandler[T]) onProductRemoved(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storev1.ProductRemoved)
	return h.productCacheRepository.Remove(ctx, payload.GetId())
}
