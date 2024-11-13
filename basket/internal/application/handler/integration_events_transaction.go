package handler

import (
	"context"

	"gorm.io/gorm"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/di"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
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

	_, err := subscriber.Subscribe(storev1.StoreAggregateChannel, evtMsgHandler, am.MessageFilter{
		storev1.StoreCreatedEvent,
		storev1.StoreRebrandedEvent,
	}, am.GroupName("baskets-stores"))
	if err != nil {
		return err
	}

	_, err = subscriber.Subscribe(storev1.ProductAggregateChannel, evtMsgHandler, am.MessageFilter{
		storev1.ProductAddedEvent,
		storev1.ProductRebrandedEvent,
		storev1.ProductPriceIncreasedEvent,
		storev1.ProductPriceDecreasedEvent,
		storev1.ProductRemovedEvent,
	}, am.GroupName("baskets-products"))

	return err
}
