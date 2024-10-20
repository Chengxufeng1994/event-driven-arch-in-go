package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/domain/repository"
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
	customerRepository repository.CustomerRepository
}

func NewEnableCustomerHandler(customerRepository repository.CustomerRepository) EnableCustomerHandler {
	return EnableCustomerHandler{
		customerRepository: customerRepository,
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

	return h.customerRepository.Update(ctx, customer)
}
