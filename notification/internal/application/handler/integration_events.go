package handler

import (
	"context"
	"time"

	customerv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/customer/api/customer/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/errorsotel"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/application"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/application/port/in/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/application/port/out/repository"
	orderv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/api/order/v1"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type integrationEventHandlers[T ddd.Event] struct {
	app       application.NotificationUseCase
	customers repository.CustomerCacheRepository
}

var _ ddd.EventHandler[ddd.Event] = (*integrationEventHandlers[ddd.Event])(nil)

func NewIntegrationEventHandler(reg registry.Registry, app application.NotificationUseCase, customers repository.CustomerCacheRepository, mws ...am.MessageHandlerMiddleware) am.MessageHandler {
	return am.NewEventHandler(reg, integrationEventHandlers[ddd.Event]{
		app:       app,
		customers: customers,
	}, mws...)
}

func RegisterIntegrationEventHandlers(subscriber am.MessageSubscriber, handlers am.MessageHandler) error {
	if _, err := subscriber.Subscribe(customerv1.CustomerAggregateChannel, handlers, am.MessageFilter{
		customerv1.CustomerRegisteredEvent,
		customerv1.CustomerSmsChangedEvent,
	}, am.GroupName("notification-customers")); err != nil {
		return err
	}

	if _, err := subscriber.Subscribe(customerv1.CustomerAggregateChannel, handlers, am.MessageFilter{
		orderv1.OrderCreatedEvent,
		orderv1.OrderReadiedEvent,
		orderv1.OrderCanceledEvent,
		orderv1.OrderCompletedEvent,
	}, am.GroupName("notification-orders")); err != nil {
		return err
	}

	return nil
}

func (h integrationEventHandlers[T]) HandleEvent(ctx context.Context, event T) (err error) {
	span := trace.SpanFromContext(ctx)
	defer func(started time.Time) {
		if err != nil {
			span.AddEvent(
				"Encountered an error handling integration event",
				trace.WithAttributes(errorsotel.ErrAttrs(err)...),
			)
		}
		span.AddEvent("Handled integration event", trace.WithAttributes(
			attribute.Int64("TookMS", time.Since(started).Milliseconds()),
		))
	}(time.Now())

	span.AddEvent("Handling integration event", trace.WithAttributes(
		attribute.String("Event", event.EventName()),
	))

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

func (h integrationEventHandlers[T]) onCustomerRegistered(ctx context.Context, event T) error {
	payload := event.Payload().(*customerv1.CustomerRegistered)
	return h.customers.Add(ctx, payload.GetId(), payload.GetName(), payload.GetSmsNumber())
}

func (h integrationEventHandlers[T]) onCustomerSmsChanged(ctx context.Context, event T) error {
	payload := event.Payload().(*customerv1.CustomerSmsChanged)
	return h.customers.UpdateSmsNumber(ctx, payload.GetId(), payload.GetSmsNumber())
}

func (h integrationEventHandlers[T]) onOrderCreated(ctx context.Context, event T) error {
	payload := event.Payload().(*orderv1.OrderCreated)
	return h.app.NotifyOrderCreated(ctx, command.NewOrderCreated(
		payload.GetId(),
		payload.GetCustomerId()),
	)
}

func (h integrationEventHandlers[T]) onOrderReadied(ctx context.Context, event T) error {
	payload := event.Payload().(*orderv1.OrderCreated)
	return h.app.NotifyOrderReady(ctx, command.NewOrderReady(
		payload.GetId(),
		payload.GetCustomerId()),
	)
}

func (h integrationEventHandlers[T]) onOrderCanceled(ctx context.Context, event T) error {
	payload := event.Payload().(*orderv1.OrderCanceled)
	return h.app.NotifyOrderCanceled(ctx, command.NewOrderCanceled(
		payload.GetId(),
		payload.GetCustomerId()),
	)
}
