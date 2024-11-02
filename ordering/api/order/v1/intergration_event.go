package orderv1

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry/serdes"
)

const (
	OrderAggregateChannel = "mallbots.ordering.events.Order"

	OrderCreatedEvent   = "ordersapi.OrderCreated"
	OrderReadiedEvent   = "ordersapi.OrderReadied"
	OrderCanceledEvent  = "ordersapi.OrderCanceled"
	OrderCompletedEvent = "ordersapi.OrderCompleted"
)

func Registrations(reg registry.Registry) error {
	serde := serdes.NewProtoSerde(reg)

	// Order events
	if err := serde.Register(&OrderCreated{}); err != nil {
		return err
	}
	if err := serde.Register(&OrderReadied{}); err != nil {
		return err
	}
	if err := serde.Register(&OrderCanceled{}); err != nil {
		return err
	}
	if err := serde.Register(&OrderCompleted{}); err != nil {
		return err
	}

	return nil
}

func (*OrderCreated) Key() string   { return OrderCreatedEvent }
func (*OrderReadied) Key() string   { return OrderReadiedEvent }
func (*OrderCanceled) Key() string  { return OrderCanceledEvent }
func (*OrderCompleted) Key() string { return OrderCompletedEvent }
