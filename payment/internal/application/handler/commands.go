package handler

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	paymentv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/payment/api/payment/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/application/port/int/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/application/usecase"
)

type commandHandler struct {
	app usecase.PaymentUseCase
}

func NewCommandHandler(app usecase.PaymentUseCase) *commandHandler {
	return &commandHandler{
		app: app,
	}
}

func RegisterCommandHandlers(subscriber am.RawMessageStream, handlers am.RawMessageHandler) error {
	_, err := subscriber.Subscribe(paymentv1.CommandChannel, handlers, am.MessageFilter{
		paymentv1.ConfirmPaymentCommand,
	}, am.GroupName("payment-commands"))
	return err
}

func (h commandHandler) HandleCommand(ctx context.Context, cmd ddd.Command) (ddd.Reply, error) {
	switch cmd.CommandName() {
	case paymentv1.ConfirmPaymentCommand:
		return h.doConfirmPayment(ctx, cmd)
	}

	return nil, nil
}

func (h commandHandler) doConfirmPayment(ctx context.Context, cmd ddd.Command) (ddd.Reply, error) {
	payload := cmd.Payload().(*paymentv1.ConfirmPayment)

	return nil, h.app.ConfirmPayment(ctx, command.ConfirmPayment{ID: payload.GetId()})
}
