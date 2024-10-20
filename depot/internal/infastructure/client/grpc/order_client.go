package grpc

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/application/port/out/client"
	"google.golang.org/grpc"
)

type OrderClient struct {
}

var _ client.OrderClient = (*OrderClient)(nil)

func NewGrpcOrderClient(conn *grpc.ClientConn) *OrderClient {
	return &OrderClient{}
}

func (client *OrderClient) Ready(ctx context.Context, orderID string) error {
	panic("unimplemented")
}
