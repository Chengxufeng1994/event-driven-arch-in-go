package command

import (
	"context"

	"github.com/stackus/errors"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/valueobject"
)

type CreateOrder struct {
	ID         string
	CustomerID string
	PaymentID  string
	Items      []valueobject.Item
}

type CreateOrderHandler struct {
	orderRepository repository.OrderRepository
	publisher       ddd.EventPublisher[ddd.Event]
}

func NewCreateOrderHandler(
	order repository.OrderRepository,
	publisher ddd.EventPublisher[ddd.Event],
) CreateOrderHandler {
	return CreateOrderHandler{
		orderRepository: order,
		publisher:       publisher,
	}
}

func (h CreateOrderHandler) CreateOrder(ctx context.Context, cmd CreateOrder) error {
	order, err := h.orderRepository.Load(ctx, cmd.ID)
	if err != nil {
		return errors.Wrap(err, "create order command")
	}

	event, err := order.CreateOrder(cmd.ID, cmd.CustomerID, cmd.PaymentID, cmd.Items)
	if err != nil {
		return errors.Wrap(err, "create order command")
	}

	if err := h.orderRepository.Save(ctx, order); err != nil {
		return errors.Wrap(err, "order creation")
	}

	return h.publisher.Publish(ctx, event)
}
