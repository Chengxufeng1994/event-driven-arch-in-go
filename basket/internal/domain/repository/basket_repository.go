package repository

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/aggregate"
)

type BasketRepository interface {
	Load(ctx context.Context, basketID string) (*aggregate.Basket, error)
	Save(ctx context.Context, basket *aggregate.Basket) error
}
