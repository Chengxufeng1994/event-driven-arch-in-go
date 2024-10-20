package logging

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/port/in/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/port/in/query"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/usecase"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/logger"
)

type Application struct {
	usecase.BasketUseCase
	logger logger.Logger
}

var _ usecase.BasketUseCase = (*Application)(nil)

func NewLogApplicationAccess(application usecase.BasketUseCase, logger logger.Logger) *Application {
	return &Application{
		BasketUseCase: application,
		logger:        logger,
	}
}

// StartBasket implements usecase.BasketUseCase.
func (a *Application) StartBasket(ctx context.Context, cmd command.StartBasket) (err error) {
	a.logger.Info("--> Baskets.StartBasket")
	defer func() { a.logger.WithError(err).Info("<-- Baskets.StartBasket") }()
	return a.BasketUseCase.StartBasket(ctx, cmd)
}

// CancelBasket implements usecase.BasketUseCase.
func (a *Application) CancelBasket(ctx context.Context, cmd command.CancelBasket) (err error) {
	a.logger.Info("--> Baskets.CancelBasket")
	defer func() { a.logger.WithError(err).Info("<-- Baskets.CancelBasket") }()
	return a.BasketUseCase.CancelBasket(ctx, cmd)
}

// CheckoutBasket implements usecase.BasketUseCase.
func (a *Application) CheckoutBasket(ctx context.Context, checkout command.CheckoutBasket) (err error) {
	a.logger.Info("--> Baskets.CheckoutBasket")
	defer func() { a.logger.WithError(err).Info("<-- Baskets.CheckBasket") }()
	return a.BasketUseCase.CheckoutBasket(ctx, checkout)
}

// AddItem implements usecase.BasketUseCase.
func (a *Application) AddItem(ctx context.Context, add command.AddItem) (err error) {
	a.logger.Info("--> Baskets.AddItem")
	defer func() { a.logger.WithError(err).Info("<-- Baskets.AddItem") }()
	return a.BasketUseCase.AddItem(ctx, add)
}

// RemoveItem implements usecase.BasketUseCase.
func (a *Application) RemoveItem(ctx context.Context, remove command.RemoveItem) (err error) {
	a.logger.Info("--> Baskets.RemoveItem")
	defer func() { a.logger.WithError(err).Info("<-- Baskets.RemoveItem") }()
	return a.BasketUseCase.RemoveItem(ctx, remove)
}

// GetBasket implements usecase.BasketUseCase.
func (a *Application) GetBasket(ctx context.Context, query query.GetBasket) (basket *aggregate.BasketAgg, err error) {
	a.logger.Info("--> Baskets.GetBasket")
	defer func() { a.logger.WithError(err).Info("<-- Baskets.GetBasket") }()
	return a.BasketUseCase.GetBasket(ctx, query)
}
