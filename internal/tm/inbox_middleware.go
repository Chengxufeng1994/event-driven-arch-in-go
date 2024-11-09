package tm

import (
	"context"
	"fmt"

	"github.com/stackus/errors"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
)

type ErrDuplicateMessage string

func (e ErrDuplicateMessage) Error() string {
	return fmt.Sprintf("duplicate message id encountered: %s", string(e))
}

type InboxStore interface {
	Save(ctx context.Context, msg am.RawMessage) error
}

type inbox struct {
	handler am.RawMessageHandler
	store   InboxStore
}

var _ am.RawMessageHandler = (*inbox)(nil)

func NewInboxHandlerMiddleware(store InboxStore) am.RawMessageHandlerMiddleware {
	i := inbox{store: store}

	return func(handler am.RawMessageHandler) am.RawMessageHandler {
		i.handler = handler

		return i
	}
}

func (i inbox) HandleMessage(ctx context.Context, msg am.IncomingRawMessage) error {
	// try to insert the message
	err := i.store.Save(ctx, msg)
	if err != nil {
		var errDupe ErrDuplicateMessage
		if errors.As(err, &errDupe) {
			// duplicate message; return without an error to let the message Ack
			return nil
		}
		return err
	}

	return i.handler.HandleMessage(ctx, msg)
}
