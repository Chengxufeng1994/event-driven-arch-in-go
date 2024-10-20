package command

import "context"

type Commands interface {
	CreateShoppingList(ctx context.Context, cmd CreateShoppingList) error
	CancelShoppingList(ctx context.Context, cmd CancelShoppingList) error
	AssignShoppingList(ctx context.Context, cmd AssignShoppingList) error
	CompleteShoppingList(ctx context.Context, cmd CompleteShoppingList) error
}
