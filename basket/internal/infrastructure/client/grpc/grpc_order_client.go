package grpc

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/port/out/client"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/aggregate"
	"google.golang.org/grpc"
)

type GrpcOrderClient struct{}

var _ client.OrderClient = (*GrpcOrderClient)(nil)

func NewGrpcOrderClient(conn *grpc.ClientConn) *GrpcOrderClient {
	return &GrpcOrderClient{}
}

func (g *GrpcOrderClient) Save(ctx context.Context, basket *aggregate.BasketAgg) (string, error) {
	panic("unimplemented")
}
