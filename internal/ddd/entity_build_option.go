package ddd

import (
	"fmt"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"
)

type IDSetter interface {
	setID(string)
}

type NameSetter interface {
	setName(string)
}

func WithID(id string) registry.BuildOption {
	return func(v interface{}) error {
		if e, ok := v.(IDSetter); ok {
			e.setID(id)
			return nil
		}

		return fmt.Errorf("%T does not have the method setID(string)", v)
	}
}

func WithName(name string) registry.BuildOption {
	return func(v interface{}) error {
		if e, ok := v.(NameSetter); ok {
			e.setName(name)
			return nil
		}
		return fmt.Errorf("%T does not have the method setName(string)", v)
	}
}
