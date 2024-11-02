package application

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
	domainevent "github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/event"
)

type StoreIntegrationEventHandler[T ddd.AggregateEvent] struct {
	publisher am.MessagePublisher[ddd.Event]
}

var _ ddd.EventHandler[ddd.AggregateEvent] = (*StoreIntegrationEventHandler[ddd.AggregateEvent])(nil)

func NewStoreIntegrationEventHandler(publisher am.MessagePublisher[ddd.Event]) *StoreIntegrationEventHandler[ddd.AggregateEvent] {
	return &StoreIntegrationEventHandler[ddd.AggregateEvent]{
		publisher: publisher,
	}
}

func (h StoreIntegrationEventHandler[T]) HandleEvent(ctx context.Context, event T) error {
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

func (h StoreIntegrationEventHandler[T]) onStoreCreated(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domainevent.StoreCreated)
	return h.publisher.Publish(ctx, storev1.StoreAggregateChannel,
		ddd.NewEventBase(storev1.StoreCreatedEvent,
			&storev1.StoreCreated{
				Id:       event.AggregateID(),
				Name:     payload.Name,
				Location: payload.Location,
			}))
}

func (h StoreIntegrationEventHandler[T]) onStoreParticipationEnabled(ctx context.Context, event ddd.AggregateEvent) error {
	return h.publisher.Publish(ctx, storev1.StoreAggregateChannel,
		ddd.NewEventBase(storev1.StoreParticipatingToggledEvent,
			&storev1.StoreParticipationToggled{
				Id:            event.AggregateID(),
				Participating: true,
			}),
	)
}

func (h StoreIntegrationEventHandler[T]) onStoreParticipationDisabled(ctx context.Context, event ddd.AggregateEvent) error {
	return h.publisher.Publish(ctx, storev1.StoreAggregateChannel,
		ddd.NewEventBase(storev1.StoreParticipatingToggledEvent,
			&storev1.StoreParticipationToggled{
				Id:            event.AggregateID(),
				Participating: false,
			}),
	)
}

func (h StoreIntegrationEventHandler[T]) onStoreRebranded(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domainevent.StoreRebranded)
	return h.publisher.Publish(ctx, storev1.StoreAggregateChannel,
		ddd.NewEventBase(storev1.StoreRebrandedEvent,
			&storev1.StoreRebranded{
				Id:   event.AggregateID(),
				Name: payload.Name,
			}),
	)
}
