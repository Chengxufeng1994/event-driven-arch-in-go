package event

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
)

type DomainEventHandlers interface {
	OnShoppingListCreated(ctx context.Context, event ddd.DomainEvent) error
	OnShoppingListCanceled(ctx context.Context, event ddd.DomainEvent) error
	OnShoppingListAssigned(ctx context.Context, event ddd.DomainEvent) error
	OnShoppingListCompleted(ctx context.Context, event ddd.DomainEvent) error
}
