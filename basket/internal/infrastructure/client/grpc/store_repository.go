package grpc

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/entity"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/repository"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
	"github.com/stackus/errors"
	"google.golang.org/grpc"
)

type GrpcStoreRepository struct {
	client storev1.StoresServiceClient
}

var _ repository.StoreRepository = (*GrpcStoreRepository)(nil)

func NewGrpcStoreRepository(conn *grpc.ClientConn) *GrpcStoreRepository {
	return &GrpcStoreRepository{client: storev1.NewStoresServiceClient(conn)}
}

func (c *GrpcStoreRepository) Find(ctx context.Context, storeID string) (*entity.Store, error) {
	resp, err := c.client.GetStore(ctx, &storev1.GetStoreRequest{
		Id: storeID,
	})
	if err != nil {
		return &entity.Store{}, errors.Wrap(err, "requesting store")
	}
	store := c.storeToDomain(resp.Store)
	return &store, nil
}

func (c *GrpcStoreRepository) storeToDomain(store *storev1.Store) entity.Store {
	return entity.NewStore(store.GetId(), store.GetName())
}
