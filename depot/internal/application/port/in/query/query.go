package query

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/aggregate"
)

type Queries interface {
	GetShoppingList(ctx context.Context, query GetShoppingList) (*aggregate.ShoppingList, error)
}
