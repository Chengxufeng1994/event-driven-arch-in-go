package query

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/aggregate"
)

type Queries interface {
	GetBasket(ctx context.Context, query GetBasket) (*aggregate.BasketAgg, error)
}
