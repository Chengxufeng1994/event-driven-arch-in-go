package grpc

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/port/out/client"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/valueobject"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
	"google.golang.org/grpc"
)

type GrpcProductClient struct {
	client storev1.StoresServiceClient
}

var _ client.ProductClient = (*GrpcProductClient)(nil)

func NewGrpcProductClient(conn *grpc.ClientConn) *GrpcProductClient {
	client := storev1.NewStoresServiceClient(conn)
	return &GrpcProductClient{
		client: client,
	}
}

func (g *GrpcProductClient) Find(ctx context.Context, productID string) (valueobject.Product, error) {
	panic("unimplemented")
}
