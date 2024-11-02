package out

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/domain"
)

type StoreRepository interface {
	Find(ctx context.Context, storeID string) (*domain.Store, error)
}
