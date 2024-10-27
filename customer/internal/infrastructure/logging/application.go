package logging

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/application/port/in/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/application/port/in/query"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/application/usecase"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/logger"
)

type Application struct {
	usecase.CustomerUsecase
	logger logger.Logger
}

var _ usecase.CustomerUsecase = (*Application)(nil)

func NewLogApplicationAccess(application usecase.CustomerUsecase, logger logger.Logger) *Application {
	return &Application{
		CustomerUsecase: application,
		logger:          logger,
	}
}

// RegisterCustomer implements usecase.CustomerUsecase.
func (a *Application) RegisterCustomer(ctx context.Context, register command.RegisterCustomer) (err error) {
	a.logger.Info("--> Customers.RegisterCustomer")
	defer func() { a.logger.WithError(err).Info("<-- Customers.RegisterCustomer") }()
	return a.CustomerUsecase.RegisterCustomer(ctx, register)
}

// AuthorizeCustomer implements usecase.CustomerUsecase.
func (a *Application) AuthorizeCustomer(ctx context.Context, authorize command.AuthorizeCustomer) (err error) {
	a.logger.Info("--> Customers.RegisterCustomer")
	defer func() { a.logger.WithError(err).Info("<-- Customers.RegisterCustomer") }()
	return a.CustomerUsecase.AuthorizeCustomer(ctx, authorize)
}

// EnableCustomer implements usecase.CustomerUsecase.
func (a *Application) EnableCustomer(ctx context.Context, enable command.EnableCustomer) (err error) {
	a.logger.Info("--> Customers.EnableCustomer")
	defer func() { a.logger.WithError(err).Info("<-- Customers.EnableCustomer") }()
	return a.CustomerUsecase.EnableCustomer(ctx, enable)
}

// DisableCustomer implements usecase.CustomerUsecase.
func (a *Application) DisableCustomer(ctx context.Context, disable command.DisableCustomer) (err error) {
	a.logger.Info("--> Customers.DisableCustomer")
	defer func() { a.logger.WithError(err).Info("<-- Customers.DisableCustomer") }()
	return a.CustomerUsecase.DisableCustomer(ctx, disable)
}

// GetCustomer implements usecase.CustomerUsecase.
func (a *Application) GetCustomer(ctx context.Context, query query.GetCustomer) (customer *aggregate.Customer, err error) {
	a.logger.Info("--> Customers.GetCustomer")
	defer func() { a.logger.WithError(err).Info("<-- Customers.GetCustomer") }()
	return a.CustomerUsecase.GetCustomer(ctx, query)
}
