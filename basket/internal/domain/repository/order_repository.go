package repository

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/aggregate"
)

type OrderRepository interface {
	Save(ctx context.Context, basket *aggregate.Basket) (string, error)
}
