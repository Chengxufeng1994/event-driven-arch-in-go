package repository

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/entity"
)

type StoreRepository interface {
	Find(ctx context.Context, storeID string) (*entity.Store, error)
}
