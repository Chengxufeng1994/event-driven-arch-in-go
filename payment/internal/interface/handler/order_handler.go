package handler

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	orderv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/api/order/v1"
)

func RegisterOrderHandlers(handler ddd.EventHandler[ddd.Event], stream am.EventSubscriber) error {
	evtMsgHandler := am.MessageHandlerFunc[am.EventMessage](func(ctx context.Context, eventMsg am.EventMessage) error {
		return handler.HandleEvent(ctx, eventMsg)
	})

	return stream.Subscribe(orderv1.OrderAggregateChannel, evtMsgHandler, am.MessageFilter{
		orderv1.OrderReadiedEvent,
	}, am.GroupName("payment-orders"))
}
