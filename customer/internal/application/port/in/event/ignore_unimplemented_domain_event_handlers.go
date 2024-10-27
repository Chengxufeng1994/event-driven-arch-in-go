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

// OnCustomerRegistered implements DomainEventHandlers.
func (i *IgnoreUnimplementedDomainEventHandler) OnCustomerRegistered(ctx context.Context, event ddd.Event) error {
	return nil
}

// OnCustomerAuthorized implements DomainEventHandlers.
func (i *IgnoreUnimplementedDomainEventHandler) OnCustomerAuthorized(ctx context.Context, event ddd.Event) error {
	return nil
}

// OnCustomerDisabled implements DomainEventHandlers.
func (i *IgnoreUnimplementedDomainEventHandler) OnCustomerDisabled(ctx context.Context, event ddd.Event) error {
	return nil
}

// OnCustomerEnabled implements DomainEventHandlers.
func (i *IgnoreUnimplementedDomainEventHandler) OnCustomerEnabled(ctx context.Context, event ddd.Event) error {
	return nil
}
