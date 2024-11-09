package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/stackus/errors"
)

type ChangeSmsNumber struct {
	ID        string
	SmsNumber string
}

func NewChangeSmsNumber(id, smsNumber string) ChangeSmsNumber {
	return ChangeSmsNumber{
		ID:        id,
		SmsNumber: smsNumber,
	}
}

type ChangeSmsNumberHandler struct {
	customerRepository   repository.CustomerRepository
	domainEventPublisher ddd.EventPublisher[ddd.AggregateEvent]
}

func NewChangeSmsNumberHandler(
	customerRepository repository.CustomerRepository,
	domainEventPublisher ddd.EventPublisher[ddd.AggregateEvent],
) ChangeSmsNumberHandler {
	return ChangeSmsNumberHandler{
		customerRepository:   customerRepository,
		domainEventPublisher: domainEventPublisher,
	}
}

func (h ChangeSmsNumberHandler) ChangeSmsNumber(ctx context.Context, changeSmsNumber ChangeSmsNumber) error {
	customer, err := h.customerRepository.Find(ctx, changeSmsNumber.ID)
	if err != nil {
		return err
	}

	if err := customer.ChangeSmsNumber(changeSmsNumber.SmsNumber); err != nil {
		return err
	}

	if err := h.customerRepository.Save(ctx, customer); err != nil {
		return err
	}

	if err := h.domainEventPublisher.Publish(ctx, customer.Events()...); err != nil {
		return errors.Wrap(err, "could not publish events")
	}

	return nil
}
