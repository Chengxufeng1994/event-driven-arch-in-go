package grpc

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/entity"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/repository"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
	"github.com/stackus/errors"
	"google.golang.org/grpc"
)

type StoreClient struct {
	client storev1.StoresServiceClient
}

var _ repository.StoreRepository = (*StoreClient)(nil)

func NewGrpcStoreClient(conn *grpc.ClientConn) *StoreClient {
	return &StoreClient{client: storev1.NewStoresServiceClient(conn)}
}

func (c *StoreClient) Find(ctx context.Context, storeID string) (*entity.Store, error) {
	resp, err := c.client.GetStore(ctx, &storev1.GetStoreRequest{Id: storeID})
	if err != nil {
		return nil, errors.Wrap(err, "requesting store")
	}

	return c.storeToDomain(resp.Store), nil
}

func (c *StoreClient) storeToDomain(store *storev1.Store) *entity.Store {
	return &entity.Store{
		ID:       store.GetId(),
		Name:     store.GetName(),
		Location: store.GetLocation(),
	}
}
