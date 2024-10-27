package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application/port/out/client"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/repository"
	"github.com/stackus/errors"
)

type CancelOrder struct {
	ID string
}

type CancelOrderHandler struct {
	orderRepository repository.OrderRepository
	shopping        client.ShoppingClient
}

func NewCancelOrderHandler(
	orderRepository repository.OrderRepository,
	shopping client.ShoppingClient,
) CancelOrderHandler {
	return CancelOrderHandler{
		orderRepository: orderRepository,
		shopping:        shopping,
	}
}

func (h CancelOrderHandler) CancelOrder(ctx context.Context, cmd CancelOrder) error {
	orderAgg, err := h.orderRepository.Load(ctx, cmd.ID)
	if err != nil {
		return errors.Wrap(err, "cancel order command")
	}

	if err = orderAgg.Cancel(); err != nil {
		return errors.Wrap(err, "cancel order command")
	}

	if err = h.shopping.Cancel(ctx, orderAgg.ShoppingID); err != nil {
		return errors.Wrap(err, "order shopping cancel")
	}

	if err = h.orderRepository.Save(ctx, orderAgg); err != nil {
		return errors.Wrap(err, "saving order")
	}

	return nil
}
