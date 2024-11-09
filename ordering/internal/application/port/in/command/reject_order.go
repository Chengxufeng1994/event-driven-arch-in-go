package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/repository"
)

type RejectOrder struct {
	ID string
}

func NewRejectOrder(id string) RejectOrder {
	return RejectOrder{ID: id}
}

func (cmd RejectOrder) Validate() error {
	return nil
}

type RejectOrderHandler struct {
	orderRepository repository.OrderRepository
	publisher       ddd.EventPublisher[ddd.Event]
}

func NewRejectOrderHandler(orderRepository repository.OrderRepository, publisher ddd.EventPublisher[ddd.Event]) RejectOrderHandler {
	return RejectOrderHandler{
		orderRepository: orderRepository,
		publisher:       publisher}
}

func (h RejectOrderHandler) RejectOrder(ctx context.Context, cmd RejectOrder) error {
	order, err := h.orderRepository.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	event, err := order.Reject()
	if err != nil {
		return err
	}

	if err = h.orderRepository.Save(ctx, order); err != nil {
		return err
	}

	return h.publisher.Publish(ctx, event)
}
