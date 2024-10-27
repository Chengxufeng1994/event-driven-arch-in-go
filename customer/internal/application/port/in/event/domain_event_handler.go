package event

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
)

type DomainEventHandlers interface {
	OnCustomerRegistered(ctx context.Context, event ddd.Event) error
	OnCustomerAuthorized(ctx context.Context, event ddd.Event) error
	OnCustomerEnabled(ctx context.Context, event ddd.Event) error
	OnCustomerDisabled(ctx context.Context, event ddd.Event) error
}
