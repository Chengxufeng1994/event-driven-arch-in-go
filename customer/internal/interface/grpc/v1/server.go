package v1

import (
	"context"

	customerv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/customer/api/customer/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/application/port/in/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/application/port/in/query"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/application/usecase"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/domain/aggregate"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type server struct {
	app usecase.CustomerUsecase
	customerv1.UnimplementedCustomersServiceServer
}

var _ customerv1.CustomersServiceServer = (*server)(nil)

func RegisterServer(_ context.Context, app usecase.CustomerUsecase, registrar grpc.ServiceRegistrar) error {
	customerv1.RegisterCustomersServiceServer(registrar, server{app: app})
	return nil
}

func (s server) RegisterCustomer(ctx context.Context, request *customerv1.RegisterCustomerRequest) (*customerv1.RegisterCustomerResponse, error) {
	id := uuid.New().String()
	err := s.app.RegisterCustomer(ctx, command.NewRegisterCustomer(id, request.GetName(), request.GetSmsNumber()))
	return &customerv1.RegisterCustomerResponse{Id: id}, err
}

func (s server) AuthorizeCustomer(ctx context.Context, request *customerv1.AuthorizeCustomerRequest) (*customerv1.AuthorizeCustomerResponse, error) {
	err := s.app.AuthorizeCustomer(ctx, command.NewAuthorizeCustomer(request.GetId()))
	return &customerv1.AuthorizeCustomerResponse{}, err
}

func (s server) GetCustomer(ctx context.Context, request *customerv1.GetCustomerRequest) (*customerv1.GetCustomerResponse, error) {
	customer, err := s.app.GetCustomer(ctx, query.NewGetCustomer(request.GetId()))
	if err != nil {
		return nil, err
	}

	return &customerv1.GetCustomerResponse{
		Customer: s.customerFromDomain(customer),
	}, nil
}

func (s server) EnableCustomer(ctx context.Context, request *customerv1.EnableCustomerRequest) (*customerv1.EnableCustomerResponse, error) {
	err := s.app.EnableCustomer(ctx, command.NewEnableCustomer(request.GetId()))
	return &customerv1.EnableCustomerResponse{}, err
}

func (s server) DisableCustomer(ctx context.Context, request *customerv1.DisableCustomerRequest) (*customerv1.DisableCustomerResponse, error) {
	err := s.app.DisableCustomer(ctx, command.NewDisableCustomer(request.GetId()))
	return &customerv1.DisableCustomerResponse{}, err
}

func (s server) customerFromDomain(customer *aggregate.CustomerAgg) *customerv1.Customer {
	return &customerv1.Customer{
		Id:        customer.ID,
		Name:      customer.Name,
		SmsNumber: customer.SmsNumber,
		Enabled:   customer.Enabled,
	}
}
