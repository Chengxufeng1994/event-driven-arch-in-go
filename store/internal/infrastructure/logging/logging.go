package logging

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/logger"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/application/port/in/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/application/port/in/query"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/application/usecase"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
)

type Application struct {
	usecase.StoreUseCase
	logger logger.Logger
}

var _ usecase.StoreUseCase = (*Application)(nil)

func NewLogApplicationAccess(application usecase.StoreUseCase, logger logger.Logger) Application {
	return Application{
		StoreUseCase: application,
		logger:       logger,
	}
}

func (a Application) CreateStore(ctx context.Context, cmd command.CreateStore) (err error) {
	a.logger.Info("--> Stores.CreateStore")
	defer func() { a.logger.WithError(err).Info("<-- Stores.CreateStore") }()
	return a.StoreUseCase.CreateStore(ctx, cmd)
}

func (a Application) EnableParticipation(ctx context.Context, cmd command.EnableParticipation) (err error) {
	a.logger.Info("--> Stores.EnableParticipation")
	defer func() { a.logger.WithError(err).Info("<-- Stores.EnableParticipation") }()
	return a.StoreUseCase.EnableParticipation(ctx, cmd)
}

func (a Application) DisableParticipation(ctx context.Context, cmd command.DisableParticipation) (err error) {
	a.logger.Info("--> Stores.DisableParticipation")
	defer func() { a.logger.WithError(err).Info("<-- Stores.DisableParticipation") }()
	return a.StoreUseCase.DisableParticipation(ctx, cmd)
}

func (a Application) AddProduct(ctx context.Context, cmd command.AddProduct) (err error) {
	a.logger.Info("--> Stores.AddProduct")
	defer func() { a.logger.WithError(err).Info("<-- Stores.AddProduct") }()
	return a.StoreUseCase.AddProduct(ctx, cmd)
}

func (a Application) RemoveProduct(ctx context.Context, cmd command.RemoveProduct) (err error) {
	a.logger.Info("--> Stores.RemoveProduct")
	defer func() { a.logger.WithError(err).Info("<-- Stores.RemoveProduct") }()
	return a.StoreUseCase.RemoveProduct(ctx, cmd)
}

func (a Application) GetStore(ctx context.Context, query query.GetStore) (store *aggregate.Store, err error) {
	a.logger.Info("--> Stores.GetStore")
	defer func() { a.logger.WithError(err).Info("<-- Stores.GetStore") }()
	return a.StoreUseCase.GetStore(ctx, query)
}

func (a Application) GetStores(ctx context.Context, query query.GetStores) (store []*aggregate.Store, err error) {
	a.logger.Info("--> Stores.GetStores")
	defer func() { a.logger.WithError(err).Info("<-- Stores.GetStores") }()
	return a.StoreUseCase.GetStores(ctx, query)
}

func (a Application) GetParticipatingStores(ctx context.Context, query query.GetParticipatingStores) (store []*aggregate.Store, err error) {
	a.logger.Info("--> Stores.GetParticipatingStores")
	defer func() { a.logger.WithError(err).Info("<-- Stores.GetParticipatingStores") }()
	return a.StoreUseCase.GetParticipatingStores(ctx, query)
}

func (a Application) GetCatalog(ctx context.Context, query query.GetCatalog) (products []*aggregate.Product, err error) {
	a.logger.Info("--> Stores.GetCatalog")
	defer func() { a.logger.WithError(err).Info("<-- Stores.GetCatalog") }()
	return a.StoreUseCase.GetCatalog(ctx, query)
}

func (a Application) GetProduct(ctx context.Context, query query.GetProduct) (product *aggregate.Product, err error) {
	a.logger.Info("--> Stores.GetProduct")
	defer func() { a.logger.WithError(err).Info("<-- Stores.GetProduct") }()
	return a.StoreUseCase.GetProduct(ctx, query)
}
