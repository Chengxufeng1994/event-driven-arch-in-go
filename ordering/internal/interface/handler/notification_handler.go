package handler

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	domainevent "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/event"
)

func RegisterNotificationHandler(handlers ddd.EventHandler[ddd.AggregateEvent], subscriber ddd.EventSubscriber[ddd.AggregateEvent]) {
	subscriber.Subscribe(domainevent.OrderCreatedEvent, handlers)
	subscriber.Subscribe(domainevent.OrderCanceledEvent, handlers)
	subscriber.Subscribe(domainevent.OrderReadiedEvent, handlers)
}
