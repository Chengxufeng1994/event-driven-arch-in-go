package application

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/application/port/in/query"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/application/port/out"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/application/usecase"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/domain"
)

type SearchApplication struct {
	orderRepository out.OrderRepository
}

var _ usecase.SearchUseCase = (*SearchApplication)(nil)

func New(orderRepository out.OrderRepository) *SearchApplication {
	return &SearchApplication{
		orderRepository: orderRepository,
	}
}

func (s *SearchApplication) SearchOrders(ctx context.Context, search query.SearchOrders) ([]*domain.Order, error) {
	panic("unimplemented")
}

func (s *SearchApplication) GetOrder(ctx context.Context, get query.GetOrder) (*domain.Order, error) {
	panic("unimplemented")
}
