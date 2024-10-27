package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/repository"
	"github.com/stackus/errors"
)

type ReadyOrder struct {
	ID string
}

type ReadyOrderHandler struct {
	orderRepository repository.OrderRepository
}

func NewReadyOrderHandler(
	orderRepository repository.OrderRepository,
) ReadyOrderHandler {
	return ReadyOrderHandler{
		orderRepository: orderRepository,
	}
}

func (h ReadyOrderHandler) ReadyOrder(ctx context.Context, cmd ReadyOrder) error {
	orderAgg, err := h.orderRepository.Load(ctx, cmd.ID)
	if err != nil {
		return errors.Wrap(err, "ready order command")
	}

	if err = orderAgg.Ready(); err != nil {
		return nil
	}

	if err := h.orderRepository.Save(ctx, orderAgg); err != nil {
		return errors.Wrap(err, "order update")
	}

	return nil
}
