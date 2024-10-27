package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/stackus/errors"
)

type EnableCustomer struct {
	ID string
}

func NewEnableCustomer(id string) EnableCustomer {
	return EnableCustomer{
		ID: id,
	}
}

type EnableCustomerHandler struct {
	customerRepository   repository.CustomerRepository
	domainEventPublisher ddd.EventPublisher[ddd.AggregateEvent]
}

func NewEnableCustomerHandler(
	customerRepository repository.CustomerRepository,
	domainEventPublisher ddd.EventPublisher[ddd.AggregateEvent],
) EnableCustomerHandler {
	return EnableCustomerHandler{
		customerRepository:   customerRepository,
		domainEventPublisher: domainEventPublisher,
	}
}

func (h EnableCustomerHandler) EnableCustomer(ctx context.Context, enable EnableCustomer) error {
	customer, err := h.customerRepository.Find(ctx, enable.ID)
	if err != nil {
		return err
	}

	err = customer.Enable()
	if err != nil {
		return err
	}

	if err := h.customerRepository.Update(ctx, customer); err != nil {
		return err
	}

	if err := h.domainEventPublisher.Publish(ctx, customer.Events()...); err != nil {
		return errors.Wrap(err, "could not publish events")
	}

	return nil
}
