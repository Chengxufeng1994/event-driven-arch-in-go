package grpc

import (
	"context"

	"github.com/stackus/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/entity"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/repository"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
)

type GrpcProductRepository struct {
	client storev1.StoresServiceClient
}

var _ repository.ProductRepository = (*GrpcProductRepository)(nil)

func NewGrpcProductRepository(conn *grpc.ClientConn) *GrpcProductRepository {
	client := storev1.NewStoresServiceClient(conn)
	return &GrpcProductRepository{client: client}
}

func (c *GrpcProductRepository) Find(ctx context.Context, productID string) (*entity.Product, error) {
	resp, err := c.client.GetProduct(ctx, &storev1.GetProductRequest{
		Id: productID,
	})
	if err != nil {
		if errors.GRPCCode(err) == codes.NotFound {
			return nil, errors.ErrNotFound.Msg("product was not located")
		}
		return &entity.Product{}, errors.Wrap(err, "requesting product")
	}

	product := c.productToDomain(resp.Product)
	return &product, nil
}

func (c *GrpcProductRepository) productToDomain(product *storev1.Product) entity.Product {
	return entity.NewProduct(
		product.GetId(),
		product.GetStoreId(),
		product.GetName(),
		product.GetPrice())
}
