package event

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
)

type DomainEventHandlers interface {
	OnCustomerRegistered(ctx context.Context, event ddd.DomainEvent) error
	OnCustomerAuthorized(ctx context.Context, event ddd.DomainEvent) error
	OnCustomerEnabled(ctx context.Context, event ddd.DomainEvent) error
	OnCustomerDisabled(ctx context.Context, event ddd.DomainEvent) error
}
