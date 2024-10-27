package ddd

import "context"

type EventHandler func(ctx context.Context, event DomainEvent) error

type DomainEvent interface {
	EventName() string
}
