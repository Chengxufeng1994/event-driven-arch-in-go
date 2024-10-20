package grpc

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/port/out/client"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/valueobject"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
	"google.golang.org/grpc"
)

type GrpcStoreClient struct {
	client storev1.StoresServiceClient
}

var _ client.StoreClient = (*GrpcStoreClient)(nil)

func NewGrpcStoreClient(conn *grpc.ClientConn) *GrpcStoreClient {
	client := storev1.NewStoresServiceClient(conn)
	return &GrpcStoreClient{
		client: client,
	}
}

func (g *GrpcStoreClient) Find(ctx context.Context, storeID string) (valueobject.Store, error) {
	panic("unimplemented")
}
