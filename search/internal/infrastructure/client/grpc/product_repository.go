package grpc

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/rpc"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/application/port/out"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/domain"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
	"google.golang.org/grpc"
)

type ProductRepository struct {
	endpoint string
}

var _ out.ProductRepository = (*ProductRepository)(nil)

func NewGrpcProductRepository(endpoint string) *ProductRepository {
	return &ProductRepository{
		endpoint: endpoint,
	}
}

func (r ProductRepository) Find(ctx context.Context, productID string) (*domain.Product, error) {
	var conn *grpc.ClientConn
	conn, err := r.dial(ctx)
	if err != nil {
		return nil, err
	}

	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)

	resp, err := storev1.NewStoresServiceClient(conn).
		GetProduct(ctx, &storev1.GetProductRequest{Id: productID})
	if err != nil {
		return nil, err
	}

	return r.productToDomain(resp.Product), nil
}

func (r ProductRepository) productToDomain(product *storev1.Product) *domain.Product {
	return &domain.Product{
		ID:      product.GetId(),
		StoreID: product.GetStoreId(),
		Name:    product.GetName(),
	}
}

func (r ProductRepository) dial(ctx context.Context) (*grpc.ClientConn, error) {
	return rpc.Dial(ctx, r.endpoint)
}
