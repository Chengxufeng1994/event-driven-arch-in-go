package v1

import (
	"context"

	"google.golang.org/grpc"
	"gorm.io/gorm"

	depotv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/depot/api/depot/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/application/usecase"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/di"
)

type serverTx struct {
	c di.Container
	depotv1.UnimplementedDepotServiceServer
}

var _ depotv1.DepotServiceServer = (*server)(nil)

func RegisterServerTx(container di.Container, registrar grpc.ServiceRegistrar) error {
	depotv1.RegisterDepotServiceServer(registrar, serverTx{
		c: container,
	})
	return nil
}

func (s serverTx) CreateShoppingList(ctx context.Context, request *depotv1.CreateShoppingListRequest) (resp *depotv1.CreateShoppingListResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *gorm.DB) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*gorm.DB))

	next := server{app: di.Get(ctx, "app").(usecase.ShoppingListUseCase)}

	return next.CreateShoppingList(ctx, request)
}

func (s serverTx) CancelShoppingList(ctx context.Context, request *depotv1.CancelShoppingListRequest) (resp *depotv1.CancelShoppingListResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *gorm.DB) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*gorm.DB))

	next := server{app: di.Get(ctx, "app").(usecase.ShoppingListUseCase)}

	return next.CancelShoppingList(ctx, request)
}

func (s serverTx) AssignShoppingList(ctx context.Context, request *depotv1.AssignShoppingListRequest) (resp *depotv1.AssignShoppingListResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *gorm.DB) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*gorm.DB))

	next := server{app: di.Get(ctx, "app").(usecase.ShoppingListUseCase)}

	return next.AssignShoppingList(ctx, request)
}

func (s serverTx) CompleteShoppingList(ctx context.Context, request *depotv1.CompleteShoppingListRequest) (resp *depotv1.CompleteShoppingListResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *gorm.DB) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*gorm.DB))

	next := server{app: di.Get(ctx, "app").(usecase.ShoppingListUseCase)}

	return next.CompleteShoppingList(ctx, request)
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
