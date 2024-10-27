package event

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/port/out/client"
	domainevent "github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/event"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
)

type OnBasketCheckOutEventHandler struct {
	orderClient client.OrderClient
}

func NewOnBasketCheckOutEventHandler(orderClient client.OrderClient) *OnBasketCheckOutEventHandler {
	return &OnBasketCheckOutEventHandler{
		orderClient: orderClient,
	}
}

func (h *OnBasketCheckOutEventHandler) OnBasketCheckedOut(ctx context.Context, event ddd.AggregateEvent) error {
	checkout := event.Payload().(*domainevent.BasketCheckedOut)
	_, err := h.orderClient.Save(ctx, checkout.PaymentID, checkout.CustomerID, checkout.Items)
	if err != nil {
		return err
	}
	return nil
}
