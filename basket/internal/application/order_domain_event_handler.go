package application

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/port/in/event"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/port/out/client"
	domainevent "github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/event"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
)

type (
	OrderDomainEventHandler[T ddd.AggregateEvent] struct {
		orderClient client.OrderClient
	}

	// appDomainEventHandlers struct {
	// 	event.OrderEventHandler
	// }
)

var _ interface {
	ddd.EventHandler[ddd.AggregateEvent]
	event.OrderEventHandler
} = (*OrderDomainEventHandler[ddd.AggregateEvent])(nil)

func NewOrderDomainEventHandler(orderClient client.OrderClient) *OrderDomainEventHandler[ddd.AggregateEvent] {
	return &OrderDomainEventHandler[ddd.AggregateEvent]{
		orderClient: orderClient,
	}
}

func (h *OrderDomainEventHandler[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case domainevent.BasketCheckedOutEvent:
		return h.OnBasketCheckedOut(ctx, event)
	default:
		return nil
	}
}

func (h *OrderDomainEventHandler[T]) OnBasketCheckedOut(ctx context.Context, event ddd.AggregateEvent) error {
	checkedOut := event.Payload().(*domainevent.BasketCheckedOut)
	_, err := h.orderClient.Save(ctx, checkedOut.PaymentID, checkedOut.CustomerID, checkedOut.Items)
	return err
}
