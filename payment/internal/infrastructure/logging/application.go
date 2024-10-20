package logging

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/logger"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/application/port/int/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/application/usecase"
)

type Application struct {
	usecase.PaymentUseCase
	logger logger.Logger
}

var _ usecase.PaymentUseCase = (*Application)(nil)

func NewLogApplicationAccess(app usecase.PaymentUseCase, logger logger.Logger) *Application {
	return &Application{
		PaymentUseCase: app,
		logger:         logger,
	}
}

func (a *Application) AuthorizePayment(ctx context.Context, authorize command.AuthorizePayment) (err error) {
	a.logger.Info("--> Payments.AuthorizePayment")
	defer func() { a.logger.WithError(err).Info("<-- Payments.AuthorizePayment") }()
	return a.PaymentUseCase.AuthorizePayment(ctx, authorize)
}

// ConfirmPayment implements usecase.PaymentUseCase.
func (a *Application) ConfirmPayment(ctx context.Context, confirm command.ConfirmPayment) (err error) {
	a.logger.Info("--> Payments.ConfirmPayment")
	defer func() { a.logger.WithError(err).Info("<-- Payments.ConfirmPayment") }()
	return a.PaymentUseCase.ConfirmPayment(ctx, confirm)
}

// CreateInvoice implements usecase.PaymentUseCase.
func (a *Application) CreateInvoice(ctx context.Context, create command.CreateInvoice) (err error) {
	a.logger.Info("--> Payments.CreateInvoice")
	defer func() { a.logger.WithError(err).Info("<-- Payments.CreateInvoice") }()
	return a.PaymentUseCase.CreateInvoice(ctx, create)
}

// AdjustInvoice implements usecase.PaymentUseCase.
func (a *Application) AdjustInvoice(ctx context.Context, adjust command.AdjustInvoice) (err error) {
	a.logger.Info("--> Payments.AdjustInvoice")
	defer func() { a.logger.WithError(err).Info("<-- Payments.AdjustInvoice") }()
	return a.PaymentUseCase.AdjustInvoice(ctx, adjust)
}

// PayInvoice implements usecase.PaymentUseCase.
func (a *Application) PayInvoice(ctx context.Context, pay command.PayInvoice) (err error) {
	a.logger.Info("--> Payments.PayInvoice")
	defer func() { a.logger.WithError(err).Info("<-- Payments.PayInvoice") }()
	return a.PaymentUseCase.PayInvoice(ctx, pay)
}

// CancelInvoice implements usecase.PaymentUseCase.
func (a *Application) CancelInvoice(ctx context.Context, cancel command.CancelInvoice) (err error) {
	a.logger.Info("--> Payments.CancelInvoice")
	defer func() { a.logger.WithError(err).Info("<-- Payments.CancelInvoice") }()
	return a.PaymentUseCase.CancelInvoice(ctx, cancel)
}
