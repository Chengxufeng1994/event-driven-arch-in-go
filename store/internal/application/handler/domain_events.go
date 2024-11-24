package handler

import (
	"context"
	"time"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/errorsotel"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
	domainevent "github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/event"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type DomainEventHandler[T ddd.Event] struct {
	publisher am.EventPublisher
}

var _ ddd.EventHandler[ddd.Event] = (*DomainEventHandler[ddd.Event])(nil)

func NewDomainEventHandlers(publisher am.EventPublisher) *DomainEventHandler[ddd.Event] {
	return &DomainEventHandler[ddd.Event]{publisher: publisher}
}

func RegisterDomainEventHandlers(subscriber ddd.EventSubscriber[ddd.Event], handlers ddd.EventHandler[ddd.Event]) {
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

func (h DomainEventHandler[T]) HandleEvent(ctx context.Context, event T) (err error) {
	span := trace.SpanFromContext(ctx)
	defer func(started time.Time) {
		if err != nil {
			span.AddEvent(
				"Encountered an error handling domain event",
				trace.WithAttributes(errorsotel.ErrAttrs(err)...),
			)
		}
		span.AddEvent("Handled domain event", trace.WithAttributes(
			attribute.Int64("TookMS", time.Since(started).Milliseconds()),
		))
	}(time.Now())

	span.AddEvent("Handling domain event", trace.WithAttributes(
		attribute.String("Event", event.EventName()),
	))

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

func (h DomainEventHandler[T]) onStoreCreated(ctx context.Context, event ddd.Event) error {
	store := event.Payload().(*aggregate.Store)
	return h.publisher.Publish(ctx, storev1.StoreAggregateChannel,
		ddd.NewEvent(storev1.StoreCreatedEvent, &storev1.StoreCreated{
			Id:       store.ID(),
			Name:     store.Name,
			Location: store.Location,
		}))
}

func (h DomainEventHandler[T]) onStoreParticipationEnabled(ctx context.Context, event ddd.Event) error {
	store := event.Payload().(*aggregate.Store)
	return h.publisher.Publish(ctx, storev1.StoreAggregateChannel,
		ddd.NewEvent(storev1.StoreParticipatingToggledEvent, &storev1.StoreParticipationToggled{
			Id:            store.ID(),
			Participating: true,
		}),
	)
}

func (h DomainEventHandler[T]) onStoreParticipationDisabled(ctx context.Context, event ddd.Event) error {
	store := event.Payload().(*aggregate.Store)
	return h.publisher.Publish(ctx, storev1.StoreAggregateChannel,
		ddd.NewEvent(storev1.StoreParticipatingToggledEvent, &storev1.StoreParticipationToggled{
			Id:            store.ID(),
			Participating: false,
		}),
	)
}

func (h DomainEventHandler[T]) onStoreRebranded(ctx context.Context, event ddd.Event) error {
	store := event.Payload().(*aggregate.Store)
	return h.publisher.Publish(ctx, storev1.StoreAggregateChannel,
		ddd.NewEvent(storev1.StoreRebrandedEvent, &storev1.StoreRebranded{
			Id:   store.ID(),
			Name: store.Name,
		}),
	)
}

func (h DomainEventHandler[T]) onProductAdded(ctx context.Context, event ddd.Event) error {
	product := event.Payload().(*aggregate.Product)
	return h.publisher.Publish(ctx, storev1.ProductAggregateChannel,
		ddd.NewEvent(storev1.ProductAddedEvent, &storev1.ProductAdded{
			Id:          product.ID(),
			StoreId:     product.StoreID,
			Name:        product.Name,
			Description: product.Description,
			Sku:         product.SKU,
			Price:       product.Price,
		}),
	)
}

func (h DomainEventHandler[T]) onProductRebranded(ctx context.Context, event ddd.Event) error {
	product := event.Payload().(*aggregate.Product)
	return h.publisher.Publish(ctx, storev1.ProductAggregateChannel,
		ddd.NewEvent(storev1.ProductRebrandedEvent, &storev1.ProductRebranded{
			Id:          product.ID(),
			Name:        product.Name,
			Description: product.Description,
		}),
	)
}

func (h DomainEventHandler[T]) onProductPriceIncreased(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*domainevent.ProductPriceDelta)
	return h.publisher.Publish(ctx, storev1.ProductAggregateChannel,
		ddd.NewEvent(storev1.ProductPriceIncreasedEvent, &storev1.ProductPriceChanged{
			Id:    payload.ProductID,
			Delta: payload.Delta,
		}),
	)
}

func (h DomainEventHandler[T]) onProductPriceDecreased(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*domainevent.ProductPriceDelta)
	return h.publisher.Publish(ctx, storev1.ProductAggregateChannel,
		ddd.NewEvent(storev1.ProductPriceDecreasedEvent, &storev1.ProductPriceChanged{
			Id:    payload.ProductID,
			Delta: payload.Delta,
		}),
	)
}

func (h DomainEventHandler[T]) onProductRemoved(ctx context.Context, event ddd.Event) error {
	product := event.Payload().(*aggregate.Product)
	return h.publisher.Publish(ctx, storev1.ProductAggregateChannel,
		ddd.NewEvent(storev1.ProductRemovedEvent, &storev1.ProductRemoved{
			Id: product.ID(),
		}),
	)
}
