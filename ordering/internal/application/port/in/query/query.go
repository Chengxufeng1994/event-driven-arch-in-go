package query

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/aggregate"
)

type Queries interface {
	GetOrder(ctx context.Context, query GetOrder) (*aggregate.OrderAgg, error)
}
