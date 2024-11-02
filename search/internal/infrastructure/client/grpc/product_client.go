package grpc

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/application/port/out"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/domain"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
	"google.golang.org/grpc"
)

type ProductClient struct {
	client storev1.StoresServiceClient
}

var _ out.ProductRepository = (*ProductClient)(nil)

func NewProductClient(conn *grpc.ClientConn) *ProductClient {
	return &ProductClient{client: storev1.NewStoresServiceClient(conn)}
}

// Find implements out.ProductRepository.
func (c ProductClient) Find(ctx context.Context, productID string) (*domain.Product, error) {
	resp, err := c.client.GetProduct(ctx, &storev1.GetProductRequest{Id: productID})
	if err != nil {
		return nil, err
	}
	return c.productToDomain(resp.Product), nil
}

func (c ProductClient) productToDomain(product *storev1.Product) *domain.Product {
	return &domain.Product{
		ID:      product.GetId(),
		StoreID: product.GetStoreId(),
		Name:    product.GetName(),
	}
}
