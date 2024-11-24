package v1

import (
	"context"

	"google.golang.org/grpc"
	"gorm.io/gorm"

	basketv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/basket/api/basket/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/usecase"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/infrastructure/constants"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/di"
)

type serverTx struct {
	c di.Container
	basketv1.UnimplementedBasketServiceServer
}

var _ basketv1.BasketServiceServer = (*serverTx)(nil)

func RegisterServerTx(container di.Container, registrar grpc.ServiceRegistrar) error {
	basketv1.RegisterBasketServiceServer(registrar, serverTx{
		c: container,
	})
	return nil
}

func (s serverTx) StartBasket(ctx context.Context, request *basketv1.StartBasketRequest) (resp *basketv1.StartBasketResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *gorm.DB) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*gorm.DB))

	next := server{app: di.Get(ctx, constants.ApplicationKey).(usecase.BasketUseCase)}

	return next.StartBasket(ctx, request)
}

func (s serverTx) CancelBasket(ctx context.Context, request *basketv1.CancelBasketRequest) (resp *basketv1.CancelBasketResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *gorm.DB) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*gorm.DB))

	next := server{app: di.Get(ctx, constants.ApplicationKey).(usecase.BasketUseCase)}

	return next.CancelBasket(ctx, request)
}

func (s serverTx) CheckoutBasket(ctx context.Context, request *basketv1.CheckoutBasketRequest) (resp *basketv1.CheckoutBasketResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *gorm.DB) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*gorm.DB))

	next := server{app: di.Get(ctx, constants.ApplicationKey).(usecase.BasketUseCase)}

	return next.CheckoutBasket(ctx, request)
}

func (s serverTx) AddItem(ctx context.Context, request *basketv1.AddItemRequest) (resp *basketv1.AddItemResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *gorm.DB) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*gorm.DB))

	next := server{app: di.Get(ctx, constants.ApplicationKey).(usecase.BasketUseCase)}

	return next.AddItem(ctx, request)
}

func (s serverTx) RemoveItem(ctx context.Context, request *basketv1.RemoveItemRequest) (resp *basketv1.RemoveItemResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *gorm.DB) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*gorm.DB))

	next := server{app: di.Get(ctx, constants.ApplicationKey).(usecase.BasketUseCase)}

	return next.RemoveItem(ctx, request)
}

func (s serverTx) GetBasket(ctx context.Context, request *basketv1.GetBasketRequest) (resp *basketv1.GetBasketResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *gorm.DB) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*gorm.DB))

	next := server{app: di.Get(ctx, constants.ApplicationKey).(usecase.BasketUseCase)}

	return next.GetBasket(ctx, request)
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
