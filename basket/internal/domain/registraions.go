package domain

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/entity"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/event"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/es"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry/serdes"
)

func Registrations(reg registry.Registry) error {
	serde := serdes.NewJSONSerde(reg)
	// Basket
	if err := serde.Register(aggregate.Basket{}, func(v interface{}) error {
		basket := v.(*aggregate.Basket)
		basket.Aggregate = es.NewAggregate("", aggregate.BasketAggregate)
		basket.Items = make(map[string]*entity.Item)
		return nil
	}); err != nil {
		return err
	}

	// basket events
	if err := serde.Register(event.BasketStarted{}); err != nil {
		return err
	}
	if err := serde.Register(event.BasketCanceled{}); err != nil {
		return err
	}
	if err := serde.Register(event.BasketCheckedOut{}); err != nil {
		return err
	}
	if err := serde.Register(event.BasketItemAdded{}); err != nil {
		return err
	}
	if err := serde.Register(event.BasketItemRemoved{}); err != nil {
		return err
	}

	// basket snapshots
	if err := serde.RegisterKey(aggregate.BasketV1{}.SnapshotName(), aggregate.BasketV1{}); err != nil {
		return err
	}

	return nil
}
