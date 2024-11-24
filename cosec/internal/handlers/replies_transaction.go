package handlers

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/cosec/internal/constants"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/di"
	"gorm.io/gorm"
)

func RegisterReplyHandlersTx(container di.Container) error {
	replyMsgHandler := am.MessageHandlerFunc(func(ctx context.Context, msg am.IncomingMessage) (err error) {
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
		}(di.Get(ctx, constants.DatabaseTransactionKey).(*gorm.DB))

		return di.Get(ctx, constants.ReplyHandlersKey).(am.MessageHandler).HandleMessage(ctx, msg)
	})

	subscriber := container.Get(constants.MessageSubscriberKey).(am.MessageSubscriber)

	return RegisterReplyHandlers(subscriber, replyMsgHandler)
}
