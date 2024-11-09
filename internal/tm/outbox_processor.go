package tm

import (
	"context"
	"time"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
)

const (
	messageLimit    = 50
	pollingInterval = 500 * time.Millisecond
)

type OutboxProcessor interface {
	Start(ctx context.Context) error
}

type outboxProcessor struct {
	publisher am.RawMessagePublisher
	store     OutboxStore
}

var _ OutboxProcessor = (*outboxProcessor)(nil)

func NewOutboxProcessor(publisher am.RawMessagePublisher, store OutboxStore) OutboxProcessor {
	return &outboxProcessor{
		publisher: publisher,
		store:     store,
	}
}

func (p outboxProcessor) Start(ctx context.Context) error {
	errc := make(chan error)

	go func() {
		errc <- p.processMessages(ctx)
	}()

	for {
		select {
		case <-ctx.Done():
			return nil
		case err := <-errc:
			return err
		}
	}
}

func (p outboxProcessor) processMessages(ctx context.Context) error {
	for {
		msgs, err := p.store.FindUnpublished(ctx, messageLimit)
		if err != nil {
			return err
		}

		if len(msgs) > 0 {
			ids := make([]string, len(msgs))
			for i, msg := range msgs {
				ids[i] = msg.ID()
				err = p.publisher.Publish(ctx, msg.Subject(), msg)
				if err != nil {
					return err
				}
			}
			err = p.store.MarkPublished(ctx, ids...)
			if err != nil {
				return err
			}

			continue
		}

		select {
		case <-ctx.Done():
			return nil
		// wait a short time before polling again
		case <-time.After(pollingInterval):
		}
	}
}
