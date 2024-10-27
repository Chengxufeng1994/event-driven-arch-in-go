package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/repository"
	"github.com/stackus/errors"
)

type CompleteOrder struct {
	ID        string
	InvoiceID string
}

type CompleteOrderHandler struct {
	orderRepository repository.OrderRepository
}

func NewCompleteOrderHandler(orderRepository repository.OrderRepository) CompleteOrderHandler {
	return CompleteOrderHandler{orderRepository: orderRepository}
}

func (h CompleteOrderHandler) CompleteOrder(ctx context.Context, cmd CompleteOrder) error {
	orderAgg, err := h.orderRepository.Load(ctx, cmd.ID)
	if err != nil {
		return errors.Wrap(err, "complete order command")
	}

	if err := orderAgg.Complete(cmd.InvoiceID); err != nil {
		return errors.Wrap(err, "complete order command")
	}

	return errors.Wrap(h.orderRepository.Save(ctx, orderAgg), "complete order command")
}
