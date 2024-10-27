package command

import "context"

type Commands interface {
	CreateStore(ctx context.Context, cmd CreateStore) error
	RebrandStore(ctx context.Context, cmd RebrandStore) error
	EnableParticipation(ctx context.Context, cmd EnableParticipation) error
	DisableParticipation(ctx context.Context, cmd DisableParticipation) error
	AddProduct(ctx context.Context, cmd AddProduct) error
	RemoveProduct(ctx context.Context, cmd RemoveProduct) error
	RebrandProduct(ctx context.Context, cmd RebrandProduct) error
	IncreaseProductPrice(ctx context.Context, cmd IncreaseProductPrice) error
	DecreaseProductPrice(ctx context.Context, cmd DecreaseProductPrice) error
}
