package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/stackus/errors"
)

type RegisterCustomer struct {
	ID        string
	Name      string
	SmsNumber string
}

func NewRegisterCustomer(id, name, smsNumber string) RegisterCustomer {
	return RegisterCustomer{
		ID:        id,
		Name:      name,
		SmsNumber: smsNumber,
	}
}

type RegisterCustomerHandler struct {
	customerRepository   repository.CustomerRepository
	domainEventPublisher ddd.EventPublisher[ddd.AggregateEvent]
}

func NewRegisterCustomerHandler(
	customerRepository repository.CustomerRepository,
	domainEventPublisher ddd.EventPublisher[ddd.AggregateEvent],
) RegisterCustomerHandler {
	return RegisterCustomerHandler{
		customerRepository:   customerRepository,
		domainEventPublisher: domainEventPublisher,
	}
}

func (h RegisterCustomerHandler) RegisterCustomer(ctx context.Context, register RegisterCustomer) error {
	customer, err := aggregate.RegisterCustomer(register.ID, register.Name, register.SmsNumber)
	if err != nil {
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
