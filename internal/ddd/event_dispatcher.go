package ddd

import (
	"context"
	"sync"
)

type (
	EventHandler[T Event] interface {
		HandleEvent(ctx context.Context, event T) error
	}

	EventHandlerFunc[T Event] func(ctx context.Context, event T) error

	EventSubscriber[T Event] interface {
		Subscribe(event string, eventHandler EventHandler[T])
	}

	EventPublisher[T Event] interface {
		Publish(ctx context.Context, events ...T) error
	}
	EventDispatcher[T Event] interface {
		EventSubscriber[T]
		EventPublisher[T]
	}
)

type EventDispatcherBase[T Event] struct {
	handlers map[string][]EventHandler[T]
	mu       sync.Mutex
}

var _ EventDispatcher[Event] = (*EventDispatcherBase[Event])(nil)

func NewEventDispatcher[T Event]() *EventDispatcherBase[T] {
	return &EventDispatcherBase[T]{
		handlers: make(map[string][]EventHandler[T]),
		mu:       sync.Mutex{},
	}
}

// Subscribe implements EventDispatcherIntf.
func (e *EventDispatcherBase[T]) Subscribe(name string, eventHandler EventHandler[T]) {
	e.mu.Lock()
	defer e.mu.Unlock()
	if _, ok := e.handlers[name]; !ok {
		e.handlers[name] = make([]EventHandler[T], 0)
	}

	e.handlers[name] = append(e.handlers[name], eventHandler)
}

// Publish implements EventDispatcherIntf.
func (e *EventDispatcherBase[T]) Publish(ctx context.Context, events ...T) error {
	for _, event := range events {
		for _, handler := range e.handlers[event.EventName()] {
			err := handler.HandleEvent(ctx, event)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (f EventHandlerFunc[T]) HandleEvent(ctx context.Context, event T) error {
	return f(ctx, event)
}
