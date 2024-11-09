package handler

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/di"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/event"
	domainevent "github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/event"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/repository"
)

type MallDomainEventHandler[T ddd.AggregateEvent] struct {
	mall repository.MallRepository
}

var _ ddd.EventHandler[ddd.AggregateEvent] = (*MallDomainEventHandler[ddd.AggregateEvent])(nil)

func NewMallDomainEventHandler(mall repository.MallRepository) *MallDomainEventHandler[ddd.AggregateEvent] {
	return &MallDomainEventHandler[ddd.AggregateEvent]{mall: mall}
}

func RegisterMallDomainEventHandlers(subscriber ddd.EventSubscriber[ddd.AggregateEvent], handlers ddd.EventHandler[ddd.AggregateEvent]) {
	subscriber.Subscribe(handlers,
		event.StoreCreatedEvent,
		event.StoreParticipationEnabledEvent,
		event.StoreParticipationDisabledEvent,
		event.StoreRebrandedEvent)
}

func RegisterMallDomainEventHandlersTx(container di.Container) {
	handlers := ddd.EventHandlerFunc[ddd.AggregateEvent](func(ctx context.Context, event ddd.AggregateEvent) error {
		mallHandlers := di.Get(ctx, "mallHandlers").(ddd.EventHandler[ddd.AggregateEvent])

		return mallHandlers.HandleEvent(ctx, event)
	})

	subscriber := container.Get("domainEventDispatcher").(ddd.EventDispatcher[ddd.AggregateEvent])

	RegisterMallDomainEventHandlers(subscriber, handlers)
}

func (h *MallDomainEventHandler[T]) HandleEvent(ctx context.Context, event T) error {
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

func (h MallDomainEventHandler[T]) onStoreCreated(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domainevent.StoreCreated)
	return h.mall.AddStore(ctx, event.AggregateID(), payload.Name, payload.Location)
}

func (h MallDomainEventHandler[T]) onStoreParticipationEnabled(ctx context.Context, event ddd.AggregateEvent) error {
	return h.mall.SetStoreParticipation(ctx, event.AggregateID(), true)
}

func (h MallDomainEventHandler[T]) onStoreParticipationDisabled(ctx context.Context, event ddd.AggregateEvent) error {
	return h.mall.SetStoreParticipation(ctx, event.AggregateID(), false)
}

func (h MallDomainEventHandler[T]) onStoreRebranded(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domainevent.StoreRebranded)
	return h.mall.RenameStore(ctx, event.AggregateID(), payload.Name)
}
