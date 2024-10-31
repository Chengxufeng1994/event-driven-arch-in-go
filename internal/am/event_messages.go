package am

import (
	"context"
	"time"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type (
	EventMessage interface {
		Message
		ddd.Event
	}

	EventPublisher  = MessagePublisher[ddd.Event]
	EventSubscriber = MessageSubscriber[EventMessage]
	EventStream     = MessageStream[ddd.Event, EventMessage]

	// EventStreamBase
	eventStream struct {
		reg    registry.Registry
		stream MessageStream[RawMessage, RawMessage]
	}

	// EventMessageBase
	eventMessage struct {
		id         string
		name       string
		payload    ddd.EventPayload
		metadata   ddd.Metadata
		occurredAt time.Time
		msg        RawMessage
	}
)

var _ EventStream = (*eventStream)(nil)
var _ EventMessage = (*eventMessage)(nil)

func NewEventStream(reg registry.Registry, stream MessageStream[RawMessage, RawMessage]) EventStream {
	return &eventStream{
		reg:    reg,
		stream: stream,
	}
}

// Publish implements MessageStream.
func (es *eventStream) Publish(ctx context.Context, topicName string, event ddd.Event) error {
	metadata, err := structpb.NewStruct(event.Metadata())
	if err != nil {
		return err
	}

	payload, err := es.reg.Serialize(event.EventName(), event.Payload())
	if err != nil {
		return err
	}

	data, err := proto.Marshal(&EventMessageData{
		Payload:    payload,
		OccurredAt: timestamppb.New(event.OccurredAt()),
		Metadata:   metadata,
	})
	if err != nil {
		return err
	}

	return es.stream.Publish(
		ctx,
		topicName,
		rawMessage{
			id:   event.ID(),
			name: event.EventName(),
			data: data,
		},
	)
}

// Subscribe implements MessageStream.
func (es *eventStream) Subscribe(topicName string, handler MessageHandlerFunc[EventMessage], options ...SubscriberOption) error {
	cfg := NewSubscriberConfig(options)

	var filters map[string]struct{}
	if len(cfg.MessageFilters()) > 0 {
		filters = make(map[string]struct{})
		for _, filter := range cfg.MessageFilters() {
			filters[filter] = struct{}{}
		}
	}

	fn := MessageHandlerFunc[RawMessage](func(ctx context.Context, msg RawMessage) error {
		var eventData EventMessageData

		if filters != nil {
			if _, exists := filters[msg.MessageName()]; !exists {
				return nil
			}
		}

		err := proto.Unmarshal(msg.Data(), &eventData)
		if err != nil {
			return err
		}

		eventName := msg.MessageName()
		payload, err := es.reg.Deserialize(eventName, eventData.GetPayload())
		if err != nil {
			return err
		}

		return handler(ctx, &eventMessage{
			id:         msg.ID(),
			name:       eventName,
			payload:    payload,
			metadata:   eventData.GetMetadata().AsMap(),
			occurredAt: eventData.GetOccurredAt().AsTime(),
			msg:        msg,
		})
	})

	return es.stream.Subscribe(topicName, fn, options...)
}

func (e eventMessage) ID() string                { return e.id }
func (e eventMessage) EventName() string         { return e.name }
func (e eventMessage) Payload() ddd.EventPayload { return e.payload }
func (e eventMessage) Metadata() ddd.Metadata    { return e.metadata }
func (e eventMessage) OccurredAt() time.Time     { return e.occurredAt }
func (e eventMessage) MessageName() string       { return e.msg.MessageName() }
func (e eventMessage) Ack() error                { return e.msg.Ack() }
func (e eventMessage) NAck() error               { return e.msg.NAck() }
func (e eventMessage) Extend() error             { return e.msg.Extend() }
func (e eventMessage) Kill() error               { return e.msg.Kill() }
