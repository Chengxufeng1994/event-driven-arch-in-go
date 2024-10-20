package command

import (
	"context"

	"github.com/stackus/errors"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application/port/out/client"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/valueobject"
)

type CreateOrder struct {
	ID         string
	CustomerID string
	PaymentID  string
	Items      []*valueobject.Item
}

type CreateOrderHandler struct {
	order        repository.OrderRepository
	customer     client.CustomerClient
	payment      client.PaymentClient
	shopping     client.ShoppingClient
	notification client.NotificationClient
}

func NewCreateOrderHandler(
	order repository.OrderRepository,
	customer client.CustomerClient,
	payment client.PaymentClient,
	shopping client.ShoppingClient,
	notification client.NotificationClient,
) CreateOrderHandler {
	return CreateOrderHandler{
		order:        order,
		customer:     customer,
		payment:      payment,
		shopping:     shopping,
		notification: notification,
	}
}

func (h CreateOrderHandler) CreateOrder(ctx context.Context, cmd CreateOrder) error {
	orderAgg, err := aggregate.CreateOrder(cmd.ID, cmd.CustomerID, cmd.PaymentID, cmd.Items)
	if err != nil {
		return errors.Wrap(err, "create order command")
	}

	// authorizeCustomer
	if err = h.customer.Authorize(ctx, orderAgg.CustomerID); err != nil {
		return errors.Wrap(err, "order customer authorization")
	}

	// validatePayment
	if err = h.payment.Confirm(ctx, orderAgg.PaymentID); err != nil {
		return errors.Wrap(err, "order payment confirmation")
	}

	// scheduleShopping
	if orderAgg.ShoppingID, err = h.shopping.Create(ctx, orderAgg); err != nil {
		return errors.Wrap(err, "order shopping scheduling")
	}

	// notifyOrderCreated
	if err = h.notification.NotifyOrderCreated(ctx, orderAgg.ID, orderAgg.CustomerID); err != nil {
		return errors.Wrap(err, "order notification")
	}

	return errors.Wrap(h.order.Save(ctx, orderAgg), "create order command")
}
