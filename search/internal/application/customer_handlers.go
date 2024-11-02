package application

import (
	"context"

	customerv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/customer/api/customer/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/application/port/out"
)

type CustomerHandlers[T ddd.Event] struct {
	cache out.CustomerCacheRepository
}

var _ ddd.EventHandler[ddd.Event] = (*CustomerHandlers[ddd.Event])(nil)

func NewCustomerHandlers(cache out.CustomerCacheRepository) CustomerHandlers[ddd.Event] {
	return CustomerHandlers[ddd.Event]{
		cache: cache,
	}
}

func (h CustomerHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case customerv1.CustomerRegisteredEvent:
		return h.onCustomerRegistered(ctx, event)
	}

	return nil
}

func (h CustomerHandlers[T]) onCustomerRegistered(ctx context.Context, event T) error {
	payload := event.Payload().(*customerv1.CustomerRegistered)
	return h.cache.Add(ctx, payload.GetId(), payload.GetName())
}
