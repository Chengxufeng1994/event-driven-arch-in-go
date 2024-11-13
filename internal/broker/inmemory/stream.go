package inmemory

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
)

type stream struct {
	subscriptions map[string][]am.RawMessageHandler
}

var _ am.RawMessageStream = (*stream)(nil)

func NewStream() stream {
	return stream{
		subscriptions: make(map[string][]am.RawMessageHandler),
	}
}

func (s stream) Publish(ctx context.Context, topicName string, v am.RawMessage) error {
	for _, handler := range s.subscriptions[topicName] {
		err := handler.HandleMessage(ctx, &rawMessage{v})
		if err != nil {
			return err
		}
	}
	return nil
}

func (s stream) Subscribe(topicName string, handler am.MessageHandler[am.IncomingRawMessage], options ...am.SubscriberOption) (am.Subscription, error) {
	cfg := am.NewSubscriberConfig(options)

	var filter map[string]struct{}
	for _, key := range cfg.MessageFilters() {
		filter[key] = struct{}{}
	}

	fn := am.MessageHandlerFunc[am.IncomingRawMessage](func(ctx context.Context, msg am.IncomingRawMessage) error {
		if filter != nil {
			if _, exists := filter[msg.MessageName()]; !exists {
				return nil
			}
		}

		return handler.HandleMessage(ctx, msg)
	})

	s.subscriptions[topicName] = append(s.subscriptions[topicName], fn)

	return nil, nil
}

func (s stream) Unsubscribe() error {
	panic("unimplemented")
}
