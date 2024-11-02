package paymentv1

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry/serdes"
)

const (
	InvoiceAggregateChannel = "mallbots.payments.events.Invoice"

	InvoicePaidEvent = "paymentsapi.InvoicePaid"
)

func Registrations(reg registry.Registry) error {
	serde := serdes.NewProtoSerde(reg)

	// Invoice events
	if err := serde.Register(&InvoicePaid{}); err != nil {
		return err
	}

	return nil
}

func (*InvoicePaid) Key() string { return InvoicePaidEvent }
