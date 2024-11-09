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

func NewLogEventHandlerAccess[T ddd.Event](domainEventHandler ddd.EventHandler[T], label string, logger logger.Logger) *EventHandler[T] {
	return &EventHandler[T]{
		EventHandler: domainEventHandler,
		label:        label,
		logger:       logger,
	}
}

// HandleEvent implements ddd.EventHandler.
func (h *EventHandler[T]) HandleEvent(ctx context.Context, event T) (err error) {
	h.logger.Infof("--> Ordering.%s.On(%s)", h.label, event.EventName())
	defer func() { h.logger.WithError(err).Infof("<-- Ordering.%s.On(%s)", h.label, event.EventName()) }()
	return h.EventHandler.HandleEvent(ctx, event)
}
