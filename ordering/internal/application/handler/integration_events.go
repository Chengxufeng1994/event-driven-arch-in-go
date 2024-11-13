package handler

import (
	"context"

	basketv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/basket/api/basket/v1"
	depotv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/depot/api/depot/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application/port/in/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application/usecase"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/valueobject"
)

type integrationEventHandlers[T ddd.Event] struct {
	app usecase.OrderUseCase
}

var _ ddd.EventHandler[ddd.Event] = (*integrationEventHandlers[ddd.Event])(nil)

func NewIntegrationEventHandlers(app usecase.OrderUseCase) *integrationEventHandlers[ddd.Event] {
	return &integrationEventHandlers[ddd.Event]{
		app: app,
	}
}

func RegisterIntegrationEventHandler(eventHandler ddd.EventHandler[ddd.Event], subscriber am.EventSubscriber) (err error) {
	evtMsgHandler := am.MessageHandlerFunc[am.IncomingEventMessage](func(ctx context.Context, eventMsg am.IncomingEventMessage) error {
		return eventHandler.HandleEvent(ctx, eventMsg)
	})

	_, err = subscriber.Subscribe(basketv1.BasketAggregateChannel, evtMsgHandler, am.MessageFilter{
		basketv1.BasketCheckedOutEvent,
	}, am.GroupName("ordering-baskets"))

	_, err = subscriber.Subscribe(depotv1.ShoppingListAggregateChannel, evtMsgHandler, am.MessageFilter{
		depotv1.ShoppingListCompletedEvent,
	}, am.GroupName("ordering-depot"))

	return err
}

func (h integrationEventHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case basketv1.BasketCheckedOutEvent:
		return h.onBasketCheckedOut(ctx, event)
	case depotv1.ShoppingListCompletedEvent:
		return h.onShoppingListCompleted(ctx, event)
	}
	return nil
}

func (h integrationEventHandlers[T]) onBasketCheckedOut(ctx context.Context, event ddd.Event) error {
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

func (h integrationEventHandlers[T]) onShoppingListCompleted(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*depotv1.ShoppingListCompleted)

	return h.app.ReadyOrder(ctx, command.ReadyOrder{ID: payload.GetOrderId()})
}
