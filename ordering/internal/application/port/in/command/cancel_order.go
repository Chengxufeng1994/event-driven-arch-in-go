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
	notification    client.NotificationClient
}

func NewCancelOrderHandler(
	orderRepository repository.OrderRepository,
	shopping client.ShoppingClient,
	notification client.NotificationClient,
) CancelOrderHandler {
	return CancelOrderHandler{
		orderRepository: orderRepository,
		shopping:        shopping,
		notification:    notification,
	}
}

func (h CancelOrderHandler) CancelOrder(ctx context.Context, cmd CancelOrder) error {
	orderAgg, err := h.orderRepository.Find(ctx, cmd.ID)
	if err != nil {
		return errors.Wrap(err, "cancel order command")
	}

	if err = orderAgg.Cancel(); err != nil {
		return errors.Wrap(err, "cancel order command")
	}

	if err = h.shopping.Cancel(ctx, orderAgg.ShoppingID); err != nil {
		return errors.Wrap(err, "order shopping cancel")
	}

	if err = h.notification.NotifyOrderCanceled(ctx, orderAgg.ID, orderAgg.CustomerID); err != nil {
		return errors.Wrap(err, "order notification")
	}

	return errors.Wrap(h.orderRepository.Update(ctx, orderAgg), "cancel order command")
}
