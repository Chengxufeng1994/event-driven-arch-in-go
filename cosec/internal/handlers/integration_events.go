package handlers

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/cosec/internal/models"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/sec"
	orderv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/api/order/v1"
)

type integrationHandlers[T ddd.Event] struct {
	orchestrator sec.Orchestrator[*models.CreateOrderData]
}

var _ ddd.EventHandler[ddd.Event] = (*integrationHandlers[ddd.Event])(nil)

func NewIntegrationEventHandlers(saga sec.Orchestrator[*models.CreateOrderData]) ddd.EventHandler[ddd.Event] {
	return integrationHandlers[ddd.Event]{
		orchestrator: saga,
	}
}

func RegisterIntegrationEventHandlers(subscriber am.RawMessageStream, handlers am.RawMessageHandler) (err error) {
	return subscriber.Subscribe(orderv1.OrderAggregateChannel, handlers, am.MessageFilter{
		orderv1.OrderCreatedEvent,
	}, am.GroupName("cosec-ordering"))
}

func (h integrationHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case orderv1.OrderCreatedEvent:
		return h.onOrderCreated(ctx, event)
	}

	return nil
}

func (h integrationHandlers[T]) onOrderCreated(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*orderv1.OrderCreated)

	var total float64
	items := make([]models.Item, len(payload.GetItems()))
	for i, item := range payload.GetItems() {
		items[i] = models.Item{
			ProductID: item.GetProductId(),
			StoreID:   item.GetStoreId(),
			Price:     item.GetPrice(),
			Quantity:  int(item.GetQuantity()),
		}
		total += float64(item.GetQuantity()) * item.GetPrice()
	}

	data := &models.CreateOrderData{
		OrderID:    payload.GetId(),
		CustomerID: payload.GetCustomerId(),
		PaymentID:  payload.GetPaymentId(),
		Items:      items,
		Total:      total,
	}

	// Start the CreateOrderSaga
	return h.orchestrator.Start(ctx, event.ID(), data)
}
