package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/domain/repository"
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
	customerRepository repository.CustomerRepository
}

func NewDisableCustomerHandler(customerRepository repository.CustomerRepository) DisableCustomerHandler {
	return DisableCustomerHandler{
		customerRepository: customerRepository,
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

	return h.customerRepository.Update(ctx, customer)
}
