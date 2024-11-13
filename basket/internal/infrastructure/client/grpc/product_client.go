package grpc

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/entity"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/repository"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
	"github.com/stackus/errors"
	"google.golang.org/grpc"
)

type GrpcProductClient struct {
	client storev1.StoresServiceClient
}

var _ repository.ProductRepository = (*GrpcProductClient)(nil)

func NewGrpcProductClient(conn *grpc.ClientConn) *GrpcProductClient {
	client := storev1.NewStoresServiceClient(conn)
	return &GrpcProductClient{
		client: client,
	}
}

func (c *GrpcProductClient) Find(ctx context.Context, productID string) (*entity.Product, error) {
	resp, err := c.client.GetProduct(ctx, &storev1.GetProductRequest{
		Id: productID,
	})
	if err != nil {
		return &entity.Product{}, errors.Wrap(err, "requesting product")
	}

	product := c.productToDomain(resp.Product)
	return &product, nil
}

func (c *GrpcProductClient) productToDomain(product *storev1.Product) entity.Product {
	return entity.NewProduct(
		product.GetId(),
		product.GetStoreId(),
		product.GetName(),
		product.GetPrice())
}
