package repository

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
)

type ParticipatingStoreRepository interface {
	FindAll(ctx context.Context) ([]*aggregate.Store, error)
}
