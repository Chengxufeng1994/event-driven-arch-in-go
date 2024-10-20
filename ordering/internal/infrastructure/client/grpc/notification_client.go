package grpc

import (
	"context"

	notificationv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/notification/api/notification/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application/port/out/client"
	"google.golang.org/grpc"
)

type GrpcNotificationClient struct {
	client notificationv1.NotificationsServiceClient
}

var _ client.NotificationClient = (*GrpcNotificationClient)(nil)

func NewGrpcNotificationClient(conn *grpc.ClientConn) *GrpcNotificationClient {
	return &GrpcNotificationClient{client: notificationv1.NewNotificationsServiceClient(conn)}
}

func (c *GrpcNotificationClient) NotifyOrderCreated(ctx context.Context, orderID string, customerID string) error {
	_, err := c.client.NotifyOrderCreated(ctx, &notificationv1.NotifyOrderCreatedRequest{
		OrderId:    orderID,
		CustomerId: customerID,
	})
	return err
}

func (c *GrpcNotificationClient) NotifyOrderCanceled(ctx context.Context, orderID string, customerID string) error {
	_, err := c.client.NotifyOrderCanceled(ctx, &notificationv1.NotifyOrderCanceledRequest{
		OrderId:    orderID,
		CustomerId: customerID,
	})
	return err
}

func (c *GrpcNotificationClient) NotifyOrderReady(ctx context.Context, orderID string, customerID string) error {
	_, err := c.client.NotifyOrderReady(ctx, &notificationv1.NotifyOrderReadyRequest{
		OrderId:    orderID,
		CustomerId: customerID,
	})
	return err
}
