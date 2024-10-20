package query

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/domain/repository"
)

type GetCustomer struct {
	ID string
}

func NewGetCustomer(id string) GetCustomer {
	return GetCustomer{
		ID: id,
	}
}

type GetCustomerHandler struct {
	customerRepository repository.CustomerRepository
}

func NewGetCustomerHandler(customerRepository repository.CustomerRepository) GetCustomerHandler {
	return GetCustomerHandler{
		customerRepository: customerRepository,
	}
}

func (h GetCustomerHandler) GetCustomer(ctx context.Context, query GetCustomer) (*aggregate.CustomerAgg, error) {
	return h.customerRepository.Find(ctx, query.ID)
}
