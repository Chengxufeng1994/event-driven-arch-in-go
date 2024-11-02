package grpc

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/application/port/out"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/domain"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
	"google.golang.org/grpc"
)

type StoreClient struct {
	client storev1.StoresServiceClient
}

var _ out.StoreRepository = (*StoreClient)(nil)

func NewStoreClient(conn *grpc.ClientConn) *StoreClient {
	return &StoreClient{client: storev1.NewStoresServiceClient(conn)}
}

func (r *StoreClient) Find(ctx context.Context, storeID string) (*domain.Store, error) {
	resp, err := r.client.GetStore(ctx, &storev1.GetStoreRequest{Id: storeID})
	if err != nil {
		return nil, err
	}

	return r.storeToDomain(resp.Store), nil
}

func (r *StoreClient) storeToDomain(store *storev1.Store) *domain.Store {
	return &domain.Store{
		ID:   store.GetId(),
		Name: store.GetName(),
	}
}
