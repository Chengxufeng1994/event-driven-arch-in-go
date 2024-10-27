package logging

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/logger"
)

type InvoiceEventHandler[T ddd.Event] struct {
	ddd.EventHandler[T]
	label  string
	logger logger.Logger
}

var _ ddd.EventHandler[ddd.Event] = (*InvoiceEventHandler[ddd.Event])(nil)

func NewLogInvoiceEventHandlerAccess[T ddd.Event](handlers ddd.EventHandler[T], label string, logger logger.Logger) *InvoiceEventHandler[T] {
	return &InvoiceEventHandler[T]{
		EventHandler: handlers,
		label:        label,
		logger:       logger,
	}
}

func (h *InvoiceEventHandler[T]) HandleEvent(ctx context.Context, event T) (err error) {
	h.logger.Infof("--> Ordering.%s.On(%s)", h.label, event.EventName())
	defer func() { h.logger.WithError(err).Infof("<-- Ordering.%s.On(%s)", h.label, event.EventName()) }()
	return h.EventHandler.HandleEvent(ctx, event)
}
