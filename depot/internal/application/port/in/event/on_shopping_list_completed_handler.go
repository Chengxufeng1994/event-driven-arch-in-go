package event

import (
	"context"
	"fmt"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/application/port/out/client"
	domainevent "github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/event"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
)

type OnShoppingListCompletedHandler struct {
	orderClient client.OrderClient
}

func NewOnShoppingListCompletedHandler(orderClient client.OrderClient) OnShoppingListCompletedHandler {
	return OnShoppingListCompletedHandler{
		orderClient: orderClient,
	}
}

func (h *OnShoppingListCompletedHandler) OnShoppingListCompleted(ctx context.Context, event ddd.Event) error {
	shoppingListCompleted, ok := event.Payload().(*domainevent.ShoppingListCompleted)
	if !ok {
		return fmt.Errorf("unexpected event type: %T", event)
	}
	return h.orderClient.Ready(ctx, shoppingListCompleted.OrderID)
}
