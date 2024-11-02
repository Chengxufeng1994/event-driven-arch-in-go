package grpc

import (
	"context"

	customerv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/customer/api/customer/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/application/port/out"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/domain"
	"google.golang.org/grpc"
)

type CustomerClient struct {
	client customerv1.CustomersServiceClient
}

var _ out.CustomerRepository = (*CustomerClient)(nil)

func NewCustomerClient(conn *grpc.ClientConn) *CustomerClient {
	return &CustomerClient{client: customerv1.NewCustomersServiceClient(conn)}
}

func (c *CustomerClient) Find(ctx context.Context, customerID string) (*domain.Customer, error) {
	resp, err := c.client.GetCustomer(ctx, &customerv1.GetCustomerRequest{Id: customerID})
	if err != nil {
		return nil, err
	}

	return &domain.Customer{
		ID:   resp.GetCustomer().GetId(),
		Name: resp.GetCustomer().GetName(),
	}, nil
}
