package handler

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
)

func RegisterStoreIntegrationEventHandlers(handler ddd.EventHandler[ddd.Event], subscriber am.EventSubscriber) error {
	evtMsgHandler := func(ctx context.Context, msg am.EventMessage) error {
		return handler.HandleEvent(ctx, msg)
	}

	return subscriber.Subscribe(storev1.StoreAggregateChannel, evtMsgHandler, am.MessageFilter{
		storev1.StoreCreatedEvent,
		storev1.StoreParticipatingToggledEvent,
		storev1.StoreRebrandedEvent,
	})
}
