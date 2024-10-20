package grpc

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/application/port/out/client"
	orderv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/api/order/v1"
	"github.com/stackus/errors"
	"google.golang.org/grpc"
)

type OrderClient struct {
	client orderv1.OrderingServiceClient
}

var _ client.OrderClient = (*OrderClient)(nil)

func NewGrpcOrderClient(conn *grpc.ClientConn) *OrderClient {
	return &OrderClient{client: orderv1.NewOrderingServiceClient(conn)}
}

func (c OrderClient) Ready(ctx context.Context, orderID string) error {
	_, err := c.client.ReadyOrder(ctx, &orderv1.ReadyOrderRequest{Id: orderID})
	if err != nil {
		return errors.Wrap(err, "requesting order")
	}
	return nil
}
