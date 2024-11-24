package v1

import (
	"context"

	"google.golang.org/grpc"
	"gorm.io/gorm"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/di"
	paymentv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/payment/api/payment/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/application/usecase"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/infrastructure/constants"
)

type serverTx struct {
	c di.Container
	paymentv1.UnimplementedPaymentsServiceServer
}

var _ paymentv1.PaymentsServiceServer = (*server)(nil)

func RegisterServerTx(container di.Container, registrar grpc.ServiceRegistrar) error {
	paymentv1.RegisterPaymentsServiceServer(registrar, serverTx{c: container})
	return nil
}

func (s serverTx) AuthorizePayment(ctx context.Context, request *paymentv1.AuthorizePaymentRequest) (resp *paymentv1.AuthorizePaymentResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *gorm.DB) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*gorm.DB))

	next := server{app: di.Get(ctx, constants.ApplicationKey).(usecase.PaymentUseCase)}

	return next.AuthorizePayment(ctx, request)
}

func (s serverTx) ConfirmPayment(ctx context.Context, request *paymentv1.ConfirmPaymentRequest) (resp *paymentv1.ConfirmPaymentResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *gorm.DB) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*gorm.DB))

	next := server{app: di.Get(ctx, constants.ApplicationKey).(usecase.PaymentUseCase)}

	return next.ConfirmPayment(ctx, request)
}

func (s serverTx) CreateInvoice(ctx context.Context, request *paymentv1.CreateInvoiceRequest) (resp *paymentv1.CreateInvoiceResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *gorm.DB) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*gorm.DB))

	next := server{app: di.Get(ctx, constants.ApplicationKey).(usecase.PaymentUseCase)}

	return next.CreateInvoice(ctx, request)
}

func (s serverTx) CancelInvoice(ctx context.Context, request *paymentv1.CancelInvoiceRequest) (resp *paymentv1.CancelInvoiceResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *gorm.DB) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*gorm.DB))

	next := server{app: di.Get(ctx, constants.ApplicationKey).(usecase.PaymentUseCase)}

	return next.CancelInvoice(ctx, request)
}

func (s serverTx) AdjustInvoice(ctx context.Context, request *paymentv1.AdjustInvoiceRequest) (resp *paymentv1.AdjustInvoiceResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *gorm.DB) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*gorm.DB))

	next := server{app: di.Get(ctx, constants.ApplicationKey).(usecase.PaymentUseCase)}

	return next.AdjustInvoice(ctx, request)
}

func (s serverTx) PayInvoice(ctx context.Context, request *paymentv1.PayInvoiceRequest) (resp *paymentv1.PayInvoiceResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *gorm.DB) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*gorm.DB))

	next := server{app: di.Get(ctx, constants.ApplicationKey).(usecase.PaymentUseCase)}

	return next.PayInvoice(ctx, request)
}

func (s serverTx) closeTx(tx *gorm.DB, err error) error {
	if p := recover(); p != nil {
		_ = tx.Rollback()
		panic(p)
	} else if err != nil {
		_ = tx.Rollback()
		return err
	} else {
		return tx.Commit().Error
	}
}
