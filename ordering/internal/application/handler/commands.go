package handler

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	orderv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/api/order/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application/port/in/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application/usecase"
)

type commandHandlers struct {
	app usecase.OrderUseCase
}

var _ ddd.CommandHandler[ddd.Command] = (*commandHandlers)(nil)

func NewCommandHandlers(app usecase.OrderUseCase) ddd.CommandHandler[ddd.Command] {
	return commandHandlers{
		app: app,
	}
}

func RegisterCommandHandlers(subscriber am.RawMessageSubscriber, handlers am.RawMessageHandler) error {
	_, err := subscriber.Subscribe(orderv1.CommandChannel, handlers, am.MessageFilter{
		orderv1.RejectOrderCommand,
		orderv1.ApproveOrderCommand,
	}, am.GroupName("ordering-commands"))

	return err
}

func (h commandHandlers) HandleCommand(ctx context.Context, cmd ddd.Command) (ddd.Reply, error) {
	switch cmd.CommandName() {
	case orderv1.RejectOrderCommand:
		return h.doRejectOrder(ctx, cmd)
	case orderv1.ApproveOrderCommand:
		return h.doApproveOrder(ctx, cmd)
	}

	return nil, nil
}

func (h commandHandlers) doRejectOrder(ctx context.Context, cmd ddd.Command) (ddd.Reply, error) {
	payload := cmd.Payload().(*orderv1.RejectOrder)

	return nil, h.app.RejectOrder(ctx, command.RejectOrder{ID: payload.GetId()})
}

func (h commandHandlers) doApproveOrder(ctx context.Context, cmd ddd.Command) (ddd.Reply, error) {
	payload := cmd.Payload().(*orderv1.ApproveOrder)

	return nil, h.app.ApproveOrder(ctx, command.ApproveOrder{
		ID:         payload.GetId(),
		ShoppingID: payload.GetShoppingId(),
	})
}
