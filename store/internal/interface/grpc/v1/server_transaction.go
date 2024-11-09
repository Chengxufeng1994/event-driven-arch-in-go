package v1

import (
	"context"

	"google.golang.org/grpc"
	"gorm.io/gorm"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/di"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/application/usecase"
)

type serverTx struct {
	c di.Container
	storev1.UnimplementedStoresServiceServer
}

var _ storev1.StoresServiceServer = (*serverTx)(nil)

func RegisterServerTx(container di.Container, registrar grpc.ServiceRegistrar) error {
	storev1.RegisterStoresServiceServer(registrar, serverTx{
		c: container,
	})
	return nil
}

func (s serverTx) CreateStore(ctx context.Context, request *storev1.CreateStoreRequest) (resp *storev1.CreateStoreResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *gorm.DB) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*gorm.DB))

	next := server{app: di.Get(ctx, "app").(usecase.StoreUseCase)}

	return next.CreateStore(ctx, request)
}

func (s serverTx) EnableParticipation(ctx context.Context, request *storev1.EnableParticipationRequest) (resp *storev1.EnableParticipationResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *gorm.DB) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*gorm.DB))

	next := server{app: di.Get(ctx, "app").(usecase.StoreUseCase)}

	return next.EnableParticipation(ctx, request)
}

func (s serverTx) DisableParticipation(ctx context.Context, request *storev1.DisableParticipationRequest) (resp *storev1.DisableParticipationResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *gorm.DB) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*gorm.DB))

	next := server{app: di.Get(ctx, "app").(usecase.StoreUseCase)}

	return next.DisableParticipation(ctx, request)
}

func (s serverTx) RebrandStore(ctx context.Context, request *storev1.RebrandStoreRequest) (resp *storev1.RebrandStoreResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *gorm.DB) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*gorm.DB))

	next := server{app: di.Get(ctx, "app").(usecase.StoreUseCase)}

	return next.RebrandStore(ctx, request)
}
func (s serverTx) GetStore(ctx context.Context, request *storev1.GetStoreRequest) (resp *storev1.GetStoreResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *gorm.DB) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*gorm.DB))

	next := server{app: di.Get(ctx, "app").(usecase.StoreUseCase)}

	return next.GetStore(ctx, request)
}

func (s serverTx) GetStores(ctx context.Context, request *storev1.GetStoresRequest) (resp *storev1.GetStoresResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *gorm.DB) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*gorm.DB))

	next := server{app: di.Get(ctx, "app").(usecase.StoreUseCase)}

	return next.GetStores(ctx, request)
}

func (s serverTx) GetParticipatingStores(ctx context.Context, request *storev1.GetParticipatingStoresRequest) (resp *storev1.GetParticipatingStoresResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *gorm.DB) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*gorm.DB))

	next := server{app: di.Get(ctx, "app").(usecase.StoreUseCase)}

	return next.GetParticipatingStores(ctx, request)
}

func (s serverTx) AddProduct(ctx context.Context, request *storev1.AddProductRequest) (resp *storev1.AddProductResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *gorm.DB) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*gorm.DB))

	next := server{app: di.Get(ctx, "app").(usecase.StoreUseCase)}

	return next.AddProduct(ctx, request)
}

func (s serverTx) RebrandProduct(ctx context.Context, request *storev1.RebrandProductRequest) (resp *storev1.RebrandProductResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *gorm.DB) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*gorm.DB))

	next := server{app: di.Get(ctx, "app").(usecase.StoreUseCase)}

	return next.RebrandProduct(ctx, request)
}

func (s serverTx) IncreaseProductPrice(ctx context.Context, request *storev1.IncreaseProductPriceRequest) (resp *storev1.IncreaseProductPriceResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *gorm.DB) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*gorm.DB))

	next := server{app: di.Get(ctx, "app").(usecase.StoreUseCase)}

	return next.IncreaseProductPrice(ctx, request)
}

func (s serverTx) DecreaseProductPrice(ctx context.Context, request *storev1.DecreaseProductPriceRequest) (resp *storev1.DecreaseProductPriceResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *gorm.DB) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*gorm.DB))

	next := server{app: di.Get(ctx, "app").(usecase.StoreUseCase)}

	return next.DecreaseProductPrice(ctx, request)
}

func (s serverTx) RemoveProduct(ctx context.Context, request *storev1.RemoveProductRequest) (resp *storev1.RemoveProductResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *gorm.DB) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*gorm.DB))

	next := server{app: di.Get(ctx, "app").(usecase.StoreUseCase)}

	return next.RemoveProduct(ctx, request)
}

func (s serverTx) GetProduct(ctx context.Context, request *storev1.GetProductRequest) (resp *storev1.GetProductResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *gorm.DB) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*gorm.DB))

	next := server{app: di.Get(ctx, "app").(usecase.StoreUseCase)}

	return next.GetProduct(ctx, request)
}

func (s serverTx) GetCatalog(ctx context.Context, request *storev1.GetCatalogRequest) (resp *storev1.GetCatalogResponse, err error) {

	ctx = s.c.Scoped(ctx)
	defer func(tx *gorm.DB) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*gorm.DB))

	next := server{app: di.Get(ctx, "app").(usecase.StoreUseCase)}

	return next.GetCatalog(ctx, request)
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
