package handler

import (
	"context"

	customerv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/customer/api/customer/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	orderv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/api/order/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/application/port/out"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/domain"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
)

type IntegrationEventHandlers[T ddd.Event] struct {
	orders    out.OrderRepository
	customers out.CustomerCacheRepository
	products  out.ProductCacheRepository
	stores    out.StoreCacheRepository
}

var _ ddd.EventHandler[ddd.Event] = (*IntegrationEventHandlers[ddd.Event])(nil)

func NewIntegrationEventHandlers(
	orders out.OrderRepository,
	customers out.CustomerCacheRepository,
	products out.ProductCacheRepository,
	stores out.StoreCacheRepository,
) *IntegrationEventHandlers[ddd.Event] {
	return &IntegrationEventHandlers[ddd.Event]{
		orders:    orders,
		customers: customers,
		products:  products,
		stores:    stores,
	}
}

func RegisterIntegrationEventHandlers(subscriber am.EventSubscriber, handlers ddd.EventHandler[ddd.Event]) error {
	evtMsgHandler := am.MessageHandlerFunc[am.IncomingEventMessage](func(ctx context.Context, eventMsg am.IncomingEventMessage) error {
		return handlers.HandleEvent(ctx, eventMsg)
	})

	err := subscriber.Subscribe(customerv1.CustomerAggregateChannel, evtMsgHandler, am.MessageFilter{
		customerv1.CustomerRegisteredEvent,
	}, am.GroupName("search-customers"))
	if err != nil {
		return err
	}

	err = subscriber.Subscribe(orderv1.OrderAggregateChannel, evtMsgHandler, am.MessageFilter{
		orderv1.OrderCreatedEvent,
		orderv1.OrderReadiedEvent,
		orderv1.OrderCanceledEvent,
		orderv1.OrderCompletedEvent,
	}, am.GroupName("search-orders"))
	if err != nil {
		return err
	}

	err = subscriber.Subscribe(storev1.ProductAggregateChannel, evtMsgHandler, am.MessageFilter{
		storev1.ProductAddedEvent,
		storev1.ProductRebrandedEvent,
		storev1.ProductRemovedEvent,
	}, am.GroupName("search-products"))
	if err != nil {
		return err
	}

	err = subscriber.Subscribe(storev1.StoreAggregateChannel, evtMsgHandler, am.MessageFilter{
		storev1.StoreCreatedEvent,
		storev1.StoreRebrandedEvent,
	}, am.GroupName("search-stores"))
	if err != nil {
		return err
	}

	return nil
}

func (h *IntegrationEventHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case customerv1.CustomerRegisteredEvent:
		return h.onCustomerRegistered(ctx, event)
	case orderv1.OrderCreatedEvent:
		return h.onOrderCreated(ctx, event)
	case orderv1.OrderReadiedEvent:
		return h.onOrderReadied(ctx, event)
	case orderv1.OrderCanceledEvent:
		return h.onOrderCanceled(ctx, event)
	case orderv1.OrderCompletedEvent:
		return h.onOrderCompleted(ctx, event)
	case storev1.ProductAddedEvent:
		return h.onProductAdded(ctx, event)
	case storev1.ProductRebrandedEvent:
		return h.onProductRebranded(ctx, event)
	case storev1.ProductRemovedEvent:
		return h.onProductRemoved(ctx, event)
	case storev1.StoreCreatedEvent:
		return h.onStoreCreated(ctx, event)
	case storev1.StoreRebrandedEvent:
		return h.onStoreRebranded(ctx, event)
	}

	return nil
}

func (h IntegrationEventHandlers[T]) onCustomerRegistered(ctx context.Context, event T) error {
	payload := event.Payload().(*customerv1.CustomerRegistered)
	return h.customers.Add(ctx, payload.GetId(), payload.GetName())
}

func (h IntegrationEventHandlers[T]) onOrderCreated(ctx context.Context, event T) error {
	payload := event.Payload().(*orderv1.OrderCreated)

	customer, err := h.customers.Find(ctx, payload.CustomerId)
	if err != nil {
		return err
	}

	var total float64
	items := make([]domain.Item, len(payload.GetItems()))
	seenStores := map[string]*domain.Store{}
	for i, item := range payload.GetItems() {
		product, err := h.products.Find(ctx, item.GetProductId())
		if err != nil {
			return err
		}
		var store *domain.Store
		var exists bool

		if store, exists = seenStores[product.StoreID]; !exists {
			store, err = h.stores.Find(ctx, product.StoreID)
			if err != nil {
				return err
			}
			seenStores[store.ID] = store
		}
		items[i] = domain.Item{
			ProductID:   product.ID,
			StoreID:     store.ID,
			ProductName: product.Name,
			StoreName:   store.Name,
			Price:       item.Price,
			Quantity:    int(item.Quantity),
		}
		total += float64(item.Quantity) * item.Price
	}
	order := &domain.Order{
		OrderID:      payload.GetId(),
		CustomerID:   customer.ID,
		CustomerName: customer.Name,
		Items:        items,
		Total:        total,
		Status:       "New",
	}
	return h.orders.Add(ctx, order)
}

func (h IntegrationEventHandlers[T]) onOrderReadied(ctx context.Context, event T) error {
	payload := event.Payload().(*orderv1.OrderReadied)
	return h.orders.UpdateStatus(ctx, payload.GetId(), "Ready For Pickup")
}

func (h IntegrationEventHandlers[T]) onOrderCanceled(ctx context.Context, event T) error {
	payload := event.Payload().(*orderv1.OrderCanceled)
	return h.orders.UpdateStatus(ctx, payload.GetId(), "Canceled")
}

func (h IntegrationEventHandlers[T]) onOrderCompleted(ctx context.Context, event T) error {
	payload := event.Payload().(*orderv1.OrderCompleted)
	return h.orders.UpdateStatus(ctx, payload.GetId(), "Completed")
}

func (h IntegrationEventHandlers[T]) onProductAdded(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storev1.ProductAdded)
	return h.products.Add(ctx, payload.GetId(), payload.GetStoreId(), payload.GetName())
}

func (h IntegrationEventHandlers[T]) onProductRebranded(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storev1.ProductRebranded)
	return h.products.Rebrand(ctx, payload.GetId(), payload.GetName())
}

func (h IntegrationEventHandlers[T]) onProductRemoved(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storev1.ProductRemoved)
	return h.products.Remove(ctx, payload.GetId())
}

func (h IntegrationEventHandlers[T]) onStoreCreated(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storev1.StoreCreated)
	return h.stores.Add(ctx, payload.GetId(), payload.GetName())
}

func (h IntegrationEventHandlers[T]) onStoreRebranded(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storev1.StoreRebranded)
	return h.stores.Rename(ctx, payload.GetId(), payload.GetName())
}
