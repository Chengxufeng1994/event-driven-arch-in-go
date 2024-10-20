package repository

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/domain/aggregate"
)

type InvoiceRepository interface {
	Save(ctx context.Context, invoice *aggregate.InvoiceAgg) error
	Update(ctx context.Context, invoice *aggregate.InvoiceAgg) error
	Find(ctx context.Context, invoiceID string) (*aggregate.InvoiceAgg, error)
}
