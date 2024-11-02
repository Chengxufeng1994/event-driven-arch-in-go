package handler

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	domainevent "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/event"
)

func RegisterIntegrationEventHandlers(eventHandlers ddd.EventHandler[ddd.AggregateEvent], domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent]) {
	domainSubscriber.Subscribe(eventHandlers,
		domainevent.OrderCreatedEvent,
		domainevent.OrderReadiedEvent,
		domainevent.OrderCanceledEvent,
		domainevent.OrderCompletedEvent,
	)
}
