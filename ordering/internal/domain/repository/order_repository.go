package repository

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/aggregate"
)

type OrderRepository interface {
	Load(ctx context.Context, orderID string) (*aggregate.Order, error)
	Save(ctx context.Context, order *aggregate.Order) error
}
