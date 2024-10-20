package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application/port/out/client"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/repository"
	"github.com/stackus/errors"
)

type ReadyOrder struct {
	ID string
}

type ReadyOrderHandler struct {
	orderRepository repository.OrderRepository
	invoiceClient   client.InvoiceClient
	notification    client.NotificationClient
}

func NewReadyOrderHandler(
	orderRepository repository.OrderRepository,
	invoiceClient client.InvoiceClient,
	notification client.NotificationClient,
) ReadyOrderHandler {
	return ReadyOrderHandler{
		orderRepository: orderRepository,
		invoiceClient:   invoiceClient,
		notification:    notification,
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
		return errors.Wrap(err, "ready order command")
	}

	if err := h.notification.NotifyOrderReady(ctx, orderAgg.ID, orderAgg.CustomerID); err != nil {
		return errors.Wrap(err, "order notification")
	}

	return h.orderRepository.Update(ctx, orderAgg)
}
