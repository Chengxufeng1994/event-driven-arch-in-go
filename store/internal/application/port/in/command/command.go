package command

import "context"

type Commands interface {
	CreateStore(ctx context.Context, cmd CreateStore) error
	EnableParticipation(ctx context.Context, cmd EnableParticipation) error
	DisableParticipation(ctx context.Context, cmd DisableParticipation) error
	AddProduct(ctx context.Context, cmd AddProduct) error
	RemoveProduct(ctx context.Context, cmd RemoveProduct) error
}
