package logging

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/application/port/in/event"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/logger"
)

type DomainEventHandler struct {
	event.DomainEventHandlers
	logger logger.Logger
}

var _ event.DomainEventHandlers = (*DomainEventHandler)(nil)

func NewLogDomainEventHandlerAccess(handlers event.DomainEventHandlers, logger logger.Logger) *DomainEventHandler {
	return &DomainEventHandler{
		DomainEventHandlers: handlers,
		logger:              logger,
	}
}

// OnShoppingListAssigned implements event.DomainEventHandlers.
func (d *DomainEventHandler) OnShoppingListAssigned(ctx context.Context, event ddd.DomainEvent) (err error) {
	d.logger.Info("--> OnShoppingListAssigned")
	defer func() { d.logger.WithError(err).Info("<-- OnShoppingListAssigned") }()
	return d.DomainEventHandlers.OnShoppingListAssigned(ctx, event)
}

// OnShoppingListCanceled implements event.DomainEventHandlers.
func (d *DomainEventHandler) OnShoppingListCanceled(ctx context.Context, event ddd.DomainEvent) (err error) {
	d.logger.Info("--> OnShoppingListCanceled")
	defer func() { d.logger.WithError(err).Info("<-- OnShoppingListCanceled") }()
	return d.DomainEventHandlers.OnShoppingListCanceled(ctx, event)
}

// OnShoppingListCompleted implements event.DomainEventHandlers.
func (d *DomainEventHandler) OnShoppingListCompleted(ctx context.Context, event ddd.DomainEvent) (err error) {
	d.logger.Info("--> OnShoppingListCompleted")
	defer func() { d.logger.WithError(err).Info("<-- OnShoppingListCompleted") }()
	return d.DomainEventHandlers.OnShoppingListCompleted(ctx, event)
}

// OnShoppingListCreated implements event.DomainEventHandlers.
func (d *DomainEventHandler) OnShoppingListCreated(ctx context.Context, event ddd.DomainEvent) (err error) {
	d.logger.Info("--> OnShoppingListCreated")
	defer func() { d.logger.WithError(err).Info("<-- OnShoppingListCreated") }()
	return d.DomainEventHandlers.OnShoppingListCreated(ctx, event)
}
