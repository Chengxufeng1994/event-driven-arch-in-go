package handler

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/di"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"

	"gorm.io/gorm"
)

func RegisterCommandHandlersTx(container di.Container) error {
	cmdMsgHandler := am.RawMessageHandlerFunc(func(ctx context.Context, msg am.IncomingRawMessage) (err error) {
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

		cmdMsgHandlers := am.RawMessageHandlerWithMiddleware(
			am.NewCommandMessageHandler(
				container.Get("registry").(registry.Registry),
				container.Get("replyStream").(am.ReplyStream),
				container.Get("commandHandlers").(ddd.CommandHandler[ddd.Command]),
			).(am.RawMessageHandler),
			container.Get("inboxMiddleware").(am.RawMessageHandlerMiddleware),
		)

		return cmdMsgHandlers.HandleMessage(ctx, msg)
	})

	subscriber := container.Get("stream").(am.RawMessageStream)

	return RegisterCommandHandlers(subscriber, cmdMsgHandler)
}
