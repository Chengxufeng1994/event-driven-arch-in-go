package client

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/valueobject"
)

type ProductClient interface {
	Find(ctx context.Context, productID string) (valueobject.Product, error)
}