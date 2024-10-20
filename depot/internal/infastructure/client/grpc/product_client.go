package grpc

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/application/port/out/client"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/valueobject"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
	"github.com/stackus/errors"
	"google.golang.org/grpc"
)

type ProductClient struct {
	client storev1.StoresServiceClient
}

var _ client.ProductClient = (*ProductClient)(nil)

func NewGrpcProductClient(conn *grpc.ClientConn) *ProductClient {
	return &ProductClient{
		client: storev1.NewStoresServiceClient(conn),
	}
}

func (c *ProductClient) Find(ctx context.Context, productID string) (valueobject.Product, error) {
	resp, err := c.client.GetProduct(ctx, &storev1.GetProductRequest{Id: productID})
	if err != nil {
		return valueobject.Product{}, errors.Wrap(err, "requesting product")
	}
	return c.productToDomain(resp.Product), nil
}

func (c *ProductClient) productToDomain(product *storev1.Product) valueobject.Product {
	return valueobject.Product{
		ID:      product.GetId(),
		StoreID: product.GetStoreId(),
		Name:    product.GetName(),
	}
}
