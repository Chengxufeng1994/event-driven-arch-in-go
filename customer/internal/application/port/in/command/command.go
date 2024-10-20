package command

import "context"

type Commands interface {
	RegisterCustomer(ctx context.Context, register RegisterCustomer) error
	AuthorizeCustomer(ctx context.Context, authorize AuthorizeCustomer) error
	EnableCustomer(ctx context.Context, enable EnableCustomer) error
	DisableCustomer(ctx context.Context, disable DisableCustomer) error
}
