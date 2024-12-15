package tm

import (
	"context"
	"sync"
	"time"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"golang.org/x/sync/errgroup"
)

const (
	workerCount = 5 // Worker 數量
)

type outboxProcessorOption func(*outboxProcessorWithWorker) *outboxProcessorWithWorker

func WithWorkerCount(workerCount int) outboxProcessorOption {
	return func(p *outboxProcessorWithWorker) *outboxProcessorWithWorker {
		p.workerCount = workerCount
		return p
	}
}

type outboxProcessorWithWorker struct {
	publisher   am.MessagePublisher
	store       OutboxStore
	workerCount int
}

var _ OutboxProcessor = (*outboxProcessorWithWorker)(nil)

func NewOutboxProcessorWithOptions(publisher am.MessagePublisher, store OutboxStore, options ...outboxProcessorOption) OutboxProcessor {
	op := &outboxProcessorWithWorker{
		publisher:   publisher,
		store:       store,
		workerCount: workerCount,
	}

	for _, option := range options {
		op = option(op)
	}

	return op
}

func (p *outboxProcessorWithWorker) Start(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	for i := 0; i < p.workerCount; i++ {
		g.Go(func() error {
			return p.processMessages(ctx)
		})
	}

	return g.Wait()
}

func (p *outboxProcessorWithWorker) processMessages(ctx context.Context) error {
	timer := time.NewTimer(0)
	defer timer.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}

		msgs, err := p.store.FindUnpublished(ctx, messageLimit)
		if err != nil {
			return err
		}

		if len(msgs) > 0 {
			if err := p.publishMessages(ctx, msgs); err != nil {
				return err
			}
			continue // 繼續處理下一批訊息
		}

		if !timer.Stop() {
			select {
			case <-timer.C:
			default:
			}
		}
		timer.Reset(pollingInterval)

		select {
		case <-ctx.Done():
			return nil
		case <-timer.C:
		}
	}
}

func (p *outboxProcessorWithWorker) publishMessages(ctx context.Context, msgs []am.Message) error {
	var mu sync.Mutex
	ids := make([]string, 0, len(msgs))

	for _, msg := range msgs {
		if err := p.publisher.Publish(ctx, msg.Subject(), msg); err != nil {
			return err
		}
		mu.Lock()
		ids = append(ids, msg.ID())
		mu.Unlock()
	}

	return p.store.MarkPublished(ctx, ids...)
}
