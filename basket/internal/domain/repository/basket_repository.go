package repository

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/aggregate"
)

type BasketRepository interface {
	Save(ctx context.Context, basket *aggregate.Basket) error
	Update(ctx context.Context, basket *aggregate.Basket) error
	Find(ctx context.Context, basketID string) (*aggregate.Basket, error)
}
