package application

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/application/port/int/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/application/usecase"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/domain/repository"
)

type (
	PaymentApplication struct {
		appCommands
		appQueries
	}

	appCommands struct {
		command.AuthorizePaymentHandler
		command.ConfirmPaymentHandler
		command.CreateInvoiceHandler
		command.AdjustInvoiceHandler
		command.PayInvoiceHandler
		command.CancelInvoiceHandler
	}

	appQueries struct{}
)

var _ usecase.PaymentUseCase = (*PaymentApplication)(nil)

func NewPaymentApplication(
	invoiceRepository repository.InvoiceRepository,
	paymentRepository repository.PaymentRepository,
	domainEventPublisher ddd.EventPublisher[ddd.Event],
) *PaymentApplication {
	return &PaymentApplication{
		appCommands: appCommands{
			AuthorizePaymentHandler: command.NewAuthorizePaymentCommandHandler(paymentRepository),
			ConfirmPaymentHandler:   command.NewConfirmPaymentHandler(paymentRepository),
			CreateInvoiceHandler:    command.NewCreateInvoiceHandler(invoiceRepository),
			AdjustInvoiceHandler:    command.NewAdjustInvoiceHandler(invoiceRepository),
			PayInvoiceHandler:       command.NewPayInvoiceHandler(invoiceRepository, domainEventPublisher),
			CancelInvoiceHandler:    command.NewCancelInvoiceHandler(invoiceRepository),
		},
		appQueries: appQueries{},
	}
}
