package query

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/domain/aggregate"
)

type Queries interface {
	GetCustomer(ctx context.Context, query GetCustomer) (*aggregate.CustomerAgg, error)
}
