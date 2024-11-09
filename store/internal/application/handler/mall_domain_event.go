package handler

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/event"
)

func RegisterMallDomainEventHandler(mallHandlers ddd.EventHandler[ddd.AggregateEvent], domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent]) {
	domainSubscriber.Subscribe(mallHandlers,
		event.StoreCreatedEvent,
		event.StoreParticipationEnabledEvent,
		event.StoreParticipationDisabledEvent,
		event.StoreRebrandedEvent)
}
