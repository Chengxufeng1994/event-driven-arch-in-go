package handlers

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/cosec/internal/models"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/sec"
)

func RegisterReplyHandlers(subscriber am.ReplySubscriber, orchestrator sec.Orchestrator[*models.CreateOrderData]) error {
	replyMsgHandler := am.MessageHandlerFunc[am.IncomingReplyMessage](func(ctx context.Context, replyMsg am.IncomingReplyMessage) error {
		return orchestrator.HandleReply(ctx, replyMsg)
	})
	_, err := subscriber.Subscribe(orchestrator.ReplyTopic(), replyMsgHandler, am.GroupName("cosec-replies"))
	return err
}
