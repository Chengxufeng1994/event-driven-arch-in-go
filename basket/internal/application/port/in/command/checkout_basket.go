package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/port/out/client"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/repository"
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
	orderClient      client.OrderClient
}

func NewCheckoutBasketHandler(basketRepository repository.BasketRepository, orderClient client.OrderClient) CheckoutBasketHandler {
	return CheckoutBasketHandler{
		basketRepository: basketRepository,
		orderClient:      orderClient,
	}
}

func (h CheckoutBasketHandler) CheckoutBasket(ctx context.Context, cmd CheckoutBasket) error {
	basket, err := h.basketRepository.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if err := basket.Checkout(cmd.PaymentID); err != nil {
		return errors.Wrap(err, "checkout basket")
	}

	// save the basket
	if err := h.basketRepository.Save(ctx, basket); err != nil {
		return errors.Wrap(err, "saving basket")
	}

	return nil
}
