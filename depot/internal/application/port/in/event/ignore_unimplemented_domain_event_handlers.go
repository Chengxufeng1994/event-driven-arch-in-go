package event

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
)

type IgnoreUnimplementedDomainEventHandler struct{}

var _ DomainEventHandlers = (*IgnoreUnimplementedDomainEventHandler)(nil)

func NewIgnoreUnimplementedDomainEventHandler() IgnoreUnimplementedDomainEventHandler {
	return IgnoreUnimplementedDomainEventHandler{}
}

// OnShoppingListAssigned implements DomainEventHandlers.
func (i *IgnoreUnimplementedDomainEventHandler) OnShoppingListAssigned(ctx context.Context, event ddd.Event) error {
	return nil
}

// OnShoppingListCanceled implements DomainEventHandlers.
func (i *IgnoreUnimplementedDomainEventHandler) OnShoppingListCanceled(ctx context.Context, event ddd.Event) error {
	return nil
}

// OnShoppingListCompleted implements DomainEventHandlers.
func (i *IgnoreUnimplementedDomainEventHandler) OnShoppingListCompleted(ctx context.Context, event ddd.Event) error {
	return nil
}

// OnShoppingListCreated implements DomainEventHandlers.
func (i *IgnoreUnimplementedDomainEventHandler) OnShoppingListCreated(ctx context.Context, event ddd.Event) error {
	return nil
}
