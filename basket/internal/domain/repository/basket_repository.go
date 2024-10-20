package repository

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/aggregate"
)

type BasketRepository interface {
	Save(ctx context.Context, basket *aggregate.BasketAgg) error
	Update(ctx context.Context, basket *aggregate.BasketAgg) error
	Find(ctx context.Context, basketID string) (*aggregate.BasketAgg, error)
}
