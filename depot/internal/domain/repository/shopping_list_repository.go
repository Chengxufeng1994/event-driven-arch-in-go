package repository

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/aggregate"
)

type ShoppingListRepository interface {
	Save(ctx context.Context, list *aggregate.ShoppingListAgg) error
	Update(ctx context.Context, list *aggregate.ShoppingListAgg) error
	Find(ctx context.Context, shoppingListID string) (*aggregate.ShoppingListAgg, error)
	FindByOrderID(ctx context.Context, orderID string) (*aggregate.ShoppingListAgg, error)
}