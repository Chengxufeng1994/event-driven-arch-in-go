package grpc

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/entity"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/repository"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
	"github.com/stackus/errors"
	"google.golang.org/grpc"
)

type ProductClient struct {
	client storev1.StoresServiceClient
}

var _ repository.ProductRepository = (*ProductClient)(nil)

func NewGrpcProductClient(conn *grpc.ClientConn) *ProductClient {
	return &ProductClient{
		client: storev1.NewStoresServiceClient(conn),
	}
}

func (c *ProductClient) Find(ctx context.Context, productID string) (*entity.Product, error) {
	resp, err := c.client.GetProduct(ctx, &storev1.GetProductRequest{Id: productID})
	if err != nil {
		return nil, errors.Wrap(err, "requesting product")
	}
	return c.productToDomain(resp.Product), nil
}

func (c *ProductClient) productToDomain(product *storev1.Product) *entity.Product {
	return &entity.Product{
		ID:      product.GetId(),
		StoreID: product.GetStoreId(),
		Name:    product.GetName(),
	}
}
