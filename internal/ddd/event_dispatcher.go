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
		Subscribe(eventHandler EventHandler[T], events ...string)
	}

	EventPublisher[T Event] interface {
		Publish(ctx context.Context, events ...T) error
	}
	EventDispatcher[T Event] interface {
		EventSubscriber[T]
		EventPublisher[T]
	}

	eventHandler[T Event] struct {
		h       EventHandler[T]
		filters map[string]struct{}
	}

	EventDispatcherBase[T Event] struct {
		handlers []eventHandler[T]
		mu       sync.Mutex
	}
)

var _ EventDispatcher[Event] = (*EventDispatcherBase[Event])(nil)

func NewEventDispatcher[T Event]() *EventDispatcherBase[T] {
	return &EventDispatcherBase[T]{
		handlers: make([]eventHandler[T], 0),
		mu:       sync.Mutex{},
	}
}

// Subscribe implements EventDispatcherIntf.
func (h *EventDispatcherBase[T]) Subscribe(handler EventHandler[T], events ...string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	var filters map[string]struct{}
	if len(events) > 0 {
		filters = make(map[string]struct{})
		for _, event := range events {
			filters[event] = struct{}{}
		}
	}

	h.handlers = append(h.handlers, eventHandler[T]{
		h:       handler,
		filters: filters,
	})
}

// Publish implements EventDispatcherIntf.
func (h *EventDispatcherBase[T]) Publish(ctx context.Context, events ...T) error {
	for _, event := range events {
		for _, handler := range h.handlers {
			if handler.filters != nil {
				if _, exists := handler.filters[event.EventName()]; !exists {
					continue
				}
			}
			err := handler.h.HandleEvent(ctx, event)
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
