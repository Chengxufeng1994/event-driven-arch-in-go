package handler

import (
	"context"

	customerv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/customer/api/customer/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/application"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/application/port/in/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/application/port/out/repository"
	orderv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/api/order/v1"
)

type IntegrationEventHandler[T ddd.Event] struct {
	app       application.NotificationUseCase
	customers repository.CustomerCacheRepository
}

var _ ddd.EventHandler[ddd.Event] = (*IntegrationEventHandler[ddd.Event])(nil)

func RegisterIntegrationEventHandlers(subscriber am.EventSubscriber, handlers ddd.EventHandler[ddd.Event]) error {
	evtMsgHandler := am.MessageHandlerFunc[am.IncomingEventMessage](func(ctx context.Context, eventMsg am.IncomingEventMessage) error {
		return handlers.HandleEvent(ctx, eventMsg)
	})

	err := subscriber.Subscribe(customerv1.CustomerAggregateChannel, evtMsgHandler, am.MessageFilter{
		customerv1.CustomerRegisteredEvent,
		customerv1.CustomerSmsChangedEvent,
	}, am.GroupName("notification-customers"))
	if err != nil {
		return err
	}

	err = subscriber.Subscribe(customerv1.CustomerAggregateChannel, evtMsgHandler, am.MessageFilter{
		orderv1.OrderCreatedEvent,
		orderv1.OrderReadiedEvent,
		orderv1.OrderCanceledEvent,
		orderv1.OrderCompletedEvent,
	}, am.GroupName("notification-orders"))
	if err != nil {
		return err
	}

	return nil
}

func NewIntegrationEventHandler(
	app application.NotificationUseCase,
	customers repository.CustomerCacheRepository,
) *IntegrationEventHandler[ddd.Event] {
	return &IntegrationEventHandler[ddd.Event]{
		app:       app,
		customers: customers,
	}
}

func (h IntegrationEventHandler[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case customerv1.CustomerRegisteredEvent:
		return h.onCustomerRegistered(ctx, event)
	case customerv1.CustomerSmsChangedEvent:
		return h.onCustomerSmsChanged(ctx, event)
	case orderv1.OrderCreatedEvent:
		return h.onOrderCreated(ctx, event)
	case orderv1.OrderReadiedEvent:
		return h.onOrderReadied(ctx, event)
	case orderv1.OrderCanceledEvent:
		return h.onOrderCanceled(ctx, event)
	}

	return nil
}

func (h IntegrationEventHandler[T]) onCustomerRegistered(ctx context.Context, event T) error {
	payload := event.Payload().(*customerv1.CustomerRegistered)
	return h.customers.Add(ctx, payload.GetId(), payload.GetName(), payload.GetSmsNumber())
}

func (h IntegrationEventHandler[T]) onCustomerSmsChanged(ctx context.Context, event T) error {
	payload := event.Payload().(*customerv1.CustomerSmsChanged)
	return h.customers.UpdateSmsNumber(ctx, payload.GetId(), payload.GetSmsNumber())
}

func (h IntegrationEventHandler[T]) onOrderCreated(ctx context.Context, event T) error {
	payload := event.Payload().(*orderv1.OrderCreated)
	return h.app.NotifyOrderCreated(ctx, command.NewOrderCreated(payload.GetId(), payload.GetCustomerId()))
}

func (h IntegrationEventHandler[T]) onOrderReadied(ctx context.Context, event T) error {
	payload := event.Payload().(*orderv1.OrderCreated)
	return h.app.NotifyOrderReady(ctx, command.NewOrderReady(payload.GetId(), payload.GetCustomerId()))
}

func (h IntegrationEventHandler[T]) onOrderCanceled(ctx context.Context, event T) error {
	payload := event.Payload().(*orderv1.OrderCanceled)
	return h.app.NotifyOrderCanceled(ctx, command.NewOrderCanceled(payload.GetId(), payload.GetCustomerId()))
}
