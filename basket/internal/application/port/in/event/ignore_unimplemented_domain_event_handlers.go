package event

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
)

type IgnoreUnimplementedDomainEventHandler struct{}

// OnBasketStarted implements DomainEventHandlers.
func (i *IgnoreUnimplementedDomainEventHandler) OnBasketStarted(ctx context.Context, event ddd.DomainEvent) error {
	return nil
}

// OnBasketItemAdded implements DomainEventHandlers.
func (i *IgnoreUnimplementedDomainEventHandler) OnBasketItemAdded(ctx context.Context, event ddd.DomainEvent) error {
	return nil
}

// OnBasketItemRemoved implements DomainEventHandlers.
func (i *IgnoreUnimplementedDomainEventHandler) OnBasketItemRemoved(ctx context.Context, event ddd.DomainEvent) error {
	return nil
}

// OnBasketCanceled implements DomainEventHandlers.
func (i *IgnoreUnimplementedDomainEventHandler) OnBasketCanceled(ctx context.Context, event ddd.DomainEvent) error {
	return nil
}

// OnBasketCheckedOut implements DomainEventHandlers.
func (i *IgnoreUnimplementedDomainEventHandler) OnBasketCheckedOut(ctx context.Context, event ddd.DomainEvent) error {
	return nil
}
