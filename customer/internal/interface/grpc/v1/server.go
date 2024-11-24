package v1

import (
	"context"

	customerv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/customer/api/customer/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/application/port/in/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/application/port/in/query"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/application/usecase"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/errorsotel"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

type server struct {
	app usecase.CustomerUsecase
	customerv1.UnimplementedCustomersServiceServer
}

var _ customerv1.CustomersServiceServer = (*server)(nil)

func RegisterServer(app usecase.CustomerUsecase, registrar grpc.ServiceRegistrar) error {
	customerv1.RegisterCustomersServiceServer(registrar, &server{
		app: app,
	})
	return nil
}

func (s *server) RegisterCustomer(ctx context.Context, request *customerv1.RegisterCustomerRequest) (*customerv1.RegisterCustomerResponse, error) {
	span := trace.SpanFromContext(ctx)

	id := uuid.New().String()

	span.SetAttributes(
		attribute.String("CustomerID", id),
	)

	err := s.app.RegisterCustomer(ctx, command.NewRegisterCustomer(
		id,
		request.GetName(),
		request.GetSmsNumber()),
	)
	if err != nil {
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
	}

	return &customerv1.RegisterCustomerResponse{Id: id}, err
}

func (s *server) ChangeSmsNumber(ctx context.Context, request *customerv1.ChangeSmsNumberRequest) (*customerv1.ChangeSmsNumberResponse, error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("CustomerID", request.GetId()),
	)

	err := s.app.ChangeSmsNumber(ctx, command.NewChangeSmsNumber(
		request.GetId(),
		request.GetSmsNumber(),
	))
	if err != nil {
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
	}

	return &customerv1.ChangeSmsNumberResponse{}, err
}

func (s *server) AuthorizeCustomer(ctx context.Context, request *customerv1.AuthorizeCustomerRequest) (*customerv1.AuthorizeCustomerResponse, error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("CustomerID", request.GetId()),
	)

	err := s.app.AuthorizeCustomer(ctx, command.NewAuthorizeCustomer(
		request.GetId(),
	))
	if err != nil {
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
	}

	return &customerv1.AuthorizeCustomerResponse{}, err
}

func (s *server) GetCustomer(ctx context.Context, request *customerv1.GetCustomerRequest) (*customerv1.GetCustomerResponse, error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("CustomerID", request.GetId()),
	)

	customer, err := s.app.GetCustomer(ctx, query.NewGetCustomer(
		request.GetId(),
	))
	if err != nil {
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	return &customerv1.GetCustomerResponse{
		Customer: s.customerFromDomain(customer),
	}, nil
}

func (s *server) EnableCustomer(ctx context.Context, request *customerv1.EnableCustomerRequest) (*customerv1.EnableCustomerResponse, error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("CustomerID", request.GetId()),
	)

	err := s.app.EnableCustomer(ctx, command.NewEnableCustomer(request.GetId()))
	if err != nil {
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
	}

	return &customerv1.EnableCustomerResponse{}, err
}

func (s *server) DisableCustomer(ctx context.Context, request *customerv1.DisableCustomerRequest) (*customerv1.DisableCustomerResponse, error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("CustomerID", request.GetId()),
	)

	err := s.app.DisableCustomer(ctx, command.NewDisableCustomer(request.GetId()))
	if err != nil {
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
	}

	return &customerv1.DisableCustomerResponse{}, err
}

func (s *server) customerFromDomain(customer *aggregate.Customer) *customerv1.Customer {
	return &customerv1.Customer{
		Id:        customer.ID(),
		Name:      customer.Name,
		SmsNumber: customer.SmsNumber,
		Enabled:   customer.Enabled,
	}
}
