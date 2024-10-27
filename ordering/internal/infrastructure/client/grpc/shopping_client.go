package grpc

import (
	"context"

	depotv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/depot/api/depot/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application/port/out/client"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/valueobject"
	"google.golang.org/grpc"
)

type GrpcShoppingClient struct {
	client depotv1.DepotServiceClient
}

var _ client.ShoppingClient = (*GrpcShoppingClient)(nil)

func NewGrpcShoppingClient(conn *grpc.ClientConn) *GrpcShoppingClient {
	return &GrpcShoppingClient{client: depotv1.NewDepotServiceClient(conn)}
}

// Create implements client.ShoppingClient.
func (c *GrpcShoppingClient) Create(ctx context.Context, order *aggregate.Order) (string, error) {
	var items []*depotv1.OrderItem
	for _, item := range order.Items {
		items = append(items, c.itemFromDomain(item))
	}

	resp, err := c.client.CreateShoppingList(ctx, &depotv1.CreateShoppingListRequest{
		OrderId: order.ID,
		Items:   items,
	})
	if err != nil {
		return "", err
	}

	return resp.GetId(), nil
}

// Cancel implements client.ShoppingClient.
func (c *GrpcShoppingClient) Cancel(ctx context.Context, shoppingID string) error {
	_, err := c.client.CancelShoppingList(ctx, &depotv1.CancelShoppingListRequest{
		Id: shoppingID,
	})

	return err
}

func (c *GrpcShoppingClient) itemFromDomain(item valueobject.Item) *depotv1.OrderItem {
	return &depotv1.OrderItem{
		ProductId: item.ProductID,
		StoreId:   item.StoreID,
		Quantity:  int32(item.Quantity),
	}
}
