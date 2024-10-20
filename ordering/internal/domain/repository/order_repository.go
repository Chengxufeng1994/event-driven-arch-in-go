package repository

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/aggregate"
)

type OrderRepository interface {
	Find(ctx context.Context, orderID string) (*aggregate.OrderAgg, error)
	Save(ctx context.Context, order *aggregate.OrderAgg) error
	Update(ctx context.Context, order *aggregate.OrderAgg) error
}