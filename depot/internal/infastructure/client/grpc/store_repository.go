package grpc

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/entity"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/rpc"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
	"github.com/stackus/errors"
	"google.golang.org/grpc"
)

type StoreRepository struct {
	endpoint string
}

var _ repository.StoreRepository = (*StoreRepository)(nil)

func NewGrpcStoreRepository(endpoint string) *StoreRepository {
	return &StoreRepository{
		endpoint: endpoint,
	}
}

func (c *StoreRepository) Find(ctx context.Context, storeID string) (*entity.Store, error) {
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
		return nil, errors.Wrap(err, "requesting store")
	}

	return c.storeToDomain(resp.Store), nil
}

func (c StoreRepository) storeToDomain(store *storev1.Store) *entity.Store {
	return &entity.Store{
		ID:       store.GetId(),
		Name:     store.GetName(),
		Location: store.GetLocation(),
	}
}

func (c StoreRepository) dial(ctx context.Context) (*grpc.ClientConn, error) {
	return rpc.Dial(ctx, c.endpoint)
}
