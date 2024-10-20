package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/domain/repository"
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
	customerRepository repository.CustomerRepository
}

func NewRegisterCustomerHandler(
	customerRepository repository.CustomerRepository,
) RegisterCustomerHandler {
	return RegisterCustomerHandler{
		customerRepository: customerRepository,
	}
}

func (h RegisterCustomerHandler) RegisterCustomer(ctx context.Context, register RegisterCustomer) error {
	customer, err := aggregate.RegisterCustomer(register.ID, register.Name, register.SmsNumber)
	if err != nil {
		return err
	}

	return h.customerRepository.Save(ctx, customer)
}
