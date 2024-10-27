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
	orderRepository      repository.OrderRepository
	domainEventPublisher ddd.EventPublisher
}

func NewReadyOrderHandler(
	orderRepository repository.OrderRepository,
	domainEventPublisher ddd.EventPublisher,
) ReadyOrderHandler {
	return ReadyOrderHandler{
		orderRepository:      orderRepository,
		domainEventPublisher: domainEventPublisher,
	}
}

func (h ReadyOrderHandler) ReadyOrder(ctx context.Context, cmd ReadyOrder) error {
	orderAgg, err := h.orderRepository.Find(ctx, cmd.ID)
	if err != nil {
		return errors.Wrap(err, "ready order command")
	}

	if err = orderAgg.Ready(); err != nil {
		return nil
	}

	if err := h.orderRepository.Update(ctx, orderAgg); err != nil {
		return errors.Wrap(err, "order update")
	}

	// publish domain events
	if err := h.domainEventPublisher.Publish(ctx, orderAgg.GetEvents()...); err != nil {
		return errors.Wrap(err, "publish domain events")
	}

	return nil
}
