package grpc

import (
	"context"

	"github.com/stackus/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/entity"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/rpc"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
)

type GrpcProductRepository struct {
	endpoint string
}

var _ repository.ProductRepository = (*GrpcProductRepository)(nil)

func NewGrpcProductRepository(endpoint string) *GrpcProductRepository {
	return &GrpcProductRepository{
		endpoint: endpoint,
	}
}

func (c GrpcProductRepository) Find(ctx context.Context, productID string) (*entity.Product, error) {
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
		if errors.GRPCCode(err) == codes.NotFound {
			return nil, errors.ErrNotFound.Msg("product was not located")
		}
		return &entity.Product{}, errors.Wrap(err, "requesting product")
	}

	product := c.productToDomain(resp.Product)
	return &product, nil
}

func (c GrpcProductRepository) productToDomain(product *storev1.Product) entity.Product {
	return entity.NewProduct(
		product.GetId(),
		product.GetStoreId(),
		product.GetName(),
		product.GetPrice())
}

func (r GrpcProductRepository) dial(ctx context.Context) (*grpc.ClientConn, error) {
	return rpc.Dial(ctx, r.endpoint)
}
