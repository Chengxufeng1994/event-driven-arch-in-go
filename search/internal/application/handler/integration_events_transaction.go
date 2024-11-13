package handler

import (
	"context"

	customerv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/customer/api/customer/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/di"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"
	orderv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/api/order/v1"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
	"gorm.io/gorm"
)

func RegisterIntegrationEventHandlersTx(container di.Container) error {
	evtMsgHandler := am.RawMessageHandlerFunc(func(ctx context.Context, msg am.IncomingRawMessage) (err error) {
		ctx = container.Scoped(ctx)
		defer func(tx *gorm.DB) {
			if p := recover(); p != nil {
				_ = tx.Rollback()
				panic(p)
			} else if err != nil {
				_ = tx.Rollback()
			} else {
				err = tx.Commit().Error
			}
		}(di.Get(ctx, "tx").(*gorm.DB))

		evtHandlers := am.RawMessageHandlerWithMiddleware(
			am.NewEventMessageHandler(
				di.Get(ctx, "registry").(registry.Registry),
				di.Get(ctx, "integrationEventHandlers").(ddd.EventHandler[ddd.Event]),
			),
			di.Get(ctx, "inboxMiddleware").(am.RawMessageHandlerMiddleware),
		)

		return evtHandlers.HandleMessage(ctx, msg)
	})

	subscriber := container.Get("stream").(am.RawMessageStream)

	if _, err := subscriber.Subscribe(customerv1.CustomerAggregateChannel, evtMsgHandler, am.MessageFilter{
		customerv1.CustomerRegisteredEvent,
	}, am.GroupName("search-customers")); err != nil {
		return err
	}

	if _, err := subscriber.Subscribe(orderv1.OrderAggregateChannel, evtMsgHandler, am.MessageFilter{
		orderv1.OrderCreatedEvent,
		orderv1.OrderReadiedEvent,
		orderv1.OrderCanceledEvent,
		orderv1.OrderCompletedEvent,
	}, am.GroupName("search-orders")); err != nil {
		return err
	}

	if _, err := subscriber.Subscribe(storev1.ProductAggregateChannel, evtMsgHandler, am.MessageFilter{
		storev1.ProductAddedEvent,
		storev1.ProductRebrandedEvent,
		storev1.ProductRemovedEvent,
	}, am.GroupName("search-products")); err != nil {
		return err
	}

	if _, err := subscriber.Subscribe(storev1.StoreAggregateChannel, evtMsgHandler, am.MessageFilter{
		storev1.StoreCreatedEvent,
		storev1.StoreRebrandedEvent,
	}, am.GroupName("search-stores")); err != nil {
		return err
	}

	return nil
}
