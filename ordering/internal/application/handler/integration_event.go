package handler

import (
	"context"

	basketv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/basket/api/basket/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application/port/in/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application/usecase"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/valueobject"
)

type IntegrationEventHandlers[T ddd.Event] struct {
	app usecase.OrderUseCase
}

var _ ddd.EventHandler[ddd.Event] = (*IntegrationEventHandlers[ddd.Event])(nil)

func RegisterIntegrationEventHandler(eventHandler ddd.EventHandler[ddd.Event], subscriber am.EventSubscriber) error {
	evtMsgHandler := am.MessageHandlerFunc[am.IncomingEventMessage](func(ctx context.Context, eventMsg am.IncomingEventMessage) error {
		return eventHandler.HandleEvent(ctx, eventMsg)
	})

	return subscriber.Subscribe(basketv1.BasketAggregateChannel, evtMsgHandler, am.MessageFilter{
		basketv1.BasketCheckedOutEvent,
	}, am.GroupName("ordering-baskets"))
}

func NewIntegrationEventHandlers(app usecase.OrderUseCase) *IntegrationEventHandlers[ddd.Event] {
	return &IntegrationEventHandlers[ddd.Event]{
		app: app,
	}
}

func (h IntegrationEventHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case basketv1.BasketCheckedOutEvent:
		return h.onBasketCheckedOut(ctx, event)
	}
	return nil
}

func (h IntegrationEventHandlers[T]) onBasketCheckedOut(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*basketv1.BasketCheckedOut)
	items := make([]valueobject.Item, 0, len(payload.Items))
	for _, item := range payload.Items {
		items = append(items, valueobject.Item{
			ProductID:   item.GetProductId(),
			StoreID:     item.GetStoreId(),
			StoreName:   item.GetStoreName(),
			ProductName: item.GetProductName(),
			Price:       item.GetPrice(),
			Quantity:    int(item.GetQuantity()),
		})
	}

	cmd := command.CreateOrder{
		ID:         payload.GetId(),
		CustomerID: payload.CustomerId,
		PaymentID:  payload.PaymentId,
		Items:      items,
	}

	return h.app.CreateOrder(ctx, cmd)
}
