package grpc

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/application/port/out/client"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/valueobject"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
	"github.com/stackus/errors"
	"google.golang.org/grpc"
)

type StoreClient struct {
	client storev1.StoresServiceClient
}

var _ client.StoreClient = (*StoreClient)(nil)

func NewGrpcStoreClient(conn *grpc.ClientConn) *StoreClient {
	return &StoreClient{client: storev1.NewStoresServiceClient(conn)}
}

func (c *StoreClient) Find(ctx context.Context, storeID string) (valueobject.Store, error) {
	resp, err := c.client.GetStore(ctx, &storev1.GetStoreRequest{Id: storeID})
	if err != nil {
		return valueobject.Store{}, errors.Wrap(err, "requesting store")
	}

	return c.storeToDomain(resp.Store), nil
}

func (c *StoreClient) storeToDomain(store *storev1.Store) valueobject.Store {
	return valueobject.Store{
		ID:       store.GetId(),
		Name:     store.GetLocation(),
		Location: store.GetLocation(),
	}
}
