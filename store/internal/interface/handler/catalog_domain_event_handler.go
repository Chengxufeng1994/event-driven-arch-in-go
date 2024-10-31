package handler

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/event"
)

// Event listener
func RegisterCatalogHandler(catalogHandlers ddd.EventHandler[ddd.AggregateEvent], domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent]) {
	domainSubscriber.Subscribe(catalogHandlers,
		event.ProductAddedEvent,
		event.ProductRebrandedEvent,
		event.ProductPriceIncreasedEvent,
		event.ProductPriceDecreasedEvent,
		event.ProductRemovedEvent)
}
