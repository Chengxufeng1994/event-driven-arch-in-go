package out

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/application/port/in/query"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/domain"
)

type OrderRepository interface {
	Add(ctx context.Context, order *domain.Order) error
	UpdateStatus(ctx context.Context, orderID, status string) error
	Search(ctx context.Context, search query.SearchOrders) ([]*domain.Order, error)
	Get(ctx context.Context, orderID string) (*domain.Order, error)
}
