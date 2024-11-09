package v1

import (
	"context"

	paymentv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/payment/api/payment/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/application/port/int/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/application/usecase"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type server struct {
	app usecase.PaymentUseCase
	paymentv1.UnimplementedPaymentsServiceServer
}

var _ paymentv1.PaymentsServiceServer = (*server)(nil)

func RegisterServer(app usecase.PaymentUseCase, registrar grpc.ServiceRegistrar) error {
	paymentv1.RegisterPaymentsServiceServer(registrar, server{app: app})
	return nil
}

func (s server) AuthorizePayment(ctx context.Context, request *paymentv1.AuthorizePaymentRequest,
) (*paymentv1.AuthorizePaymentResponse, error) {
	id := uuid.New().String()
	err := s.app.AuthorizePayment(ctx, command.AuthorizePayment{
		ID:         id,
		CustomerID: request.GetCustomerId(),
		Amount:     request.GetAmount(),
	})
	return &paymentv1.AuthorizePaymentResponse{Id: id}, err
}

func (s server) ConfirmPayment(ctx context.Context, request *paymentv1.ConfirmPaymentRequest,
) (*paymentv1.ConfirmPaymentResponse, error) {
	err := s.app.ConfirmPayment(ctx, command.ConfirmPayment{ID: request.GetId()})
	return &paymentv1.ConfirmPaymentResponse{}, err
}

func (s server) CreateInvoice(ctx context.Context, request *paymentv1.CreateInvoiceRequest,
) (*paymentv1.CreateInvoiceResponse, error) {
	id := uuid.New().String()
	err := s.app.CreateInvoice(ctx, command.CreateInvoice{
		ID:      id,
		OrderID: request.GetOrderId(),
		Amount:  request.GetAmount(),
	})
	return &paymentv1.CreateInvoiceResponse{Id: id}, err
}

func (s server) AdjustInvoice(ctx context.Context, request *paymentv1.AdjustInvoiceRequest,
) (*paymentv1.AdjustInvoiceResponse, error) {
	err := s.app.AdjustInvoice(ctx, command.AdjustInvoice{
		ID:     request.GetId(),
		Amount: request.GetAmount(),
	})
	return &paymentv1.AdjustInvoiceResponse{}, err
}

func (s server) PayInvoice(ctx context.Context, request *paymentv1.PayInvoiceRequest,
) (*paymentv1.PayInvoiceResponse, error) {
	err := s.app.PayInvoice(ctx, command.PayInvoice{
		ID: request.GetId(),
	})
	return &paymentv1.PayInvoiceResponse{}, err
}

func (s server) CancelInvoice(ctx context.Context, request *paymentv1.CancelInvoiceRequest,
) (*paymentv1.CancelInvoiceResponse, error) {
	err := s.app.CancelInvoice(ctx, command.CancelInvoice{
		ID: request.GetId(),
	})
	return &paymentv1.CancelInvoiceResponse{}, err
}
