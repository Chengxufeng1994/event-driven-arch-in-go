package grpc

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/application/port/out/client"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/valueobject"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
	"google.golang.org/grpc"
)

type StoreClient struct {
	client storev1.StoresServiceClient
}

var _ client.StoreClient = (*StoreClient)(nil)

func NewGrpcStoreClient(conn *grpc.ClientConn) *StoreClient {
	client := storev1.NewStoresServiceClient(conn)
	return &StoreClient{
		client: client,
	}
}

func (s *StoreClient) Find(ctx context.Context, storeID string) (valueobject.Store, error) {
	panic("unimplemented")
}
