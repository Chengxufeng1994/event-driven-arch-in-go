package application

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/application/port/in/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/application/port/out/repository"
)

type (
	OrderReady struct {
		OrderID    string
		CustomerID string
	}

	NotificationUseCase interface {
		command.Commands
	}

	NotificationApplication struct {
		cache repository.CustomerCacheRepository
	}
)

var _ NotificationUseCase = (*NotificationApplication)(nil)

func New(cache repository.CustomerCacheRepository) *NotificationApplication {
	return &NotificationApplication{
		cache: cache,
	}
}

func (a NotificationApplication) NotifyOrderCreated(_ context.Context, notify command.OrderCreated) error {
	// not implemented

	return nil
}

func (a NotificationApplication) NotifyOrderCanceled(_ context.Context, notify command.OrderCanceled) error {
	// not implemented

	return nil
}

func (a NotificationApplication) NotifyOrderReady(_ context.Context, notify command.OrderReady) error {
	// not implemented

	return nil
}
