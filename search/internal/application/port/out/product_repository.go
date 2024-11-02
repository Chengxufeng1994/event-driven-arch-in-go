package out

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/domain"
)

type ProductRepository interface {
	Find(ctx context.Context, productID string) (*domain.Product, error)
}
