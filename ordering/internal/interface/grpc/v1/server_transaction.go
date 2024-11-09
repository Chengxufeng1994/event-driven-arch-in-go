package v1

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/di"
	orderv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/api/order/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application/usecase"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type serverTx struct {
	c di.Container
	orderv1.UnimplementedOrderingServiceServer
}

var _ orderv1.OrderingServiceServer = (*server)(nil)

func RegisterServerTx(container di.Container, registrar grpc.ServiceRegistrar) error {
	orderv1.RegisterOrderingServiceServer(registrar, serverTx{c: container})
	return nil
}

func (s serverTx) CreateOrder(ctx context.Context, request *orderv1.CreateOrderRequest) (resp *orderv1.CreateOrderResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *gorm.DB) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*gorm.DB))

	next := server{app: di.Get(ctx, "app").(usecase.OrderUseCase)}

	return next.CreateOrder(ctx, request)
}

func (s serverTx) CancelOrder(ctx context.Context, request *orderv1.CancelOrderRequest) (resp *orderv1.CancelOrderResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *gorm.DB) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*gorm.DB))

	next := server{app: di.Get(ctx, "app").(usecase.OrderUseCase)}

	return next.CancelOrder(ctx, request)
}

func (s serverTx) ReadyOrder(ctx context.Context, request *orderv1.ReadyOrderRequest) (resp *orderv1.ReadyOrderResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *gorm.DB) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*gorm.DB))

	next := server{app: di.Get(ctx, "app").(usecase.OrderUseCase)}

	return next.ReadyOrder(ctx, request)
}

func (s serverTx) CompleteOrder(ctx context.Context, request *orderv1.CompleteOrderRequest) (resp *orderv1.CompleteOrderResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *gorm.DB) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*gorm.DB))

	next := server{app: di.Get(ctx, "app").(usecase.OrderUseCase)}

	return next.CompleteOrder(ctx, request)
}

func (s serverTx) GetOrder(ctx context.Context, request *orderv1.GetOrderRequest) (resp *orderv1.GetOrderResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *gorm.DB) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*gorm.DB))

	next := server{app: di.Get(ctx, "app").(usecase.OrderUseCase)}

	return next.GetOrder(ctx, request)
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
