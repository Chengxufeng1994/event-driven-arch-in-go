package grpc

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application/port/out/client"
	paymentv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/payment/api/payment/v1"
	"google.golang.org/grpc"
)

type GrpcPaymentClient struct {
	client paymentv1.PaymentsServiceClient
}

var _ client.PaymentClient = (*GrpcPaymentClient)(nil)

func NewGrpcPaymentClient(conn *grpc.ClientConn) *GrpcPaymentClient {
	return &GrpcPaymentClient{client: paymentv1.NewPaymentsServiceClient(conn)}
}

// Confirm implements client.PaymentClient.
func (c *GrpcPaymentClient) Confirm(ctx context.Context, paymentID string) error {
	_, err := c.client.ConfirmPayment(ctx, &paymentv1.ConfirmPaymentRequest{Id: paymentID})
	return err
}
