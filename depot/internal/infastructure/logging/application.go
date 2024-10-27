package logging

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/application/port/in/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/application/port/in/query"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/application/usecase"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/logger"
)

type Application struct {
	usecase.ShoppingListUseCase
	logger logger.Logger
}

var _ usecase.ShoppingListUseCase = (*Application)(nil)

func NewLogApplicationAccess(shoppingListApplication usecase.ShoppingListUseCase, logger logger.Logger) *Application {
	return &Application{
		ShoppingListUseCase: shoppingListApplication,
		logger:              logger,
	}
}

func (a *Application) CreateShoppingList(ctx context.Context, cmd command.CreateShoppingList) (err error) {
	a.logger.Info("--> Depot.CreateShoppingList")
	defer func() { a.logger.WithError(err).Info("<-- Depot.CreateShoppingList") }()
	return a.ShoppingListUseCase.CreateShoppingList(ctx, cmd)
}

func (a *Application) AssignShoppingList(ctx context.Context, cmd command.AssignShoppingList) (err error) {
	a.logger.Info("--> Depot.AssignShoppingList")
	defer func() { a.logger.WithError(err).Info("<-- Depot.AssignShoppingList") }()
	return a.ShoppingListUseCase.AssignShoppingList(ctx, cmd)
}

func (a *Application) CompleteShoppingList(ctx context.Context, cmd command.CompleteShoppingList) (err error) {
	a.logger.Info("--> Depot.CompleteShoppingList")
	defer func() { a.logger.WithError(err).Info("<-- Depot.CompleteShoppingList") }()
	return a.ShoppingListUseCase.CompleteShoppingList(ctx, cmd)
}

func (a *Application) CancelShoppingList(ctx context.Context, cmd command.CancelShoppingList) (err error) {
	a.logger.Info("--> Depot.CancelShoppingList")
	defer func() { a.logger.WithError(err).Info("<-- Depot.CancelShoppingList") }()
	return a.ShoppingListUseCase.CancelShoppingList(ctx, cmd)
}

func (a *Application) GetShoppingList(ctx context.Context, query query.GetShoppingList) (*aggregate.ShoppingList, error) {
	a.logger.Info("--> Depot.GetShoppingList")
	defer func() { a.logger.Info("<-- Depot.GetShoppingList") }()
	return a.ShoppingListUseCase.GetShoppingList(ctx, query)
}
