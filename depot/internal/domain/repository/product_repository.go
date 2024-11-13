package repository

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/entity"
)

type ProductRepository interface {
	Find(ctx context.Context, productID string) (*entity.Product, error)
}
