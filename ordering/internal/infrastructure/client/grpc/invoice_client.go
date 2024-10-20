package grpc

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application/port/out/client"
	paymentv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/payment/api/payment/v1"
	"google.golang.org/grpc"
)

type GrpcInvoiceClient struct {
	client paymentv1.PaymentsServiceClient
}

var _ client.InvoiceClient = (*GrpcInvoiceClient)(nil)

func NewGrpcInvoiceClient(conn *grpc.ClientConn) *GrpcInvoiceClient {
	return &GrpcInvoiceClient{client: paymentv1.NewPaymentsServiceClient(conn)}
}

// Save implements client.InvoiceClient.
func (c *GrpcInvoiceClient) Save(ctx context.Context, orderID string, paymentID string, amount float64) error {
	_, err := c.client.CreateInvoice(ctx, &paymentv1.CreateInvoiceRequest{
		OrderId:   orderID,
		PaymentId: paymentID,
		Amount:    amount,
	})
	return err
}

// Delete implements client.InvoiceClient.
func (c *GrpcInvoiceClient) Delete(ctx context.Context, invoiceID string) error {
	_, err := c.client.CancelInvoice(ctx, &paymentv1.CancelInvoiceRequest{Id: invoiceID})
	return err
}
