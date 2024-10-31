package application

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/logger"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
)

type ProductIntegrationEventHandler[T ddd.Event] struct {
	logger logger.Logger
}

var _ ddd.EventHandler[ddd.Event] = (*ProductIntegrationEventHandler[ddd.Event])(nil)

func NewProductIntegrationEventHandler(logger logger.Logger) *ProductIntegrationEventHandler[ddd.Event] {
	return &ProductIntegrationEventHandler[ddd.Event]{
		logger: logger,
	}
}

func (h ProductIntegrationEventHandler[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case storev1.ProductAddedEvent:
		return h.onProductAdded(ctx, event)

	case storev1.ProductRebrandedEvent:
		return h.onProductRebranded(ctx, event)

	case storev1.ProductPriceIncreasedEvent:
	case storev1.ProductPriceDecreasedEvent:
		return h.onProductPriceChanged(ctx, event)

	case storev1.ProductRemovedEvent:
		return h.onProductRemoved(ctx, event)
	}
	return nil
}

func (h ProductIntegrationEventHandler[T]) onProductAdded(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storev1.ProductAdded)
	h.logger.Infof(`ID: %s, StoreID: "%s", Name: "%s", Description: "%s", Price: "%d"`, payload.GetId(), payload.GetStoreId(), payload.GetName(), payload.GetDescription(), payload.GetPrice())
	return nil
}

func (h ProductIntegrationEventHandler[T]) onProductRebranded(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storev1.ProductRebranded)
	h.logger.Infof(`ID: %s, Name: "%s", Description: "%s"`, payload.GetId(), payload.GetName(), payload.GetDescription())
	return nil
}

func (h ProductIntegrationEventHandler[T]) onProductPriceChanged(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storev1.ProductPriceChanged)
	h.logger.Infof(`ID: %s, Delta: "%d"`, payload.GetId(), payload.GetDelta())
	return nil
}

func (h ProductIntegrationEventHandler[T]) onProductRemoved(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storev1.ProductRemoved)
	h.logger.Debugf(`ID: %s`, payload.GetId(), payload.GetId())
	return nil
}
