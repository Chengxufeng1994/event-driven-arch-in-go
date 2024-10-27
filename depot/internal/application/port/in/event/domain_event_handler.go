package event

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
)

type DomainEventHandlers interface {
	OnShoppingListCreated(ctx context.Context, event ddd.Event) error
	OnShoppingListCanceled(ctx context.Context, event ddd.Event) error
	OnShoppingListAssigned(ctx context.Context, event ddd.Event) error
	OnShoppingListCompleted(ctx context.Context, event ddd.Event) error
}
