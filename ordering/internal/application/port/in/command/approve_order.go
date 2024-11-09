package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/repository"
)

type ApproveOrder struct {
	ID         string
	ShoppingID string
}

func NewApproveOrder(id string, shoppingID string) ApproveOrder {
	return ApproveOrder{
		ID:         id,
		ShoppingID: shoppingID,
	}
}

func (ApproveOrder) Validate() error {
	return nil
}

type ApproveOrderHandler struct {
	orderRepository repository.OrderRepository
	publisher       ddd.EventPublisher[ddd.Event]
}

func NewApproveOrderHandler(orderRepository repository.OrderRepository, publisher ddd.EventPublisher[ddd.Event]) ApproveOrderHandler {
	return ApproveOrderHandler{
		orderRepository: orderRepository,
		publisher:       publisher,
	}
}

func (h ApproveOrderHandler) ApproveOrder(ctx context.Context, cmd ApproveOrder) error {
	order, err := h.orderRepository.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	event, err := order.Approve(cmd.ShoppingID)
	if err != nil {
		return err
	}

	if err = h.orderRepository.Save(ctx, order); err != nil {
		return err
	}

	return h.publisher.Publish(ctx, event)
}
