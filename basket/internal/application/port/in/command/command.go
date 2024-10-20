package command

import "context"

type Commands interface {
	StartBasket(ctx context.Context, cmd StartBasket) error
	CancelBasket(ctx context.Context, cmd CancelBasket) error
	CheckoutBasket(ctx context.Context, checkout CheckoutBasket) error
	AddItem(ctx context.Context, add AddItem) error
	RemoveItem(ctx context.Context, remove RemoveItem) error
}
