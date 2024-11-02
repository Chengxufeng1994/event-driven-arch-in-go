package handler

import (
	"context"

	customerv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/customer/api/customer/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
)

func RegisterCustomerIntegrationEventHandlers(customerHandler ddd.EventHandler[ddd.Event], stream am.EventSubscriber) error {
	evtMsgHandler := am.MessageHandlerFunc[am.EventMessage](func(ctx context.Context, eventMsg am.EventMessage) error {
		return customerHandler.HandleEvent(ctx, eventMsg)
	})

	return stream.Subscribe(customerv1.CustomerAggregateChannel, evtMsgHandler, am.MessageFilter{
		customerv1.CustomerRegisteredEvent,
		customerv1.CustomerSmsChangedEvent,
	}, am.GroupName("notification-customers"))
}