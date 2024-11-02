package handler

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/domain/event"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
)

func RegisterCustomerIntegrationHandler(eventHandler ddd.EventHandler[ddd.AggregateEvent], domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent]) {
	domainSubscriber.Subscribe(eventHandler,
		event.CustomerRegisteredEvent,
		event.CustomerSmsChangedEvent,
		event.CustomerEnabledEvent,
		event.CustomerDisabledEvent,
	)
}
