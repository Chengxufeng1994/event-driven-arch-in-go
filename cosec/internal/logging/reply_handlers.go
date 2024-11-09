package logging

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/logger"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/sec"
)

type sagaReplyHandlers[T any] struct {
	sec.Orchestrator[T]
	label  string
	logger logger.Logger
}

var _ sec.Orchestrator[any] = (*sagaReplyHandlers[any])(nil)

func NewLogReplyHandlerAccess[T any](orc sec.Orchestrator[T], label string, logger logger.Logger) sec.Orchestrator[T] {
	return sagaReplyHandlers[T]{
		Orchestrator: orc,
		label:        label,
		logger:       logger,
	}
}

func (h sagaReplyHandlers[T]) HandleReply(ctx context.Context, reply ddd.Reply) (err error) {
	h.logger.Infof("--> COSEC.%s.On(%s)", h.label, reply.ReplyName())
	defer func() { h.logger.WithError(err).Infof("<-- COSEC.%s.On(%s)", h.label, reply.ReplyName()) }()
	return h.Orchestrator.HandleReply(ctx, reply)
}
