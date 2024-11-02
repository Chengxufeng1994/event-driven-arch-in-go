package usecase

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/application/port/in/query"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/domain"
)

type (
	SearchUseCase interface {
		SearchOrders(ctx context.Context, search query.SearchOrders) ([]*domain.Order, error)
		GetOrder(ctx context.Context, get query.GetOrder) (*domain.Order, error)
	}
)
