package handler

import (
	"context"
	"time"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/di"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/errorsotel"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
	domainevent "github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/event"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/infrastructure/constants"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type MallHandlers[T ddd.Event] struct {
	mall repository.MallRepository
}

var _ ddd.EventHandler[ddd.Event] = (*MallHandlers[ddd.Event])(nil)

func NewMallDomainEventHandler(mall repository.MallRepository) *MallHandlers[ddd.Event] {
	return &MallHandlers[ddd.Event]{
		mall: mall,
	}
}

func RegisterMallDomainEventHandlers(subscriber ddd.EventSubscriber[ddd.Event], handlers ddd.EventHandler[ddd.Event]) {
	subscriber.Subscribe(handlers,
		domainevent.StoreCreatedEvent,
		domainevent.StoreParticipationEnabledEvent,
		domainevent.StoreParticipationDisabledEvent,
		domainevent.StoreRebrandedEvent)
}

func RegisterMallDomainEventHandlersTx(container di.Container) {
	handlers := ddd.EventHandlerFunc[ddd.Event](func(ctx context.Context, event ddd.Event) error {
		mallHandlers := di.Get(ctx, constants.MallHandlersKey).(ddd.EventHandler[ddd.Event])

		return mallHandlers.HandleEvent(ctx, event)
	})

	subscriber := container.Get(constants.DomainDispatcherKey).(ddd.EventDispatcher[ddd.Event])

	RegisterMallDomainEventHandlers(subscriber, handlers)
}

func (h MallHandlers[T]) HandleEvent(ctx context.Context, event T) (err error) {
	span := trace.SpanFromContext(ctx)
	defer func(started time.Time) {
		if err != nil {
			span.AddEvent(
				"Encountered an error handling mall event",
				trace.WithAttributes(errorsotel.ErrAttrs(err)...),
			)
		}
		span.AddEvent("Handled mall event", trace.WithAttributes(
			attribute.Int64("TookMS", time.Since(started).Milliseconds()),
		))
	}(time.Now())

	span.AddEvent("Handling mall event", trace.WithAttributes(
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
	}
	return nil
}

func (h MallHandlers[T]) onStoreCreated(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*aggregate.Store)
	return h.mall.AddStore(ctx, payload.ID(), payload.Name, payload.Location)
}

func (h MallHandlers[T]) onStoreParticipationEnabled(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*aggregate.Store)
	return h.mall.SetStoreParticipation(ctx, payload.ID(), true)
}

func (h MallHandlers[T]) onStoreParticipationDisabled(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*aggregate.Store)
	return h.mall.SetStoreParticipation(ctx, payload.ID(), false)
}

func (h MallHandlers[T]) onStoreRebranded(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*aggregate.Store)
	return h.mall.RenameStore(ctx, payload.ID(), payload.Name)
}
