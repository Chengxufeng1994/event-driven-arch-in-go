package grpc

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/application/port/out/client"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/valueobject"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
	"google.golang.org/grpc"
)

type ProductClient struct {
	client storev1.StoresServiceClient
}

var _ client.ProductClient = (*ProductClient)(nil)

func NewGrpcProductClient(conn *grpc.ClientConn) *ProductClient {
	client := storev1.NewStoresServiceClient(conn)
	return &ProductClient{
		client: client,
	}
}

// Find implements client.ProductClient.
func (p *ProductClient) Find(ctx context.Context, productID string) (valueobject.Product, error) {
	panic("unimplemented")
}
