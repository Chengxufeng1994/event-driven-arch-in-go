package grpc

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/rpc"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/application/port/out"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/domain"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
	"google.golang.org/grpc"
)

type StoreRepository struct {
	endpoint string
}

var _ out.StoreRepository = (*StoreRepository)(nil)

func NewGrpcStoreRepository(endpoint string) *StoreRepository {
	return &StoreRepository{
		endpoint: endpoint,
	}
}

func (r StoreRepository) Find(ctx context.Context, storeID string) (*domain.Store, error) {
	var conn *grpc.ClientConn
	conn, err := r.dial(ctx)
	if err != nil {
		return nil, err
	}

	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)

	resp, err := storev1.NewStoresServiceClient(conn).
		GetStore(ctx, &storev1.GetStoreRequest{Id: storeID})
	if err != nil {
		return nil, err
	}

	return r.storeToDomain(resp.Store), nil
}

func (r StoreRepository) storeToDomain(store *storev1.Store) *domain.Store {
	return &domain.Store{
		ID:   store.GetId(),
		Name: store.GetName(),
	}
}

func (r StoreRepository) dial(ctx context.Context) (*grpc.ClientConn, error) {
	return rpc.Dial(ctx, r.endpoint)
}
