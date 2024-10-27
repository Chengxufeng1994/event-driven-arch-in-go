package logging

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/port/in/event"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/logger"
)

type DomainEventHandler struct {
	event.DomainEventHandlers
	logger logger.Logger
}

var _ event.DomainEventHandlers = (*DomainEventHandler)(nil)

func NewLogDomainEventHandlerAccess(domainEventHandler event.DomainEventHandlers, logger logger.Logger) *DomainEventHandler {
	return &DomainEventHandler{domainEventHandler, logger}
}

func (d *DomainEventHandler) OnBasketCanceled(ctx context.Context, event ddd.DomainEvent) (err error) {
	d.logger.Info("--> OnBasketCanceled")
	defer func() { d.logger.WithError(err).Info("<-- OnBasketCanceled") }()
	return d.DomainEventHandlers.OnBasketCanceled(ctx, event)
}

// OnBasketCheckedOut implements event.DomainEventHandlers.
func (d *DomainEventHandler) OnBasketCheckedOut(ctx context.Context, event ddd.DomainEvent) (err error) {
	d.logger.Info("--> OnBasketCheckedOut")
	defer func() { d.logger.WithError(err).Info("<-- OnBasketCheckedOut") }()
	return d.DomainEventHandlers.OnBasketCheckedOut(ctx, event)
}

// OnBasketItemAdded implements event.DomainEventHandlers.
func (d *DomainEventHandler) OnBasketItemAdded(ctx context.Context, event ddd.DomainEvent) (err error) {
	d.logger.Info("--> OnBasketItemAdded")
	defer func() { d.logger.WithError(err).Info("<-- OnBasketItemAdded") }()
	return d.DomainEventHandlers.OnBasketItemAdded(ctx, event)
}

// OnBasketItemRemoved implements event.DomainEventHandlers.
func (d *DomainEventHandler) OnBasketItemRemoved(ctx context.Context, event ddd.DomainEvent) (err error) {
	d.logger.Info("--> OnBasketItemRemoved")
	defer func() { d.logger.WithError(err).Info("<-- OnBasketItemRemoved") }()
	return d.DomainEventHandlers.OnBasketItemRemoved(ctx, event)
}

// OnBasketStarted implements event.DomainEventHandlers.
func (d *DomainEventHandler) OnBasketStarted(ctx context.Context, event ddd.DomainEvent) (err error) {
	d.logger.Info("--> OnBasketStarted")
	defer func() { d.logger.WithError(err).Info("<-- OnBasketStarted") }()
	return d.DomainEventHandlers.OnBasketStarted(ctx, event)
}
