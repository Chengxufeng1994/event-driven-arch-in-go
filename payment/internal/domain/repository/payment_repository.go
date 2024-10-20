package repository

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/domain/aggregate"
)

type PaymentRepository interface {
	Save(ctx context.Context, payment *aggregate.PaymentAgg) error
	Find(ctx context.Context, paymentID string) (*aggregate.PaymentAgg, error)
}
