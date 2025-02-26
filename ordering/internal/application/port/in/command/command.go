package command

import "context"

type Commands interface {
	CreateOrder(ctx context.Context, cmd CreateOrder) error
	RejectOrder(ctx context.Context, cmd RejectOrder) error
	ApproveOrder(ctx context.Context, cmd ApproveOrder) error
	CancelOrder(ctx context.Context, cmd CancelOrder) error
	ReadyOrder(ctx context.Context, cmd ReadyOrder) error
	CompleteOrder(ctx context.Context, cmd CompleteOrder) error
}
