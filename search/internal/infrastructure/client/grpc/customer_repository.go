package grpc

import (
	"context"

	"google.golang.org/grpc"

	customerv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/customer/api/customer/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/rpc"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/application/port/out"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/domain"
)

type CustomerRepository struct {
	endpoint string
}

var _ out.CustomerRepository = (*CustomerRepository)(nil)

func NewGrpcCustomerRepository(endpoint string) *CustomerRepository {
	return &CustomerRepository{
		endpoint: endpoint,
	}
}

func (r CustomerRepository) Find(ctx context.Context, customerID string) (*domain.Customer, error) {
	var conn *grpc.ClientConn
	conn, err := r.dial(ctx)
	if err != nil {
		return nil, err
	}

	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)

	resp, err := customerv1.NewCustomersServiceClient(conn).
		GetCustomer(ctx, &customerv1.GetCustomerRequest{Id: customerID})
	if err != nil {
		return nil, err
	}

	return &domain.Customer{
		ID:   resp.GetCustomer().GetId(),
		Name: resp.GetCustomer().GetName(),
	}, nil
}

func (r CustomerRepository) dial(ctx context.Context) (*grpc.ClientConn, error) {
	return rpc.Dial(ctx, r.endpoint)
}
