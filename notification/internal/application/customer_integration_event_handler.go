package application

import (
	"context"

	customerv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/customer/api/customer/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/application/port/out/repository"
)

type CustomerIntegrationEventHandler[T ddd.Event] struct {
	cache repository.CustomerCacheRepository
}

var _ ddd.EventHandler[ddd.Event] = (*CustomerIntegrationEventHandler[ddd.Event])(nil)

func NewCustomerIntegrationEventHandler(cache repository.CustomerCacheRepository) *CustomerIntegrationEventHandler[ddd.Event] {
	return &CustomerIntegrationEventHandler[ddd.Event]{
		cache: cache,
	}
}

func (h CustomerIntegrationEventHandler[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case customerv1.CustomerRegisteredEvent:
		return h.onCustomerRegistered(ctx, event)

	case customerv1.CustomerSmsChangedEvent:
		return h.onCustomerSmsChanged(ctx, event)
	}

	return nil
}

func (h CustomerIntegrationEventHandler[T]) onCustomerRegistered(ctx context.Context, event T) error {
	payload := event.Payload().(*customerv1.CustomerRegistered)
	return h.cache.Add(ctx, payload.GetId(), payload.GetName(), payload.GetSmsNumber())
}

func (h CustomerIntegrationEventHandler[T]) onCustomerSmsChanged(ctx context.Context, event T) error {
	payload := event.Payload().(*customerv1.CustomerSmsChanged)
	return h.cache.UpdateSmsNumber(ctx, payload.GetId(), payload.GetSmsNumber())
}
