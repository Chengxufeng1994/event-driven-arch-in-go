package event

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
)

type IgnoreUnimplementedDomainEventHandler struct{}

var _ DomainEventHandlers = (*IgnoreUnimplementedDomainEventHandler)(nil)

// OnOrderCanceled implements DomainEventHandlers.
func (i *IgnoreUnimplementedDomainEventHandler) OnOrderCanceled(ctx context.Context, event ddd.DomainEvent) error {
	return nil
}

// OnOrderCompleted implements DomainEventHandlers.
func (i *IgnoreUnimplementedDomainEventHandler) OnOrderCompleted(ctx context.Context, event ddd.DomainEvent) error {
	return nil
}

// OnOrderCreated implements DomainEventHandlers.
func (i *IgnoreUnimplementedDomainEventHandler) OnOrderCreated(ctx context.Context, event ddd.DomainEvent) error {
	return nil
}

// OnOrderReadied implements DomainEventHandlers.
func (i *IgnoreUnimplementedDomainEventHandler) OnOrderReadied(ctx context.Context, event ddd.DomainEvent) error {
	return nil
}
