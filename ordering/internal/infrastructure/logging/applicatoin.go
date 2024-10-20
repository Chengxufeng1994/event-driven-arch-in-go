package logging

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/logger"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application/port/in/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application/port/in/query"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application/usecase"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/aggregate"
)

type Application struct {
	usecase.OrderUseCase
	logger logger.Logger
}

var _ usecase.OrderUseCase = (*Application)(nil)

func NewLogApplicationAccess(application usecase.OrderUseCase, logger logger.Logger) *Application {
	return &Application{
		OrderUseCase: application,
		logger:       logger,
	}
}

// CreateOrder implements usecase.OrderUseCase.
func (a *Application) CreateOrder(ctx context.Context, cmd command.CreateOrder) (err error) {
	a.logger.Info("--> Ordering.CreateOrder")
	defer func() { a.logger.WithError(err).Info("<-- Ordering.CreateOrder") }()
	return a.OrderUseCase.CreateOrder(ctx, cmd)
}

// CancelOrder implements usecase.OrderUseCase.
func (a *Application) CancelOrder(ctx context.Context, cmd command.CancelOrder) (err error) {
	a.logger.Info("--> Ordering.CancelOrder")
	defer func() { a.logger.WithError(err).Info("<-- Ordering.CancelOrder") }()
	return a.OrderUseCase.CancelOrder(ctx, cmd)
}

// ReadyOrder implements usecase.OrderUseCase.
func (a *Application) ReadyOrder(ctx context.Context, cmd command.ReadyOrder) (err error) {
	a.logger.Info("--> Ordering.ReadyOrder")
	defer func() { a.logger.WithError(err).Info("<-- Ordering.ReadyOrder") }()
	return a.OrderUseCase.ReadyOrder(ctx, cmd)
}

// CompleteOrder implements usecase.OrderUseCase.
func (a *Application) CompleteOrder(ctx context.Context, cmd command.CompleteOrder) (err error) {
	a.logger.Info("--> Ordering.CompleteOrder")
	defer func() { a.logger.WithError(err).Info("<-- Ordering.CompleteOrder") }()
	return a.OrderUseCase.CompleteOrder(ctx, cmd)
}

// GetOrder implements usecase.OrderUseCase.
func (a *Application) GetOrder(ctx context.Context, query query.GetOrder) (order *aggregate.OrderAgg, err error) {
	a.logger.Info("--> Ordering.GetOrder")
	defer func() { a.logger.WithError(err).Info("<-- Ordering.GetOrder") }()
	return a.OrderUseCase.GetOrder(ctx, query)
}
