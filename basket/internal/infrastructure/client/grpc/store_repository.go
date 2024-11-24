package grpc

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/entity"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/rpc"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
	"github.com/stackus/errors"
	"google.golang.org/grpc"
)

type GrpcStoreRepository struct {
	endpoint string
}

var _ repository.StoreRepository = (*GrpcStoreRepository)(nil)

func NewGrpcStoreRepository(endpoint string) *GrpcStoreRepository {
	return &GrpcStoreRepository{
		endpoint: endpoint,
	}
}

func (c GrpcStoreRepository) Find(ctx context.Context, storeID string) (*entity.Store, error) {
	var conn *grpc.ClientConn
	conn, err := c.dial(ctx)
	if err != nil {
		return nil, err
	}

	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)

	resp, err := storev1.NewStoresServiceClient(conn).
		GetStore(ctx, &storev1.GetStoreRequest{Id: storeID})
	if err != nil {
		return &entity.Store{}, errors.Wrap(err, "requesting store")
	}
	store := c.storeToDomain(resp.Store)
	return &store, nil
}

func (c GrpcStoreRepository) storeToDomain(store *storev1.Store) entity.Store {
	return entity.NewStore(store.GetId(), store.GetName())
}

func (r GrpcStoreRepository) dial(ctx context.Context) (*grpc.ClientConn, error) {
	return rpc.Dial(ctx, r.endpoint)
}
