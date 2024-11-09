package handlers

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/cosec/internal"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/cosec/internal/models"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/di"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/sec"
	"gorm.io/gorm"
)

func RegisterReplyHandlersTx(container di.Container) error {
	replyMsgHandler := am.RawMessageHandlerFunc(func(ctx context.Context, msg am.IncomingRawMessage) (err error) {
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

		replyHandlers := am.RawMessageHandlerWithMiddleware(
			am.NewReplyMessageHandler(
				di.Get(ctx, "registry").(registry.Registry),
				di.Get(ctx, "orchestrator").(sec.Orchestrator[*models.CreateOrderData]),
			),
			di.Get(ctx, "inboxMiddleware").(am.RawMessageHandlerMiddleware),
		)

		return replyHandlers.HandleMessage(ctx, msg)
	})

	subscriber := container.Get("stream").(am.RawMessageStream)

	return subscriber.Subscribe(internal.CreateOrderReplyChannel, replyMsgHandler, am.GroupName("cosec-replies"))
}
