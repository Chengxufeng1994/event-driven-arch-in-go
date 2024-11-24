package v1

import (
	"context"

	"google.golang.org/grpc"
	"gorm.io/gorm"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/di"
	searchv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/search/api/search/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/application/usecase"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/infrastructure/constants"
)

type serverTx struct {
	c di.Container
	searchv1.UnimplementedSearchServiceServer
}

var _ searchv1.SearchServiceServer = (*serverTx)(nil)

func RegisterServerTx(container di.Container, register grpc.ServiceRegistrar) error {
	searchv1.RegisterSearchServiceServer(register, serverTx{
		c: container,
	})
	return nil
}

func (s serverTx) GetOrder(ctx context.Context, request *searchv1.GetOrderRequest) (resp *searchv1.GetOrderResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *gorm.DB) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*gorm.DB))

	next := server{app: di.Get(ctx, constants.ApplicationKey).(usecase.SearchUseCase)}

	return next.GetOrder(ctx, request)
}

func (s serverTx) SearchOrders(ctx context.Context, request *searchv1.SearchOrdersRequest) (resp *searchv1.SearchOrdersResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *gorm.DB) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*gorm.DB))

	next := server{app: di.Get(ctx, constants.ApplicationKey).(usecase.SearchUseCase)}

	return next.SearchOrders(ctx, request)
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
