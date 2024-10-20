package client

import "context"

type OrderClient interface {
	Ready(ctx context.Context, orderID string) error
}
