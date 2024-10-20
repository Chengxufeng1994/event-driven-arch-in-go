package command

import (
	"context"

	"github.com/stackus/errors"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/domain/repository"
)

type ConfirmPayment struct {
	ID string
}

func NewConfirmPayment(id string) ConfirmPayment {
	return ConfirmPayment{
		ID: id,
	}
}

type ConfirmPaymentHandler struct {
	paymentRepository repository.PaymentRepository
}

func NewConfirmPaymentHandler(paymentRepository repository.PaymentRepository) ConfirmPaymentHandler {
	return ConfirmPaymentHandler{
		paymentRepository: paymentRepository,
	}
}

func (h ConfirmPaymentHandler) ConfirmPayment(ctx context.Context, confirm ConfirmPayment) error {
	paymentEnt, err := h.paymentRepository.Find(ctx, confirm.ID)
	if err != nil || paymentEnt == nil {
		return errors.Wrap(errors.ErrNotFound, "confirm payment command")
	}

	return nil
}
