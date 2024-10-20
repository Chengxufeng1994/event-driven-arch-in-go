package grpc

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/port/out/client"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/valueobject"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
	"github.com/stackus/errors"
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

func (c *GrpcStoreClient) Find(ctx context.Context, storeID string) (valueobject.Store, error) {
	resp, err := c.client.GetStore(ctx, &storev1.GetStoreRequest{
		Id: storeID,
	})

	if err != nil {
		return valueobject.Store{}, errors.Wrap(err, "requesting store")
	}

	return c.storeToDomain(resp.Store), nil
}

func (c *GrpcStoreClient) storeToDomain(store *storev1.Store) valueobject.Store {
	return valueobject.NewStore(
		store.GetId(),
		store.GetName(),
		store.GetLocation())
}