package logging

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/logger"
)

type CommandHandlers[T ddd.Command] struct {
	ddd.CommandHandler[T]
	label  string
	logger logger.Logger
}

var _ ddd.CommandHandler[ddd.Command] = (*CommandHandlers[ddd.Command])(nil)

func NewLogCommandHandlerAccess[T ddd.Command](handler ddd.CommandHandler[T], label string, logger logger.Logger) ddd.CommandHandler[T] {
	return &CommandHandlers[T]{
		CommandHandler: handler,
		label:          label,
		logger:         logger,
	}
}

func (h CommandHandlers[T]) HandleEvent(ctx context.Context, event T) (reply ddd.Reply, err error) {
	h.logger.Infof("--> Customers.%s.On(%s)", h.label, event.CommandName())
	defer func() { h.logger.WithError(err).Infof("<-- Customers.%s.On(%s)", h.label, event.CommandName()) }()
	return h.CommandHandler.HandleCommand(ctx, event)
}
