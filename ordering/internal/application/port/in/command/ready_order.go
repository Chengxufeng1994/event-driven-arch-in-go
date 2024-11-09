package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/repository"
	"github.com/stackus/errors"
)

type ReadyOrder struct {
	ID string
}

type ReadyOrderHandler struct {
	orderRepository repository.OrderRepository
	publisher       ddd.EventPublisher[ddd.Event]
}

func NewReadyOrderHandler(
	orderRepository repository.OrderRepository,
	publisher ddd.EventPublisher[ddd.Event],
) ReadyOrderHandler {
	return ReadyOrderHandler{
		orderRepository: orderRepository,
		publisher:       publisher,
	}
}

func (h ReadyOrderHandler) ReadyOrder(ctx context.Context, cmd ReadyOrder) error {
	orderAgg, err := h.orderRepository.Load(ctx, cmd.ID)
	if err != nil {
		return errors.Wrap(err, "ready order command")
	}

	event, err := orderAgg.Ready()
	if err != nil {
		return nil
	}

	if err := h.orderRepository.Save(ctx, orderAgg); err != nil {
		return errors.Wrap(err, "order update")
	}

	return h.publisher.Publish(ctx, event)
}
