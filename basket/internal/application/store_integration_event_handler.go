package application

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/logger"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
)

type StoreIntegrationEventHandler[T ddd.Event] struct {
	logger logger.Logger
}

var _ ddd.EventHandler[ddd.Event] = (*StoreIntegrationEventHandler[ddd.Event])(nil)

func NewStoreIntegrationEventHandler(logger logger.Logger) *StoreIntegrationEventHandler[ddd.Event] {
	return &StoreIntegrationEventHandler[ddd.Event]{
		logger: logger,
	}
}

func (h StoreIntegrationEventHandler[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case storev1.StoreCreatedEvent:
		return h.onStoreCreated(ctx, event)

	case storev1.StoreParticipatingToggledEvent:
		return h.onStoreParticipationToggled(ctx, event)

	case storev1.StoreRebrandedEvent:
		return h.onStoreRebranded(ctx, event)
	}
	return nil
}

func (h StoreIntegrationEventHandler[T]) onStoreCreated(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storev1.StoreCreated)
	h.logger.Infof(`ID: %s, Name: "%s", Location: "%s"`, payload.GetId(), payload.GetName(), payload.GetLocation())
	return nil
}

func (h StoreIntegrationEventHandler[T]) onStoreParticipationToggled(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storev1.StoreParticipationToggled)
	h.logger.Infof(`ID: %s, Participating: %t`, payload.GetId(), payload.GetParticipating())
	return nil
}

func (h StoreIntegrationEventHandler[T]) onStoreRebranded(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storev1.StoreRebranded)
	h.logger.Infof(`ID: %s, Name: "%s"`, payload.GetId(), payload.GetName())
	return nil
}
