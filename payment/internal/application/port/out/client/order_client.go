package client

import "context"

type OrderClient interface {
	Complete(ctx context.Context, invoiceID, orderID string) error
}
