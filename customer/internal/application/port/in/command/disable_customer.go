package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/stackus/errors"
)

type DisableCustomer struct {
	ID string
}

func NewDisableCustomer(id string) DisableCustomer {
	return DisableCustomer{
		ID: id,
	}
}

type DisableCustomerHandler struct {
	customerRepository   repository.CustomerRepository
	domainEventPublisher ddd.EventPublisher
}

func NewDisableCustomerHandler(
	customerRepository repository.CustomerRepository,
	domainEventPublisher ddd.EventPublisher,
) DisableCustomerHandler {
	return DisableCustomerHandler{
		customerRepository:   customerRepository,
		domainEventPublisher: domainEventPublisher,
	}
}

func (h DisableCustomerHandler) DisableCustomer(ctx context.Context, disable DisableCustomer) error {
	customer, err := h.customerRepository.Find(ctx, disable.ID)
	if err != nil {
		return err
	}

	err = customer.Disable()
	if err != nil {
		return err
	}

	if err := h.customerRepository.Update(ctx, customer); err != nil {
		return err
	}

	if err := h.domainEventPublisher.Publish(ctx, customer.GetEvents()...); err != nil {
		return errors.Wrap(err, "could not publish events")
	}

	return nil
}
