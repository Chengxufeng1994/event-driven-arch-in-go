package event

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
)

type DomainEventHandlers interface {
	OnBasketStarted(ctx context.Context, event ddd.DomainEvent) error
	OnBasketItemAdded(ctx context.Context, event ddd.DomainEvent) error
	OnBasketItemRemoved(ctx context.Context, event ddd.DomainEvent) error
	OnBasketCanceled(ctx context.Context, event ddd.DomainEvent) error
	OnBasketCheckedOut(ctx context.Context, event ddd.DomainEvent) error
}
