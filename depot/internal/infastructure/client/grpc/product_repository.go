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

type ProductRepository struct {
	endpoint string
}

var _ repository.ProductRepository = (*ProductRepository)(nil)

func NewGrpcProductRepository(endpoint string) *ProductRepository {
	return &ProductRepository{
		endpoint: endpoint,
	}
}

func (c ProductRepository) Find(ctx context.Context, productID string) (*entity.Product, error) {
	var conn *grpc.ClientConn
	conn, err := c.dial(ctx)
	if err != nil {
		return nil, err
	}

	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)

	resp, err := storev1.NewStoresServiceClient(conn).
		GetProduct(ctx, &storev1.GetProductRequest{Id: productID})
	if err != nil {
		return nil, errors.Wrap(err, "requesting product")
	}
	return c.productToDomain(resp.Product), nil
}

func (c ProductRepository) productToDomain(product *storev1.Product) *entity.Product {
	return &entity.Product{
		ID:      product.GetId(),
		StoreID: product.GetStoreId(),
		Name:    product.GetName(),
	}
}

func (c ProductRepository) dial(ctx context.Context) (*grpc.ClientConn, error) {
	return rpc.Dial(ctx, c.endpoint)
}
