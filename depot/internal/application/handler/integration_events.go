package handler

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
)

type IntegrationEventHandlers[T ddd.Event] struct {
	products repository.ProductCacheRepository
	stores   repository.StoreCacheRepository
}

var _ ddd.EventHandler[ddd.Event] = (*IntegrationEventHandlers[ddd.Event])(nil)

func RegisterIntegrationEventHandlers(subscriber am.EventSubscriber, handler ddd.EventHandler[ddd.Event]) error {
	evtMsgHandler := am.MessageHandlerFunc[am.IncomingEventMessage](func(ctx context.Context, eventMsg am.IncomingEventMessage) error {
		return handler.HandleEvent(ctx, eventMsg)
	})

	if err := subscriber.Subscribe(storev1.StoreAggregateChannel, evtMsgHandler, am.MessageFilter{
		storev1.StoreCreatedEvent,
		storev1.StoreRebrandedEvent,
	}, am.GroupName("depot-stores")); err != nil {
		return err
	}

	if err := subscriber.Subscribe(storev1.ProductAggregateChannel, evtMsgHandler, am.MessageFilter{
		storev1.ProductAddedEvent,
		storev1.ProductRebrandedEvent,
		storev1.ProductPriceIncreasedEvent,
		storev1.ProductPriceDecreasedEvent,
		storev1.ProductRemovedEvent,
	}, am.GroupName("depot-products")); err != nil {
		return err
	}

	return nil
}

func NewIntegrationEventHandlers(
	products repository.ProductCacheRepository,
	stores repository.StoreCacheRepository,
) *IntegrationEventHandlers[ddd.Event] {
	return &IntegrationEventHandlers[ddd.Event]{
		products: products,
		stores:   stores,
	}
}

func (h IntegrationEventHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case storev1.StoreCreatedEvent:
		return h.onStoreCreated(ctx, event)
	case storev1.StoreRebrandedEvent:
		return h.onStoreRebranded(ctx, event)
	case storev1.ProductAddedEvent:
		return h.onProductAdded(ctx, event)
	case storev1.ProductRebrandedEvent:
		return h.onProductRebranded(ctx, event)
	case storev1.ProductRemovedEvent:
		return h.onProductRemoved(ctx, event)
	}
	return nil
}

func (h IntegrationEventHandlers[T]) onStoreCreated(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storev1.StoreCreated)
	return h.stores.Add(ctx, payload.GetId(), payload.GetName(), payload.GetLocation())
}

func (h IntegrationEventHandlers[T]) onStoreRebranded(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storev1.StoreRebranded)
	return h.stores.Rename(ctx, payload.GetId(), payload.GetName())
}

func (h IntegrationEventHandlers[T]) onProductAdded(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storev1.ProductAdded)
	return h.products.Add(ctx, payload.GetId(), payload.GetStoreId(), payload.GetName())
}

func (h IntegrationEventHandlers[T]) onProductRebranded(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storev1.ProductRebranded)
	return h.products.Rebrand(ctx, payload.GetId(), payload.GetName())
}

func (h IntegrationEventHandlers[T]) onProductRemoved(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storev1.ProductRemoved)
	return h.products.Remove(ctx, payload.GetId())
}
