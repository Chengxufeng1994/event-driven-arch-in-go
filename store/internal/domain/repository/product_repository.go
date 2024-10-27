package repository

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
)

type ProductRepository interface {
	Load(ctx context.Context, id string) (*aggregate.Product, error)
	Save(ctx context.Context, product *aggregate.Product) error
}
