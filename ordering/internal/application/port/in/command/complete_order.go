package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/repository"
	"github.com/stackus/errors"
)

type CompleteOrder struct {
	ID        string
	InvoiceID string
}

type CompleteOrderHandler struct {
	orderRepository repository.OrderRepository
	publisher       ddd.EventPublisher[ddd.Event]
}

func NewCompleteOrderHandler(orderRepository repository.OrderRepository, publisher ddd.EventPublisher[ddd.Event]) CompleteOrderHandler {
	return CompleteOrderHandler{orderRepository: orderRepository, publisher: publisher}
}

func (h CompleteOrderHandler) CompleteOrder(ctx context.Context, cmd CompleteOrder) error {
	orderAgg, err := h.orderRepository.Load(ctx, cmd.ID)
	if err != nil {
		return errors.Wrap(err, "complete order command")
	}

	event, err := orderAgg.Complete(cmd.InvoiceID)
	if err != nil {
		return errors.Wrap(err, "complete order command")
	}

	err = h.orderRepository.Save(ctx, orderAgg)
	if err != nil {
		return errors.Wrap(err, "complete order command")
	}

	return h.publisher.Publish(ctx, event)
}
