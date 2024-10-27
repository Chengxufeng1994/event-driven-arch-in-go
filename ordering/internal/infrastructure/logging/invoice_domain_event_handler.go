package logging

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/logger"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application/port/in/event"
)

type InvoiceDomainEventHandler struct {
	event.DomainEventHandlers
	logger logger.Logger
}

var _ event.DomainEventHandlers = (*InvoiceDomainEventHandler)(nil)

func NewLogInvoiceDomainEventHandlerAccess(handlers event.DomainEventHandlers, logger logger.Logger) *InvoiceDomainEventHandler {
	return &InvoiceDomainEventHandler{
		DomainEventHandlers: handlers,
		logger:              logger,
	}
}

// OnOrderReadied implements event.DomainEventHandlers.
func (i *InvoiceDomainEventHandler) OnOrderReadied(ctx context.Context, event ddd.DomainEvent) (err error) {
	i.logger.Info("--> Ordering.OnOrderReadied")
	defer func() { i.logger.WithError(err).Info("<-- Ordering.OnOrderReadied") }()
	return i.DomainEventHandlers.OnOrderReadied(ctx, event)
}
