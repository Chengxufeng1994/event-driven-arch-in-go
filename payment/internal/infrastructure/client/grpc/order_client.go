package grpc

import (
	"context"

	orderv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/api/order/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/application/port/out/client"
	"google.golang.org/grpc"
)

type GrpcOrderClient struct {
	client orderv1.OrderingServiceClient
}

var _ client.OrderClient = (*GrpcOrderClient)(nil)

func NewGrpcOrderClient(conn *grpc.ClientConn) *GrpcOrderClient {
	return &GrpcOrderClient{client: orderv1.NewOrderingServiceClient(conn)}
}

func (c *GrpcOrderClient) Complete(ctx context.Context, invoiceID string, orderID string) error {
	_, err := c.client.CompleteOrder(ctx, &orderv1.CompleteOrderRequest{InvoiceId: invoiceID, Id: orderID})
	return err
}
