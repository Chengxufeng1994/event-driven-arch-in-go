package logging

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/logger"
)

type EventHandler[T ddd.Event] struct {
	ddd.EventHandler[T]
	label  string
	logger logger.Logger
}

var _ ddd.EventHandler[ddd.Event] = (*EventHandler[ddd.Event])(nil)

func NewLogEventHandlerAccess[T ddd.Event](handler ddd.EventHandler[T], label string, logger logger.Logger) *EventHandler[T] {
	return &EventHandler[T]{
		EventHandler: handler,
		label:        label,
		logger:       logger,
	}
}

func (h *EventHandler[T]) HandleEvent(ctx context.Context, event T) (err error) {
	h.logger.Infof("--> Customers.%s.On(%s)", h.label, event.EventName())
	defer func() { h.logger.WithError(err).Infof("<-- Customers.%s.On(%s)", h.label, event.EventName()) }()
	return h.EventHandler.HandleEvent(ctx, event)
}
