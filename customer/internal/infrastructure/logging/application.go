package logging

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/application/port/in/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/application/port/in/query"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/application/service"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/application/usecase"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/logger"
)

type Application struct {
	application *service.CustomerApplication
	logger      logger.Logger
}

var _ usecase.CustomerUsecase = (*Application)(nil)

func NewLogApplicationAccess(
	application *service.CustomerApplication,
	logger logger.Logger,
) *Application {
	return &Application{
		application: application,
		logger:      logger,
	}
}

// RegisterCustomer implements usecase.CustomerUsecase.
func (a *Application) RegisterCustomer(ctx context.Context, register command.RegisterCustomer) (err error) {
	a.logger.Info("--> Customers.RegisterCustomer")
	defer func() { a.logger.WithError(err).Info("<-- Customers.RegisterCustomer") }()
	return a.application.RegisterCustomer(ctx, register)
}

// AuthorizeCustomer implements usecase.CustomerUsecase.
func (a *Application) AuthorizeCustomer(ctx context.Context, authorize command.AuthorizeCustomer) (err error) {
	a.logger.Info("--> Customers.RegisterCustomer")
	defer func() { a.logger.WithError(err).Info("<-- Customers.RegisterCustomer") }()
	return a.application.AuthorizeCustomer(ctx, authorize)
}

// EnableCustomer implements usecase.CustomerUsecase.
func (a *Application) EnableCustomer(ctx context.Context, enable command.EnableCustomer) (err error) {
	a.logger.Info("--> Customers.EnableCustomer")
	defer func() { a.logger.WithError(err).Info("<-- Customers.EnableCustomer") }()
	return a.application.EnableCustomer(ctx, enable)
}

// DisableCustomer implements usecase.CustomerUsecase.
func (a *Application) DisableCustomer(ctx context.Context, disable command.DisableCustomer) (err error) {
	a.logger.Info("--> Customers.DisableCustomer")
	defer func() { a.logger.WithError(err).Info("<-- Customers.DisableCustomer") }()
	return a.application.DisableCustomer(ctx, disable)
}

// GetCustomer implements usecase.CustomerUsecase.
func (a *Application) GetCustomer(ctx context.Context, query query.GetCustomer) (customer *aggregate.CustomerAgg, err error) {
	a.logger.Info("--> Customers.GetCustomer")
	defer func() { a.logger.WithError(err).Info("<-- Customers.GetCustomer") }()
	return a.application.GetCustomer(ctx, query)
}
