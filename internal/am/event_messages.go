package am

import (
	"context"
	"time"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"
)

type (
	EventMessage interface {
		MessageBase
		ddd.Event
	}

	IncomingEventMessage interface {
		IncomingMessageBase
		ddd.Event
	}

	EventPublisher interface {
		Publish(ctx context.Context, topicName string, event ddd.Event) error
	}

	eventPublisher struct {
		reg       registry.Registry
		publisher MessagePublisher
	}

	eventMessage struct {
		id         string
		name       string
		payload    ddd.EventPayload
		occurredAt time.Time
		msg        IncomingMessageBase
	}

	eventMsgHandler struct {
		reg     registry.Registry
		handler ddd.EventHandler[ddd.Event]
	}
)

var _ EventMessage = (*eventMessage)(nil)
var _ EventPublisher = (*eventPublisher)(nil)
var _ MessageHandler = (*eventMsgHandler)(nil)

func NewEventPublisher(reg registry.Registry, msgPublisher MessagePublisher, mws ...MessagePublisherMiddleware) EventPublisher {
	return eventPublisher{
		reg:       reg,
		publisher: MessagePublisherWithMiddleware(msgPublisher, mws...),
	}
}

func (e eventPublisher) Publish(ctx context.Context, topicName string, event ddd.Event) error {
	payload, err := e.reg.Serialize(event.EventName(), event.Payload())
	if err != nil {
		return err
	}

	data, err := proto.Marshal(&EventMessageData{
		Payload:    payload,
		OccurredAt: timestamppb.New(event.OccurredAt()),
	})
	if err != nil {
		return err
	}

	return e.publisher.Publish(ctx, topicName, message{
		id:       event.ID(),
		name:     event.EventName(),
		subject:  topicName,
		data:     data,
		metadata: event.Metadata(),
		sentAt:   time.Now(),
	})
}

func (e eventMessage) ID() string                { return e.id }
func (e eventMessage) EventName() string         { return e.name }
func (e eventMessage) Payload() ddd.EventPayload { return e.payload }
func (e eventMessage) Metadata() ddd.Metadata    { return e.msg.Metadata() }
func (e eventMessage) OccurredAt() time.Time     { return e.occurredAt }
func (e eventMessage) Subject() string           { return e.msg.Subject() }
func (e eventMessage) MessageName() string       { return e.msg.MessageName() }
func (e eventMessage) SentAt() time.Time         { return e.msg.SentAt() }
func (e eventMessage) ReceivedAt() time.Time     { return e.msg.ReceivedAt() }
func (e eventMessage) Ack() error                { return e.msg.Ack() }
func (e eventMessage) NAck() error               { return e.msg.NAck() }
func (e eventMessage) Extend() error             { return e.msg.Extend() }
func (e eventMessage) Kill() error               { return e.msg.Kill() }

func NewEventHandler(reg registry.Registry, handler ddd.EventHandler[ddd.Event], mws ...MessageHandlerMiddleware) MessageHandler {
	return MessageHandlerWithMiddleware(eventMsgHandler{
		reg:     reg,
		handler: handler,
	}, mws...)
}

func (h eventMsgHandler) HandleMessage(ctx context.Context, msg IncomingMessage) error {
	var eventData EventMessageData

	err := proto.Unmarshal(msg.Data(), &eventData)
	if err != nil {
		return err
	}

	eventName := msg.MessageName()

	payload, err := h.reg.Deserialize(eventName, eventData.GetPayload())
	if err != nil {
		return err
	}

	eventMsg := eventMessage{
		id:         msg.ID(),
		name:       eventName,
		payload:    payload,
		occurredAt: eventData.GetOccurredAt().AsTime(),
		msg:        msg,
	}

	return h.handler.HandleEvent(ctx, eventMsg)
}
