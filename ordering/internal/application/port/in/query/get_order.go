package query

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/repository"
)

type GetOrder struct {
	OrderID string
}

type GetOrderHandler struct {
	orderRepository repository.OrderRepository
}

func NewGetOrderHandler(orderRepository repository.OrderRepository) GetOrderHandler {
	return GetOrderHandler{orderRepository: orderRepository}
}

func (h GetOrderHandler) GetOrder(ctx context.Context, query GetOrder) (*aggregate.OrderAgg, error) {
	return h.orderRepository.Find(ctx, query.OrderID)
}
