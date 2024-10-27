package client

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/valueobject"
)

type ShoppingClient interface {
	Create(ctx context.Context, orderID string, items []valueobject.Item) (string, error)
	Cancel(ctx context.Context, shoppingID string) error
}
