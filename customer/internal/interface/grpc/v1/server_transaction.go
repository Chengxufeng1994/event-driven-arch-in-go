package v1

import (
	"context"

	customerv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/customer/api/customer/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/application/usecase"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/di"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type serverTx struct {
	c di.Container
	customerv1.UnimplementedCustomersServiceServer
}

var _ customerv1.CustomersServiceServer = (*server)(nil)

func RegisterServerTx(container di.Container, registrar grpc.ServiceRegistrar) error {
	customerv1.RegisterCustomersServiceServer(registrar, serverTx{
		c: container,
	})
	return nil
}

func (s serverTx) RegisterCustomer(ctx context.Context, request *customerv1.RegisterCustomerRequest) (resp *customerv1.RegisterCustomerResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *gorm.DB) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*gorm.DB))

	next := server{app: di.Get(ctx, "app").(usecase.CustomerUsecase)}

	return next.RegisterCustomer(ctx, request)
}

func (s serverTx) ChangeSmsNumber(ctx context.Context, request *customerv1.ChangeSmsNumberRequest) (resp *customerv1.ChangeSmsNumberResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *gorm.DB) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*gorm.DB))

	next := server{app: di.Get(ctx, "app").(usecase.CustomerUsecase)}

	return next.ChangeSmsNumber(ctx, request)
}

func (s serverTx) AuthorizeCustomer(ctx context.Context, request *customerv1.AuthorizeCustomerRequest) (resp *customerv1.AuthorizeCustomerResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *gorm.DB) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*gorm.DB))

	next := server{app: di.Get(ctx, "app").(usecase.CustomerUsecase)}

	return next.AuthorizeCustomer(ctx, request)
}

func (s serverTx) GetCustomer(ctx context.Context, request *customerv1.GetCustomerRequest) (resp *customerv1.GetCustomerResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *gorm.DB) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*gorm.DB))

	next := server{app: di.Get(ctx, "app").(usecase.CustomerUsecase)}

	return next.GetCustomer(ctx, request)
}

func (s serverTx) EnableCustomer(ctx context.Context, request *customerv1.EnableCustomerRequest) (resp *customerv1.EnableCustomerResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *gorm.DB) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*gorm.DB))

	next := server{app: di.Get(ctx, "app").(usecase.CustomerUsecase)}

	return next.EnableCustomer(ctx, request)
}

func (s serverTx) DisableCustomer(ctx context.Context, request *customerv1.DisableCustomerRequest) (resp *customerv1.DisableCustomerResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *gorm.DB) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*gorm.DB))

	next := server{app: di.Get(ctx, "app").(usecase.CustomerUsecase)}

	return next.DisableCustomer(ctx, request)
}

func (s serverTx) closeTx(tx *gorm.DB, err error) error {
	if r := recover(); r != nil {
		tx.Rollback()
		panic(r)
	}

	if err := tx.Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
