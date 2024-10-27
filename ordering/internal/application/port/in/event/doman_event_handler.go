package event

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
)

type DomainEventHandlers interface {
	OnOrderCreated(ctx context.Context, event ddd.DomainEvent) error
	OnOrderReadied(ctx context.Context, event ddd.DomainEvent) error
	OnOrderCanceled(ctx context.Context, event ddd.DomainEvent) error
	OnOrderCompleted(ctx context.Context, event ddd.DomainEvent) error
}
