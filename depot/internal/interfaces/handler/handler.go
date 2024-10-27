package handler

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/application/port/in/event"
	domainevent "github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/event"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
)

func RegisterDomainEventHandlers(_ context.Context, domainEventHandlers event.DomainEventHandlers, domainSubscriber ddd.EventSubscriber) {
	domainSubscriber.Subscribe(domainevent.ShoppingListCompleted{}, domainEventHandlers.OnShoppingListCompleted)
}
