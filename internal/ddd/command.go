package ddd

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type (
	CommandHandler[T Command] interface {
		HandleCommand(ctx context.Context, cmd T) (Reply, error)
	}

	CommandHandlerFunc[T Command] func(ctx context.Context, cmd T) (Reply, error)

	CommandOption interface {
		configureCommand(*command)
	}

	CommandPayload any

	Command interface {
		IDer
		CommandName() string
		Payload() CommandPayload
		Metadata() Metadata
		OccurredAt() time.Time
	}

	command struct {
		EntityBase
		payload    CommandPayload
		metadata   Metadata
		occurredAt time.Time
	}
)

var _ Command = (*command)(nil)

func NewCommand(name string, payload CommandPayload, options ...CommandOption) Command {
	return newCommand(name, payload, options...)
}

func newCommand(name string, payload CommandPayload, options ...CommandOption) command {
	evt := command{
		EntityBase: NewEntityBase(uuid.New().String(), name),
		payload:    payload,
		metadata:   make(Metadata),
		occurredAt: time.Now(),
	}

	for _, option := range options {
		option.configureCommand(&evt)
	}

	return evt
}

func (e command) CommandName() string     { return e.name }
func (e command) Payload() CommandPayload { return e.payload }
func (e command) Metadata() Metadata      { return e.metadata }
func (e command) OccurredAt() time.Time   { return e.occurredAt }

var _ CommandHandler[Command] = (*CommandHandlerFunc[Command])(nil)

func (f CommandHandlerFunc[T]) HandleCommand(ctx context.Context, cmd T) (Reply, error) {
	return f(ctx, cmd)
}