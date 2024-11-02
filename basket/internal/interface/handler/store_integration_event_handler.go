package handler

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
)

func RegisterStoreIntegrationEventHandlers(storeHandler ddd.EventHandler[ddd.Event], stream am.EventSubscriber) error {
	evtMsgHandler := am.MessageHandlerFunc[am.EventMessage](func(ctx context.Context, eventMsg am.EventMessage) error {
		return storeHandler.HandleEvent(ctx, eventMsg)
	})

	return stream.Subscribe(storev1.StoreAggregateChannel, evtMsgHandler, am.MessageFilter{
		storev1.StoreCreatedEvent,
		storev1.StoreRebrandedEvent,
	}, am.GroupName("baskets-stores"))
}
