package handler

import (
	"context"
	"time"

	customerv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/customer/api/customer/v1"
	domainevent "github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/domain/event"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/errorsotel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type DomainEventHandler[T ddd.AggregateEvent] struct {
	publisher am.EventPublisher
}

var _ ddd.EventHandler[ddd.AggregateEvent] = (*DomainEventHandler[ddd.AggregateEvent])(nil)

func NewDomainEventHandlers(publisher am.EventPublisher) *DomainEventHandler[ddd.AggregateEvent] {
	return &DomainEventHandler[ddd.AggregateEvent]{
		publisher: publisher,
	}
}

func RegisterDomainEventHandlers(subscriber ddd.EventSubscriber[ddd.AggregateEvent], handlers ddd.EventHandler[ddd.AggregateEvent]) {
	subscriber.Subscribe(handlers,
		domainevent.CustomerRegisteredEvent,
		domainevent.CustomerSmsChangedEvent,
		domainevent.CustomerEnabledEvent,
		domainevent.CustomerDisabledEvent,
	)
}

func (h DomainEventHandler[T]) HandleEvent(ctx context.Context, event ddd.AggregateEvent) (err error) {
	span := trace.SpanFromContext(ctx)
	defer func(started time.Time) {
		if err != nil {
			span.AddEvent(
				"Encountered an error handling domain event",
				trace.WithAttributes(errorsotel.ErrAttrs(err)...),
			)
		}
		span.AddEvent("Handled domain event", trace.WithAttributes(
			attribute.Int64("TookMS", time.Since(started).Milliseconds()),
		))
	}(time.Now())

	span.AddEvent("Handling domain event", trace.WithAttributes(
		attribute.String("Event", event.EventName()),
	))

	switch event.EventName() {
	case domainevent.CustomerRegisteredEvent:
		return h.onCustomerRegistered(ctx, event)
	case domainevent.CustomerSmsChangedEvent:
		return h.onCustomerSmsChanged(ctx, event)
	case domainevent.CustomerEnabledEvent:
		return h.onCustomerEnabled(ctx, event)
	case domainevent.CustomerDisabledEvent:
		return h.onCustomerDisabled(ctx, event)
	}
	return nil
}

func (h DomainEventHandler[T]) onCustomerRegistered(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domainevent.CustomerRegistered)
	return h.publisher.Publish(ctx, customerv1.CustomerAggregateChannel,
		ddd.NewEvent(customerv1.CustomerRegisteredEvent, &customerv1.CustomerRegistered{
			Id:        event.AggregateID(),
			Name:      payload.Name,
			SmsNumber: payload.SmsNumber,
		}),
	)
}

func (h DomainEventHandler[T]) onCustomerSmsChanged(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domainevent.CustomerSmsChanged)
	return h.publisher.Publish(ctx, customerv1.CustomerAggregateChannel,
		ddd.NewEvent(customerv1.CustomerSmsChangedEvent, &customerv1.CustomerSmsChanged{
			Id:        event.AggregateID(),
			SmsNumber: payload.SmsNumber,
		}),
	)
}

func (h DomainEventHandler[T]) onCustomerEnabled(ctx context.Context, event ddd.AggregateEvent) error {
	return h.publisher.Publish(ctx, customerv1.CustomerAggregateChannel,
		ddd.NewEvent(customerv1.CustomerEnabledEvent, &customerv1.CustomerEnabled{
			Id: event.AggregateID(),
		}),
	)
}

func (h DomainEventHandler[T]) onCustomerDisabled(ctx context.Context, event ddd.AggregateEvent) error {
	return h.publisher.Publish(ctx, customerv1.CustomerAggregateChannel,
		ddd.NewEvent(customerv1.CustomerDisabledEvent, &customerv1.CustomerDisabled{
			Id: event.AggregateID(),
		}),
	)
}
