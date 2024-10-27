package command

import (
	"context"

	"github.com/stackus/errors"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
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
	order                repository.OrderRepository
	customer             client.CustomerClient
	payment              client.PaymentClient
	shopping             client.ShoppingClient
	domainEventPublisher ddd.EventPublisher
}

func NewCreateOrderHandler(
	order repository.OrderRepository,
	customer client.CustomerClient,
	payment client.PaymentClient,
	shopping client.ShoppingClient,
	domainEventPublisher ddd.EventPublisher,
) CreateOrderHandler {
	return CreateOrderHandler{
		order:                order,
		customer:             customer,
		payment:              payment,
		shopping:             shopping,
		domainEventPublisher: domainEventPublisher,
	}
}

func (h CreateOrderHandler) CreateOrder(ctx context.Context, cmd CreateOrder) error {
	order, err := aggregate.CreateOrder(cmd.ID, cmd.CustomerID, cmd.PaymentID, cmd.Items)
	if err != nil {
		return errors.Wrap(err, "create order command")
	}

	// authorizeCustomer
	if err = h.customer.Authorize(ctx, order.CustomerID); err != nil {
		return errors.Wrap(err, "order customer authorization")
	}

	// validatePayment
	if err = h.payment.Confirm(ctx, order.PaymentID); err != nil {
		return errors.Wrap(err, "order payment confirmation")
	}

	// scheduleShopping
	if order.ShoppingID, err = h.shopping.Create(ctx, order); err != nil {
		return errors.Wrap(err, "order shopping scheduling")
	}

	if err := h.order.Save(ctx, order); err != nil {
		return errors.Wrap(err, "order creation")
	}

	if err := h.domainEventPublisher.Publish(ctx, order.GetEvents()...); err != nil {
		return errors.Wrap(err, "publish domain events")
	}

	return nil
}
