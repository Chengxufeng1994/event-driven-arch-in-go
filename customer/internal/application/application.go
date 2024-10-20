package application

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/application/port/in/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/application/port/in/query"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/application/usecase"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/domain/repository"
)

type (
	CustomerApplication struct {
		appCommands
		appQueries
	}

	appCommands struct {
		command.RegisterCustomerHandler
		command.AuthorizeCustomerHandler
		command.EnableCustomerHandler
		command.DisableCustomerHandler
	}

	appQueries struct {
		query.GetCustomerHandler
	}
)

var _ usecase.CustomerUsecase = (*CustomerApplication)(nil)

func NewCustomerApplication(
	customerRepository repository.CustomerRepository,
) *CustomerApplication {
	return &CustomerApplication{
		appCommands: appCommands{
			RegisterCustomerHandler:  command.NewRegisterCustomerHandler(customerRepository),
			AuthorizeCustomerHandler: command.NewAuthorizeCustomerHandler(customerRepository),
			EnableCustomerHandler:    command.NewEnableCustomerHandler(customerRepository),
			DisableCustomerHandler:   command.NewDisableCustomerHandler(customerRepository),
		},
		appQueries: appQueries{
			GetCustomerHandler: query.NewGetCustomerHandler(customerRepository),
		},
	}
}
