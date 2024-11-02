package application

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	orderv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/api/order/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/application/port/out"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/domain"
)

type OrderHandlers[T ddd.Event] struct {
	orders    out.OrderRepository
	customers out.CustomerRepository
	stores    out.StoreRepository
	products  out.ProductRepository
}

var _ ddd.EventHandler[ddd.Event] = (*OrderHandlers[ddd.Event])(nil)

func NewOrderHandlers(orders out.OrderRepository, customers out.CustomerRepository, stores out.StoreRepository, products out.ProductRepository) OrderHandlers[ddd.Event] {
	return OrderHandlers[ddd.Event]{
		orders:    orders,
		customers: customers,
		stores:    stores,
		products:  products,
	}
}

func (h OrderHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case orderv1.OrderCreatedEvent:
		return h.onOrderCreated(ctx, event)
	case orderv1.OrderReadiedEvent:
		return h.onOrderReadied(ctx, event)
	case orderv1.OrderCanceledEvent:
		return h.onOrderCanceled(ctx, event)
	case orderv1.OrderCompletedEvent:
		return h.onOrderCompleted(ctx, event)
	}
	return nil
}

func (h OrderHandlers[T]) onOrderCreated(ctx context.Context, event T) error {
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

func (h OrderHandlers[T]) onOrderReadied(ctx context.Context, event T) error {
	payload := event.Payload().(*orderv1.OrderReadied)
	return h.orders.UpdateStatus(ctx, payload.GetId(), "Ready For Pickup")
}

func (h OrderHandlers[T]) onOrderCanceled(ctx context.Context, event T) error {
	payload := event.Payload().(*orderv1.OrderCanceled)
	return h.orders.UpdateStatus(ctx, payload.GetId(), "Canceled")
}

func (h OrderHandlers[T]) onOrderCompleted(ctx context.Context, event T) error {
	payload := event.Payload().(*orderv1.OrderCompleted)
	return h.orders.UpdateStatus(ctx, payload.GetId(), "Completed")
}
