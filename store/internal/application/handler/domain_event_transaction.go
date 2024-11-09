package handler

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/di"
)

func RegisterDomainEventHandlersTx(container di.Container) {
	handlers := ddd.EventHandlerFunc[ddd.AggregateEvent](func(ctx context.Context, event ddd.AggregateEvent) error {
		domainEventHandlers := di.Get(ctx, "domainEventHandlers").(ddd.EventHandler[ddd.AggregateEvent])
		return domainEventHandlers.HandleEvent(ctx, event)
	})

	subscriber := container.Get("domainEventDispatcher").(ddd.EventDispatcher[ddd.AggregateEvent])

	RegisterDomainEventHandlers(subscriber, handlers)
}
