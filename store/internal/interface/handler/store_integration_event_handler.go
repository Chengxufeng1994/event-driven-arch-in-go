package handler

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/event"
)

func RegisterStoreIntegrationHandler(eventHandler ddd.EventHandler[ddd.AggregateEvent], domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent]) {
	domainSubscriber.Subscribe(eventHandler,
		event.StoreCreatedEvent,
		event.StoreParticipationEnabledEvent,
		event.StoreParticipationDisabledEvent,
		event.StoreRebrandedEvent)
}
