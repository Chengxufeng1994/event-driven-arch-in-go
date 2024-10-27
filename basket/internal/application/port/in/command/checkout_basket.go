package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/port/out/client"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/stackus/errors"
)

type CheckoutBasket struct {
	ID        string
	PaymentID string
}

func NewCheckoutBasket(id, paymentID string) CheckoutBasket {
	return CheckoutBasket{
		ID:        id,
		PaymentID: paymentID,
	}
}

type CheckoutBasketHandler struct {
	basketRepository     repository.BasketRepository
	orderClient          client.OrderClient
	domainEventPublisher ddd.EventPublisher
}

func NewCheckoutBasketHandler(basketRepository repository.BasketRepository, orderClient client.OrderClient, domainEventPublisher ddd.EventPublisher) CheckoutBasketHandler {
	return CheckoutBasketHandler{
		basketRepository:     basketRepository,
		orderClient:          orderClient,
		domainEventPublisher: domainEventPublisher,
	}
}

func (h CheckoutBasketHandler) CheckoutBasket(ctx context.Context, cmd CheckoutBasket) error {
	basket, err := h.basketRepository.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if err := basket.Checkout(cmd.PaymentID); err != nil {
		return errors.Wrap(err, "checkout basket")
	}

	// update the basket
	if err := h.basketRepository.Update(ctx, basket); err != nil {
		return errors.Wrap(err, "updating basket")
	}

	// publish domain events
	return h.domainEventPublisher.Publish(ctx, basket.GetEvents()...)
}
