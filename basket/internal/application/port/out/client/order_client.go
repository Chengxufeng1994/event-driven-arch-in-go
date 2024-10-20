package client

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/aggregate"
)

type OrderClient interface {
	Save(ctx context.Context, basket *aggregate.BasketAgg) (string, error)
}
