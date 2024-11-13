package grpc

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/entity"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/repository"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
	"github.com/stackus/errors"
	"google.golang.org/grpc"
)

type GrpcStoreClient struct {
	client storev1.StoresServiceClient
}

var _ repository.StoreRepository = (*GrpcStoreClient)(nil)

func NewGrpcStoreClient(conn *grpc.ClientConn) *GrpcStoreClient {
	return &GrpcStoreClient{client: storev1.NewStoresServiceClient(conn)}
}

func (c *GrpcStoreClient) Find(ctx context.Context, storeID string) (*entity.Store, error) {
	resp, err := c.client.GetStore(ctx, &storev1.GetStoreRequest{
		Id: storeID,
	})
	if err != nil {
		return &entity.Store{}, errors.Wrap(err, "requesting store")
	}
	store := c.storeToDomain(resp.Store)
	return &store, nil
}

func (c *GrpcStoreClient) storeToDomain(store *storev1.Store) entity.Store {
	return entity.NewStore(store.GetId(), store.GetName())
}
