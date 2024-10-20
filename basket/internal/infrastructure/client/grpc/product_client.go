package grpc

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/port/out/client"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/valueobject"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
	"github.com/stackus/errors"
	"google.golang.org/grpc"
)

type GrpcProductClient struct {
	client storev1.StoresServiceClient
}

var _ client.ProductClient = (*GrpcProductClient)(nil)

func NewGrpcProductClient(conn *grpc.ClientConn) *GrpcProductClient {
	client := storev1.NewStoresServiceClient(conn)
	return &GrpcProductClient{
		client: client,
	}
}

func (c *GrpcProductClient) Find(ctx context.Context, productID string) (valueobject.Product, error) {
	resp, err := c.client.GetProduct(ctx, &storev1.GetProductRequest{
		Id: productID,
	})
	if err != nil {
		return valueobject.Product{}, errors.Wrap(err, "requesting product")
	}

	return c.productToDomain(resp.Product), nil
}

func (c *GrpcProductClient) productToDomain(product *storev1.Product) valueobject.Product {
	return valueobject.NewProduct(
		product.GetId(),
		product.GetStoreId(),
		product.GetName(),
		product.GetPrice())
}
