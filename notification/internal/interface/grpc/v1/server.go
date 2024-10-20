package v1

import (
	"context"

	notificationv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/notification/api/notification/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/application"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/application/port/in/command"
	"google.golang.org/grpc"
)

type server struct {
	app application.NotificationUseCase
	notificationv1.UnimplementedNotificationsServiceServer
}

var _ notificationv1.NotificationsServiceServer = (*server)(nil)

func RegisterServer(_ context.Context, app application.NotificationUseCase, registrar grpc.ServiceRegistrar) error {
	notificationv1.RegisterNotificationsServiceServer(registrar, server{app: app})
	return nil
}

func (s server) NotifyOrderCreated(ctx context.Context, request *notificationv1.NotifyOrderCreatedRequest) (*notificationv1.NotifyOrderCreatedResponse, error) {
	err := s.app.NotifyOrderCreated(ctx, command.NewOrderCreated(request.GetOrderId(), request.GetCustomerId()))
	return &notificationv1.NotifyOrderCreatedResponse{}, err
}

func (s server) NotifyOrderCanceled(ctx context.Context, request *notificationv1.NotifyOrderCanceledRequest) (*notificationv1.NotifyOrderCanceledResponse, error) {
	err := s.app.NotifyOrderCanceled(ctx, command.NewOrderCanceled(request.GetOrderId(), request.GetCustomerId()))
	return &notificationv1.NotifyOrderCanceledResponse{}, err
}

func (s server) NotifyOrderReady(ctx context.Context, request *notificationv1.NotifyOrderReadyRequest) (*notificationv1.NotifyOrderReadyResponse, error) {
	err := s.app.NotifyOrderReady(ctx, command.NewOrderReady(request.GetOrderId(), request.GetCustomerId()))
	return &notificationv1.NotifyOrderReadyResponse{}, err
}
