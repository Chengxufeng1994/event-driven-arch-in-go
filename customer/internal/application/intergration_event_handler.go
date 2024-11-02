package application

import (
	"context"

	customerv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/customer/api/customer/v1"
	domainevent "github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/domain/event"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
)

type CustomerIntegrationEventHandler[T ddd.Event] struct {
	publisher am.MessagePublisher[ddd.Event]
}

var _ ddd.EventHandler[ddd.AggregateEvent] = (*CustomerIntegrationEventHandler[ddd.AggregateEvent])(nil)

func NewCustomerDomainEventHandler(publisher am.MessagePublisher[ddd.Event]) *CustomerIntegrationEventHandler[ddd.AggregateEvent] {
	return &CustomerIntegrationEventHandler[ddd.AggregateEvent]{
		publisher: publisher,
	}
}

func (h CustomerIntegrationEventHandler[T]) HandleEvent(ctx context.Context, event ddd.AggregateEvent) error {
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

func (h CustomerIntegrationEventHandler[T]) onCustomerRegistered(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domainevent.CustomerRegistered)
	return h.publisher.Publish(ctx, customerv1.CustomerAggregateChannel,
		ddd.NewEventBase(customerv1.CustomerRegisteredEvent, &customerv1.CustomerRegistered{
			Id:        event.AggregateID(),
			Name:      payload.Name,
			SmsNumber: payload.SmsNumber,
		}),
	)
}

func (h CustomerIntegrationEventHandler[T]) onCustomerSmsChanged(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domainevent.CustomerSmsChanged)
	return h.publisher.Publish(ctx, customerv1.CustomerAggregateChannel,
		ddd.NewEventBase(customerv1.CustomerSmsChangedEvent, &customerv1.CustomerSmsChanged{
			Id:        event.AggregateID(),
			SmsNumber: payload.SmsNumber,
		}),
	)
}

func (h CustomerIntegrationEventHandler[T]) onCustomerEnabled(ctx context.Context, event ddd.AggregateEvent) error {
	return h.publisher.Publish(ctx, customerv1.CustomerAggregateChannel,
		ddd.NewEventBase(customerv1.CustomerEnabledEvent, &customerv1.CustomerEnabled{
			Id: event.AggregateID(),
		}),
	)
}

func (h CustomerIntegrationEventHandler[T]) onCustomerDisabled(ctx context.Context, event ddd.AggregateEvent) error {
	return h.publisher.Publish(ctx, customerv1.CustomerAggregateChannel,
		ddd.NewEventBase(customerv1.CustomerDisabledEvent, &customerv1.CustomerDisabled{
			Id: event.AggregateID(),
		}),
	)
}
