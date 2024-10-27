package handler

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application/port/in/event"
	domainevent "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/event"
)

func RegisterInvoiceDomainEventHandlers(handlers event.DomainEventHandlers, subscriber ddd.EventSubscriber) {
	subscriber.Subscribe(domainevent.OrderReadied{}, handlers.OnOrderReadied)
}
