package command

import "context"

type Commands interface {
	NotifyOrderCreated(context.Context, OrderCreated) error
	NotifyOrderReady(context.Context, OrderReady) error
	NotifyOrderCanceled(context.Context, OrderCanceled) error
}
