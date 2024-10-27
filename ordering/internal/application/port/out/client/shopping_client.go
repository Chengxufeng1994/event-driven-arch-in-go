package client

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/aggregate"
)

type ShoppingClient interface {
	Create(ctx context.Context, order *aggregate.Order) (string, error)
	Cancel(ctx context.Context, shoppingID string) error
}
