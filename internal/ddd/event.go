package ddd

import (
	"time"

	"github.com/google/uuid"
)

type (
	EventPayload interface{}

	Event interface {
		IDer
		EventName() string
		Payload() EventPayload
		Metadata() Metadata
		OccurredAt() time.Time
	}

	eventBase struct {
		EntityBase
		payload    EventPayload
		metadata   Metadata
		occurredAt time.Time
	}
)

var _ Event = (*eventBase)(nil)

func NewEventBase(name string, payload EventPayload, options ...EventOption) eventBase {
	return newEventBase(name, payload, options...)
}

func newEventBase(name string, payload EventPayload, options ...EventOption) eventBase {
	evt := eventBase{
		EntityBase: NewEntityBase(uuid.New().String(), name),
		payload:    payload,
		metadata:   make(Metadata),
		occurredAt: time.Now(),
	}

	for _, option := range options {
		option.configureEvent(&evt)
	}

	return evt
}

func (e eventBase) EventName() string     { return e.name }
func (e eventBase) Payload() EventPayload { return e.payload }
func (e eventBase) Metadata() Metadata    { return e.metadata }
func (e eventBase) OccurredAt() time.Time { return e.occurredAt }
