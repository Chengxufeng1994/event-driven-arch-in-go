package handler

import (
	"context"

	basketv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/basket/api/basket/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/aggregate"
	domainevent "github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/event"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
)

type DomainEventHandler[T ddd.Event] struct {
	publisher am.MessagePublisher[ddd.Event]
}

var _ ddd.EventHandler[ddd.Event] = (*DomainEventHandler[ddd.Event])(nil)

func NewDomainEventHandler(publisher am.MessagePublisher[ddd.Event]) *DomainEventHandler[ddd.Event] {
	return &DomainEventHandler[ddd.Event]{
		publisher: publisher,
	}
}

func RegisterDomainEventHandlers(subscriber ddd.EventSubscriber[ddd.Event], handlers ddd.EventHandler[ddd.Event]) {
	subscriber.Subscribe(handlers,
		domainevent.BasketStartedEvent,
		domainevent.BasketCheckedOutEvent,
	)
}

func (h *DomainEventHandler[T]) HandleEvent(ctx context.Context, event ddd.Event) error {
	switch event.EventName() {
	case domainevent.BasketStartedEvent:
		return h.onBasketStarted(ctx, event)
	case domainevent.BasketCanceledEvent:
		return h.onBasketCanceled(ctx, event)
	case domainevent.BasketCheckedOutEvent:
		return h.onBasketCheckedOut(ctx, event)
	}
	return nil
}

func (h *DomainEventHandler[T]) onBasketStarted(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*aggregate.Basket)
	return h.publisher.Publish(ctx, basketv1.BasketAggregateChannel,
		ddd.NewEvent(basketv1.BasketStartedEvent, &basketv1.BasketStarted{
			Id:         payload.ID(),
			CustomerId: payload.CustomerID,
		}),
	)
}

func (h *DomainEventHandler[T]) onBasketCanceled(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*aggregate.Basket)
	return h.publisher.Publish(ctx, basketv1.BasketAggregateChannel,
		ddd.NewEvent(basketv1.BasketCanceledEvent, &basketv1.BasketCanceled{
			Id: payload.ID(),
		}),
	)
}

func (h *DomainEventHandler[T]) onBasketCheckedOut(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*aggregate.Basket)
	items := make([]*basketv1.BasketCheckedOut_Item, 0, len(payload.Items))
	for _, item := range payload.Items {
		items = append(items, &basketv1.BasketCheckedOut_Item{
			StoreId:     item.StoreID,
			ProductId:   item.ProductID,
			StoreName:   item.StoreName,
			ProductName: item.ProductName,
			Price:       item.ProductPrice,
			Quantity:    int32(item.Quantity),
		})
	}
	return h.publisher.Publish(ctx, basketv1.BasketAggregateChannel,
		ddd.NewEvent(basketv1.BasketCheckedOutEvent, &basketv1.BasketCheckedOut{
			Id:         payload.ID(),
			CustomerId: payload.CustomerID,
			PaymentId:  payload.PaymentID,
			Items:      items,
		}),
	)
}
