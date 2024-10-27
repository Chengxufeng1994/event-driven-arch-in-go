package ddd

import (
	"context"
	"sync"
)

type EventSubscriber interface {
	Subscribe(event DomainEvent, eventHandler EventHandler)
}

type EventPublisher interface {
	Publish(ctx context.Context, events ...DomainEvent) error
}

type EventDispatcherIntf interface {
	EventSubscriber
	EventPublisher
}

// another way
// var _ interface {
// 	EventSubscriber
// 	EventPublisher
// } = (*EventDispatcher)(nil)

type EventDispatcher struct {
	handlers map[string][]EventHandler
	mu       sync.Mutex
}

var _ EventDispatcherIntf = (*EventDispatcher)(nil)

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandler),
		mu:       sync.Mutex{},
	}
}

// Subscribe implements EventDispatcherIntf.
func (e *EventDispatcher) Subscribe(event DomainEvent, eventHandler EventHandler) {
	e.mu.Lock()
	defer e.mu.Unlock()
	if _, ok := e.handlers[event.EventName()]; !ok {
		e.handlers[event.EventName()] = make([]EventHandler, 0)
	}

	e.handlers[event.EventName()] = append(e.handlers[event.EventName()], eventHandler)
}

// Publish implements EventDispatcherIntf.
func (e *EventDispatcher) Publish(ctx context.Context, events ...DomainEvent) error {
	for _, event := range events {
		for _, handler := range e.handlers[event.EventName()] {
			err := handler(ctx, event)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
