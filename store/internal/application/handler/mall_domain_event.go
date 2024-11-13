package handler

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/di"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
	domainevent "github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/event"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/repository"
)

type MallDomainEventHandler[T ddd.Event] struct {
	mall repository.MallRepository
}

var _ ddd.EventHandler[ddd.Event] = (*MallDomainEventHandler[ddd.Event])(nil)

func NewMallDomainEventHandler(mall repository.MallRepository) *MallDomainEventHandler[ddd.Event] {
	return &MallDomainEventHandler[ddd.Event]{
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
		mallHandlers := di.Get(ctx, "mallHandlers").(ddd.EventHandler[ddd.Event])

		return mallHandlers.HandleEvent(ctx, event)
	})

	subscriber := container.Get("domainEventDispatcher").(ddd.EventDispatcher[ddd.Event])

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

func (h MallDomainEventHandler[T]) onStoreCreated(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*aggregate.Store)
	return h.mall.AddStore(ctx, payload.ID(), payload.Name, payload.Location)
}

func (h MallDomainEventHandler[T]) onStoreParticipationEnabled(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*aggregate.Store)
	return h.mall.SetStoreParticipation(ctx, payload.ID(), true)
}

func (h MallDomainEventHandler[T]) onStoreParticipationDisabled(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*aggregate.Store)
	return h.mall.SetStoreParticipation(ctx, payload.ID(), false)
}

func (h MallDomainEventHandler[T]) onStoreRebranded(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*aggregate.Store)
	return h.mall.RenameStore(ctx, payload.ID(), payload.Name)
}
