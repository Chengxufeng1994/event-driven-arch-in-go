package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
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
	publisher       ddd.EventPublisher[ddd.Event]
}

func NewCancelOrderHandler(
	orderRepository repository.OrderRepository,
	shopping client.ShoppingClient,
	publisher ddd.EventPublisher[ddd.Event],
) CancelOrderHandler {
	return CancelOrderHandler{
		orderRepository: orderRepository,
		shopping:        shopping,
		publisher:       publisher,
	}
}

func (h CancelOrderHandler) CancelOrder(ctx context.Context, cmd CancelOrder) error {
	orderAgg, err := h.orderRepository.Load(ctx, cmd.ID)
	if err != nil {
		return errors.Wrap(err, "cancel order command")
	}

	event, err := orderAgg.Cancel()
	if err != nil {
		return errors.Wrap(err, "cancel order command")
	}

	// TODO CH8 remove; handled in the saga
	if err = h.shopping.Cancel(ctx, orderAgg.ShoppingID); err != nil {
		return errors.Wrap(err, "order shopping cancel")
	}

	if err = h.orderRepository.Save(ctx, orderAgg); err != nil {
		return errors.Wrap(err, "saving order")
	}

	return h.publisher.Publish(ctx, event)
}
