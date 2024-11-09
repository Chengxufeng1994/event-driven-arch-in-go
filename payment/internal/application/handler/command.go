package handler

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	paymentv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/payment/api/payment/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/application/port/int/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/application/usecase"
)

type CommandHandler[T ddd.Command] struct {
	app usecase.PaymentUseCase
}

var _ ddd.CommandHandler[ddd.Command] = (*CommandHandler[ddd.Command])(nil)

func NewCommandHandler(app usecase.PaymentUseCase) *CommandHandler[ddd.Command] {
	return &CommandHandler[ddd.Command]{
		app: app,
	}
}

func RegisterCommandHandler(subscriber am.CommandSubscriber, handlers ddd.CommandHandler[ddd.Command]) error {
	cmdMsgHandler := am.CommandMessageHandlerFunc(func(ctx context.Context, cmdMsg am.IncomingCommandMessage) (ddd.Reply, error) {
		return handlers.HandleCommand(ctx, cmdMsg)
	})

	return subscriber.Subscribe(paymentv1.CommandChannel, cmdMsgHandler, am.MessageFilter{
		paymentv1.ConfirmPaymentCommand,
	}, am.GroupName("payment-commands"))
}

func (h *CommandHandler[T]) HandleCommand(ctx context.Context, cmd T) (ddd.Reply, error) {
	switch cmd.CommandName() {
	case paymentv1.ConfirmPaymentCommand:
		return h.doConfirmPayment(ctx, cmd)
	}

	return nil, nil
}

func (h *CommandHandler[T]) doConfirmPayment(ctx context.Context, cmd T) (ddd.Reply, error) {
	payload := cmd.Payload().(*paymentv1.ConfirmPayment)
	return nil, h.app.ConfirmPayment(ctx, command.ConfirmPayment{ID: payload.GetId()})
}
