package handler

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
	domainevent "github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/event"
)

type DomainEventHandler[T ddd.AggregateEvent] struct {
	publisher am.MessagePublisher[ddd.Event]
}

var _ ddd.EventHandler[ddd.AggregateEvent] = (*DomainEventHandler[ddd.AggregateEvent])(nil)

func NewDomainEventHandler(publisher am.MessagePublisher[ddd.Event]) *DomainEventHandler[ddd.AggregateEvent] {
	return &DomainEventHandler[ddd.AggregateEvent]{
		publisher: publisher,
	}
}

func RegisterDomainEventHandlers(subscriber ddd.EventSubscriber[ddd.AggregateEvent], handlers ddd.EventHandler[ddd.AggregateEvent]) {
	subscriber.Subscribe(handlers,
		domainevent.StoreCreatedEvent,
		domainevent.StoreParticipationEnabledEvent,
		domainevent.StoreParticipationDisabledEvent,
		domainevent.StoreRebrandedEvent,
		domainevent.ProductAddedEvent,
		domainevent.ProductRebrandedEvent,
		domainevent.ProductPriceIncreasedEvent,
		domainevent.ProductPriceDecreasedEvent,
		domainevent.ProductRemovedEvent,
	)
}

func (h DomainEventHandler[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case domainevent.StoreCreatedEvent:
		return h.onStoreCreated(ctx, event)
	case domainevent.StoreParticipationEnabledEvent:
		return h.onStoreParticipationEnabled(ctx, event)
	case domainevent.StoreParticipationDisabledEvent:
		return h.onStoreParticipationDisabled(ctx, event)
	case domainevent.StoreRebrandedEvent:
		return h.onStoreRebranded(ctx, event)

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

func (h DomainEventHandler[T]) onStoreCreated(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domainevent.StoreCreated)
	return h.publisher.Publish(ctx, storev1.StoreAggregateChannel,
		ddd.NewEvent(storev1.StoreCreatedEvent, &storev1.StoreCreated{
			Id:       event.AggregateID(),
			Name:     payload.Name,
			Location: payload.Location,
		}))
}

func (h DomainEventHandler[T]) onStoreParticipationEnabled(ctx context.Context, event ddd.AggregateEvent) error {
	return h.publisher.Publish(ctx, storev1.StoreAggregateChannel,
		ddd.NewEvent(storev1.StoreParticipatingToggledEvent, &storev1.StoreParticipationToggled{
			Id:            event.AggregateID(),
			Participating: true,
		}),
	)
}

func (h DomainEventHandler[T]) onStoreParticipationDisabled(ctx context.Context, event ddd.AggregateEvent) error {
	return h.publisher.Publish(ctx, storev1.StoreAggregateChannel,
		ddd.NewEvent(storev1.StoreParticipatingToggledEvent, &storev1.StoreParticipationToggled{
			Id:            event.AggregateID(),
			Participating: false,
		}),
	)
}

func (h DomainEventHandler[T]) onStoreRebranded(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domainevent.StoreRebranded)
	return h.publisher.Publish(ctx, storev1.StoreAggregateChannel,
		ddd.NewEvent(storev1.StoreRebrandedEvent, &storev1.StoreRebranded{
			Id:   event.AggregateID(),
			Name: payload.Name,
		}),
	)
}

func (h DomainEventHandler[T]) onProductAdded(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domainevent.ProductAdded)
	return h.publisher.Publish(ctx, storev1.ProductAggregateChannel,
		ddd.NewEvent(storev1.ProductAddedEvent, &storev1.ProductAdded{
			Id:          event.AggregateID(),
			StoreId:     payload.StoreID,
			Name:        payload.Name,
			Description: payload.Description,
			Sku:         payload.SKU,
			Price:       payload.Price,
		}),
	)
}

func (h DomainEventHandler[T]) onProductRebranded(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domainevent.ProductRebranded)
	return h.publisher.Publish(ctx, storev1.ProductAggregateChannel,
		ddd.NewEvent(storev1.ProductRebrandedEvent, &storev1.ProductRebranded{
			Id:          event.AggregateID(),
			Name:        payload.Name,
			Description: payload.Description,
		}),
	)
}

func (h DomainEventHandler[T]) onProductPriceIncreased(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domainevent.ProductPriceChanged)
	return h.publisher.Publish(ctx, storev1.ProductAggregateChannel,
		ddd.NewEvent(storev1.ProductPriceIncreasedEvent, &storev1.ProductPriceChanged{
			Id:    event.AggregateID(),
			Delta: payload.Delta,
		}),
	)
}

func (h DomainEventHandler[T]) onProductPriceDecreased(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domainevent.ProductPriceChanged)
	return h.publisher.Publish(ctx, storev1.ProductAggregateChannel,
		ddd.NewEvent(storev1.ProductPriceDecreasedEvent, &storev1.ProductPriceChanged{
			Id:    event.AggregateID(),
			Delta: payload.Delta,
		}),
	)
}

func (h DomainEventHandler[T]) onProductRemoved(ctx context.Context, event ddd.AggregateEvent) error {
	return h.publisher.Publish(ctx, storev1.ProductAggregateChannel,
		ddd.NewEvent(storev1.ProductRemovedEvent, &storev1.ProductRemoved{
			Id: event.AggregateID(),
		}),
	)
}
