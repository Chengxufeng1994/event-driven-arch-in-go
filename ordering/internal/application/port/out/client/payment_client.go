package client

import "context"

type PaymentClient interface {
	Confirm(ctx context.Context, paymentID string) error
}
