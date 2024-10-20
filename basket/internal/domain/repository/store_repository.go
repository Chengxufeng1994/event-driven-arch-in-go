package repository

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/valueobject"
)

type StoreRepository interface {
	Find(ctx context.Context, storeID string) (valueobject.Store, error)
}
