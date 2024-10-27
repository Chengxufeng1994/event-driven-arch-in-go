package event

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
)

type DomainEventHandlers interface {
	OnStoreCreated(ctx context.Context, event ddd.DomainEvent) error
	OnStoreParticipationEnabled(ctx context.Context, event ddd.DomainEvent) error
	OnStoreParticipationDisabled(ctx context.Context, event ddd.DomainEvent) error
	OnProductAdded(ctx context.Context, event ddd.DomainEvent) error
	OnProductRemoved(ctx context.Context, event ddd.DomainEvent) error
}

type IgnoreUnimplementedDomainEventHandler struct{}

var _ DomainEventHandlers = (*IgnoreUnimplementedDomainEventHandler)(nil)

func (IgnoreUnimplementedDomainEventHandler) OnStoreCreated(ctx context.Context, event ddd.DomainEvent) error {
	return nil
}

func (IgnoreUnimplementedDomainEventHandler) OnStoreParticipationEnabled(ctx context.Context, event ddd.DomainEvent) error {
	return nil
}

func (IgnoreUnimplementedDomainEventHandler) OnStoreParticipationDisabled(ctx context.Context, event ddd.DomainEvent) error {
	return nil
}

func (IgnoreUnimplementedDomainEventHandler) OnProductAdded(ctx context.Context, event ddd.DomainEvent) error {
	return nil
}

func (IgnoreUnimplementedDomainEventHandler) OnProductRemoved(ctx context.Context, event ddd.DomainEvent) error {
	return nil
}
