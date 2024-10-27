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
	Items      []valueobject.Item
}

type CreateOrderHandler struct {
	order    repository.OrderRepository
	customer client.CustomerClient
	payment  client.PaymentClient
	shopping client.ShoppingClient
}

func NewCreateOrderHandler(
	order repository.OrderRepository,
	customer client.CustomerClient,
	payment client.PaymentClient,
	shopping client.ShoppingClient,
) CreateOrderHandler {
	return CreateOrderHandler{
		order:    order,
		customer: customer,
		payment:  payment,
		shopping: shopping,
	}
}

// FIXME:
func (h CreateOrderHandler) CreateOrder(ctx context.Context, cmd CreateOrder) error {
	var err error
	order := aggregate.NewOrder(cmd.ID)

	// authorizeCustomer
	if err := h.customer.Authorize(ctx, cmd.CustomerID); err != nil {
		return errors.Wrap(err, "order customer authorization")
	}

	// validatePayment
	if err = h.payment.Confirm(ctx, cmd.PaymentID); err != nil {
		return errors.Wrap(err, "order payment confirmation")
	}

	// scheduleShopping
	var shoppingID string
	if shoppingID, err = h.shopping.Create(ctx, cmd.ID, cmd.Items); err != nil {
		return errors.Wrap(err, "order shopping scheduling")
	}

	err = order.CreateOrder(cmd.ID, cmd.CustomerID, cmd.PaymentID, shoppingID, cmd.Items)
	if err != nil {
		return errors.Wrap(err, "create order command")
	}

	if err := h.order.Save(ctx, order); err != nil {
		return errors.Wrap(err, "order creation")
	}

	return nil
}
