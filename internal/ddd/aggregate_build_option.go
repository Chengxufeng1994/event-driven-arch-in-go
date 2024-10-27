package ddd

import (
	"fmt"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"
)

type EventSetter interface {
	setEvents([]Event)
}

func WithEvents(events ...Event) registry.BuildOption {
	return func(v interface{}) error {
		if agg, ok := v.(EventSetter); ok {
			agg.setEvents(events)
			return nil
		}
		return fmt.Errorf("%T does not have the method setEvents([]ddd.Event)", v)
	}
}
