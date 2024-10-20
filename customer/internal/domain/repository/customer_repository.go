package repository

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/domain/aggregate"
)

type CustomerRepository interface {
	Save(ctx context.Context, customer *aggregate.CustomerAgg) error
	Find(ctx context.Context, customerID string) (*aggregate.CustomerAgg, error)
	Update(ctx context.Context, customer *aggregate.CustomerAgg) error
}
