package event

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
)

type DomainEventHandlers interface {
	OnOrderCreated(ctx context.Context, event ddd.Event) error
	OnOrderReadied(ctx context.Context, event ddd.Event) error
	OnOrderCanceled(ctx context.Context, event ddd.Event) error
	OnOrderCompleted(ctx context.Context, event ddd.Event) error
}
