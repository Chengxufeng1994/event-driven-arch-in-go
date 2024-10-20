package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/domain/repository"
	"github.com/stackus/errors"
)

type AuthorizeCustomer struct {
	ID string
}

func NewAuthorizeCustomer(id string) AuthorizeCustomer {
	return AuthorizeCustomer{
		ID: id,
	}
}

type AuthorizeCustomerHandler struct {
	customerRepository repository.CustomerRepository
}

func NewAuthorizeCustomerHandler(
	customerRepository repository.CustomerRepository,
) AuthorizeCustomerHandler {
	return AuthorizeCustomerHandler{
		customerRepository: customerRepository,
	}
}

func (h AuthorizeCustomerHandler) AuthorizeCustomer(ctx context.Context, authorize AuthorizeCustomer) error {
	customer, err := h.customerRepository.Find(ctx, authorize.ID)
	if err != nil {
		return err
	}

	if !customer.Enabled {
		return errors.Wrap(errors.ErrUnauthorized, "customer is not authorized")
	}

	return nil
}
