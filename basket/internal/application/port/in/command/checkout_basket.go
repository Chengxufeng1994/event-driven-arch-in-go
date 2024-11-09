package command

import (
	"context"

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
	basketRepository repository.BasketRepository
	publisher        ddd.EventPublisher[ddd.Event]
}

func NewCheckoutBasketHandler(basketRepository repository.BasketRepository, publisher ddd.EventPublisher[ddd.Event]) CheckoutBasketHandler {
	return CheckoutBasketHandler{
		basketRepository: basketRepository,
		publisher:        publisher,
	}
}

func (h CheckoutBasketHandler) CheckoutBasket(ctx context.Context, cmd CheckoutBasket) error {
	basket, err := h.basketRepository.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	event, err := basket.Checkout(cmd.PaymentID)
	if err != nil {
		return errors.Wrap(err, "checkout basket")
	}

	// save the basket
	if err := h.basketRepository.Save(ctx, basket); err != nil {
		return errors.Wrap(err, "saving basket")
	}

	return h.publisher.Publish(ctx, event)
}
