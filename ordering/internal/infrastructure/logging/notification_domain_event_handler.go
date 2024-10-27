package logging

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/logger"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application/port/in/event"
)

type NotificationDomainEventHandler struct {
	event.DomainEventHandlers
	logger logger.Logger
}

var _ event.DomainEventHandlers = (*NotificationDomainEventHandler)(nil)

func NewLogNotificationDomainEventHandlerAccess(handlers event.DomainEventHandlers, logger logger.Logger) *NotificationDomainEventHandler {
	return &NotificationDomainEventHandler{
		DomainEventHandlers: handlers,
		logger:              logger,
	}
}

// OnOrderCreated implements event.DomainEventHandlers.
func (n *NotificationDomainEventHandler) OnOrderCreated(ctx context.Context, event ddd.DomainEvent) (err error) {
	n.logger.Info("--> Ordering.OnOrderCreated")
	defer func() { n.logger.WithError(err).Info("<-- Ordering.OnOrderCreated") }()
	return n.DomainEventHandlers.OnOrderCreated(ctx, event)
}

// OnOrderCanceled implements event.DomainEventHandlers.
func (n *NotificationDomainEventHandler) OnOrderCanceled(ctx context.Context, event ddd.DomainEvent) (err error) {
	n.logger.Info("--> Ordering.OnOrderCanceled")
	defer func() { n.logger.WithError(err).Info("<-- Ordering.OnOrderCanceled") }()
	return n.DomainEventHandlers.OnOrderCanceled(ctx, event)
}

// OnOrderReadied implements event.DomainEventHandlers.
func (n *NotificationDomainEventHandler) OnOrderReadied(ctx context.Context, event ddd.DomainEvent) (err error) {
	n.logger.Info("--> Ordering.OnOrderReadied")
	defer func() { n.logger.WithError(err).Info("<-- Ordering.OnOrderReadied") }()
	return n.DomainEventHandlers.OnOrderReadied(ctx, event)
}
