package application

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application/port/in/event"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application/port/out/client"
	domainevent "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/event"
)

type (
	NotificationDomainEventHandler struct {
		notificationClient client.NotificationClient
		event.IgnoreUnimplementedDomainEventHandler
	}
)

var _ event.DomainEventHandlers = (*NotificationDomainEventHandler)(nil)

func NewNotificationDomainEventHandler(notificationClient client.NotificationClient) *NotificationDomainEventHandler {
	return &NotificationDomainEventHandler{
		notificationClient: notificationClient,
	}
}

func (h NotificationDomainEventHandler) OnOrderCreated(ctx context.Context, event ddd.DomainEvent) error {
	orderCreated := event.(*domainevent.OrderCreated)
	return h.notificationClient.NotifyOrderCreated(ctx, orderCreated.OrderID, orderCreated.CustomerID)
}

func (h NotificationDomainEventHandler) OnOrderReadied(ctx context.Context, event ddd.DomainEvent) error {
	orderReadied := event.(*domainevent.OrderReadied)
	return h.notificationClient.NotifyOrderReady(ctx, orderReadied.OrderID, orderReadied.CustomerID)
}

func (h NotificationDomainEventHandler) OnOrderCanceled(ctx context.Context, event ddd.DomainEvent) error {
	orderCanceled := event.(*domainevent.OrderCanceled)
	return h.notificationClient.NotifyOrderCanceled(ctx, orderCanceled.OrderID, orderCanceled.CustomerID)
}
