package handler

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	orderv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/api/order/v1"
)

func RegisterOrderHandlers(orderHandlers ddd.EventHandler[ddd.Event], stream am.EventSubscriber) error {
	evtMsgHandler := am.MessageHandlerFunc[am.EventMessage](func(ctx context.Context, eventMsg am.EventMessage) error {
		return orderHandlers.HandleEvent(ctx, eventMsg)
	})

	return stream.Subscribe(orderv1.OrderAggregateChannel, evtMsgHandler, am.MessageFilter{
		orderv1.OrderCreatedEvent,
		orderv1.OrderReadiedEvent,
		orderv1.OrderCanceledEvent,
		orderv1.OrderCompletedEvent,
	}, am.GroupName("notification-orders"))
}
