package command

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/domain/repository"
)

type AuthorizePayment struct {
	ID         string
	CustomerID string
	Amount     float64
}

func NewAuthorizePayment(id, customerID string, amount float64) AuthorizePayment {
	return AuthorizePayment{
		ID:         id,
		CustomerID: customerID,
		Amount:     amount,
	}
}

type AuthorizePaymentHandler struct {
	payments repository.PaymentRepository
}

func NewAuthorizePaymentHandler(payments repository.PaymentRepository) AuthorizePaymentHandler {
	return AuthorizePaymentHandler{
		payments: payments,
	}
}

func (h AuthorizePaymentHandler) AuthorizePayment(ctx context.Context, authorize AuthorizePayment) error {
	return h.payments.Save(ctx, aggregate.NewPayment(
		authorize.ID,
		authorize.CustomerID,
		authorize.Amount))
}
