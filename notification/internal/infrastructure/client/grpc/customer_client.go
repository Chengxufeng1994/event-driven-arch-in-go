package grpc

import (
	"context"

	customerv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/customer/api/customer/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/application/port/out/client"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/domain/valueobject"
	"google.golang.org/grpc"
)

type CustomerClient struct {
	client customerv1.CustomersServiceClient
}

// Find implements client.CustomerClient.
var _ client.CustomerClient = (*CustomerClient)(nil)

func NewCustomerClient(conn *grpc.ClientConn) *CustomerClient {
	return &CustomerClient{
		client: customerv1.NewCustomersServiceClient(conn),
	}
}

func (c *CustomerClient) Find(ctx context.Context, customerID string) (valueobject.Customer, error) {
	resp, err := c.client.GetCustomer(ctx, &customerv1.GetCustomerRequest{Id: customerID})
	if err != nil {
		return valueobject.Customer{}, err
	}

	return c.customerToDomain(resp.GetCustomer()), nil
}

func (c *CustomerClient) customerToDomain(customer *customerv1.Customer) valueobject.Customer {
	return valueobject.NewCustomer(
		customer.GetId(),
		customer.GetName(),
		customer.GetSmsNumber(),
	)
}
