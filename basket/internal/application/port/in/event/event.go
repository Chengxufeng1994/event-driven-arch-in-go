package event

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
)

type OrderEventHandler interface {
	OnBasketCheckedOut(ctx context.Context, event ddd.AggregateEvent) error
}
