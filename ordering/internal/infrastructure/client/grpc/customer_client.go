package grpc

import (
	"context"

	customerv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/customer/api/customer/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application/port/out/client"
	"google.golang.org/grpc"
)

type GrpcCustomerClient struct {
	client customerv1.CustomersServiceClient
}

var _ client.CustomerClient = (*GrpcCustomerClient)(nil)

func NewGrpcCustomerClient(conn *grpc.ClientConn) *GrpcCustomerClient {
	return &GrpcCustomerClient{client: customerv1.NewCustomersServiceClient(conn)}
}

// Authorize implements client.CustomerClient.
func (c *GrpcCustomerClient) Authorize(ctx context.Context, customerID string) error {
	_, err := c.client.AuthorizeCustomer(ctx, &customerv1.AuthorizeCustomerRequest{Id: customerID})
	if err != nil {
		return err
	}
	return nil
}
