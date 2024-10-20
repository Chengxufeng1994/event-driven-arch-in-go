package command

import "context"

type Commands interface {
	AuthorizePayment(ctx context.Context, authorize AuthorizePayment) error
	ConfirmPayment(ctx context.Context, confirm ConfirmPayment) error
	CreateInvoice(ctx context.Context, create CreateInvoice) error
	AdjustInvoice(ctx context.Context, adjust AdjustInvoice) error
	PayInvoice(ctx context.Context, pay PayInvoice) error
	CancelInvoice(ctx context.Context, cancel CancelInvoice) error
}
