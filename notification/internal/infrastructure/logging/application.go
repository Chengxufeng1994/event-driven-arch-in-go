package logging

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/logger"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/application"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/application/port/in/command"
)

type Application struct {
	application.NotificationUseCase
	logger logger.Logger
}

var _ application.NotificationUseCase = (*Application)(nil)

func NewLogApplicationAccess(application application.NotificationUseCase, logger logger.Logger) *Application {
	return &Application{
		NotificationUseCase: application,
		logger:              logger,
	}
}

// NotifyOrderCreated implements application.NotificationUseCase.
func (a *Application) NotifyOrderCreated(ctx context.Context, notify command.OrderCreated) (err error) {
	a.logger.Info("--> Notifications.NotifyOrderCreated")
	defer func() { a.logger.WithError(err).Info("<-- Notifications.NotifyOrderCreated") }()
	return a.NotificationUseCase.NotifyOrderCreated(ctx, notify)
}

// NotifyOrderCanceled implements application.NotificationUseCase.
func (a *Application) NotifyOrderCanceled(ctx context.Context, notify command.OrderCanceled) (err error) {
	a.logger.Info("--> Notifications.NotifyOrderCanceled")
	defer func() { a.logger.WithError(err).Info("<-- Notifications.NotifyOrderCanceled") }()
	return a.NotificationUseCase.NotifyOrderCanceled(ctx, notify)
}

// NotifyOrderReady implements application.NotificationUseCase.
func (a *Application) NotifyOrderReady(ctx context.Context, notify command.OrderReady) (err error) {
	a.logger.Info("--> Notifications.NotifyOrderReady")
	defer func() { a.logger.WithError(err).Info("<-- Notifications.NotifyOrderReady") }()
	return a.NotificationUseCase.NotifyOrderReady(ctx, notify)
}
