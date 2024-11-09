package logging

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/logger"
)

type CommandHandler[T ddd.Command] struct {
	ddd.CommandHandler[T]
	label  string
	logger logger.Logger
}

var _ ddd.CommandHandler[ddd.Command] = (*CommandHandler[ddd.Command])(nil)

func NewLogCommandHandlerAccess[T ddd.Command](handler ddd.CommandHandler[T], label string, logger logger.Logger) *CommandHandler[T] {
	return &CommandHandler[T]{
		CommandHandler: handler,
		label:          label,
		logger:         logger,
	}
}

// HandleEvent implements ddd.EventHandler.
func (h *CommandHandler[T]) HandleEvent(ctx context.Context, command T) (replay ddd.Reply, err error) {
	h.logger.Infof("--> Payment.%s.On(%s)", h.label, command.CommandName())
	defer func() { h.logger.WithError(err).Infof("<-- Payment.%s.On(%s)", h.label, command.CommandName()) }()
	return h.CommandHandler.HandleCommand(ctx, command)
}
