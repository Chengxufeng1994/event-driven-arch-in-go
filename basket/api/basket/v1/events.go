package basketv1

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry/serdes"
)

const (
	BasketAggregateChannel = "mallbots.baskets.events.Basket"

	BasketStartedEvent    = "basketsapi.BasketStarted"
	BasketCanceledEvent   = "basketsapi.BasketCanceled"
	BasketCheckedOutEvent = "basketsapi.BasketCheckedOut"
)

func Registrations(reg registry.Registry) error {
	serde := serdes.NewProtoSerde(reg)

	// Basket events
	if err := serde.Register(&BasketStarted{}); err != nil {
		return err
	}
	if err := serde.Register(&BasketCanceled{}); err != nil {
		return err
	}
	if err := serde.Register(&BasketCheckedOut{}); err != nil {
		return err
	}

	return nil
}

func (*BasketStarted) Key() string    { return BasketStartedEvent }
func (*BasketCanceled) Key() string   { return BasketCanceledEvent }
func (*BasketCheckedOut) Key() string { return BasketCheckedOutEvent }
