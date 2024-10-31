package handler

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
)

func RegisterProductIntegrationEventHandlers(handler ddd.EventHandler[ddd.Event], subscriber am.EventSubscriber) error {
	evtMsgHandler := func(ctx context.Context, msg am.EventMessage) error {
		return handler.HandleEvent(ctx, msg)
	}

	return subscriber.Subscribe(storev1.ProductAggregateChannel, evtMsgHandler, am.MessageFilter{
		storev1.ProductAddedEvent,
		storev1.ProductRebrandedEvent,
		storev1.ProductPriceIncreasedEvent,
		storev1.ProductPriceDecreasedEvent,
		storev1.ProductRemovedEvent,
	}, am.GroupName("baskets-products"))
}
