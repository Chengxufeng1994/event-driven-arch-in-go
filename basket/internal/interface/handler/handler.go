package handler

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/port/in/event"
	domainevent "github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/event"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
)

func RegisterDomainEventHandlers(_ context.Context, handlers event.DomainEventHandlers, domainSubscriber ddd.EventSubscriber) {
	domainSubscriber.Subscribe(domainevent.BasketCheckedOut{}, handlers.OnBasketCheckedOut)
}
