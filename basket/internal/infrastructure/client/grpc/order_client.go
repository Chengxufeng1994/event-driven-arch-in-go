package grpc

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/port/out/client"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/aggregate"
	orderv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/api/order/v1"
	"github.com/stackus/errors"
	"google.golang.org/grpc"
)

type GrpcOrderClient struct {
	client orderv1.OrderingServiceClient
}

var _ client.OrderClient = (*GrpcOrderClient)(nil)

func NewGrpcOrderClient(conn *grpc.ClientConn) *GrpcOrderClient {
	return &GrpcOrderClient{client: orderv1.NewOrderingServiceClient(conn)}
}

func (c *GrpcOrderClient) Save(ctx context.Context, basket *aggregate.BasketAgg) (string, error) {
	items := make([]*orderv1.Item, 0, len(basket.Items))
	for _, item := range basket.Items {
		items = append(items, &orderv1.Item{
			StoreId:     item.StoreID,
			ProductId:   item.ProductID,
			StoreName:   item.StoreName,
			ProductName: item.ProductName,
			Price:       item.ProductPrice,
			Quantity:    int32(item.Quantity),
		})
	}

	resp, err := c.client.CreateOrder(ctx, &orderv1.CreateOrderRequest{
		Items:      items,
		CustomerId: basket.CustomerID,
		PaymentId:  basket.PaymentID,
	})

	if err != nil {
		return "", errors.Wrap(err, "saving order")
	}
	return resp.Id, nil
}
