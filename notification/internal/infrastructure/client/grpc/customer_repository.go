package grpc

import (
	"context"

	customerv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/customer/api/customer/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/rpc"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/application/port/out/client"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/domain/valueobject"
	"google.golang.org/grpc"
)

type CustomerRepository struct {
	endpoint string
}

var _ client.CustomerRepository = (*CustomerRepository)(nil)

func NewGrpcCustomerRepository(endpoint string) *CustomerRepository {
	return &CustomerRepository{
		endpoint: endpoint,
	}
}

func (r CustomerRepository) Find(ctx context.Context, customerID string) (valueobject.Customer, error) {
	var conn *grpc.ClientConn
	conn, err := r.dial(ctx)
	if err != nil {
		return valueobject.Customer{}, err
	}

	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)

	resp, err := customerv1.NewCustomersServiceClient(conn).
		GetCustomer(ctx, &customerv1.GetCustomerRequest{Id: customerID})
	if err != nil {
		return valueobject.Customer{}, err
	}

	return r.customerToDomain(resp.GetCustomer()), nil
}

func (r CustomerRepository) customerToDomain(customer *customerv1.Customer) valueobject.Customer {
	return valueobject.NewCustomer(
		customer.GetId(),
		customer.GetName(),
		customer.GetSmsNumber(),
	)
}

func (r CustomerRepository) dial(ctx context.Context) (*grpc.ClientConn, error) {
	return rpc.Dial(ctx, r.endpoint)
}
