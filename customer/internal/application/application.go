package application

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/application/port/in/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/application/port/in/event"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/application/port/in/query"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/application/usecase"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
)

type (
	CustomerApplication struct {
		appCommands
		appQueries
		appEvents
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

	appEvents struct {
		event.IgnoreUnimplementedDomainEventHandler
	}
)

var _ usecase.CustomerUsecase = (*CustomerApplication)(nil)

func NewCustomerApplication(
	customerRepository repository.CustomerRepository,
	domainEventDispatcher ddd.EventDispatcherIntf,
) *CustomerApplication {
	return &CustomerApplication{
		appCommands: appCommands{
			RegisterCustomerHandler:  command.NewRegisterCustomerHandler(customerRepository, domainEventDispatcher),
			AuthorizeCustomerHandler: command.NewAuthorizeCustomerHandler(customerRepository, domainEventDispatcher),
			EnableCustomerHandler:    command.NewEnableCustomerHandler(customerRepository, domainEventDispatcher),
			DisableCustomerHandler:   command.NewDisableCustomerHandler(customerRepository, domainEventDispatcher),
		},
		appQueries: appQueries{
			GetCustomerHandler: query.NewGetCustomerHandler(customerRepository),
		},
		appEvents: appEvents{
			IgnoreUnimplementedDomainEventHandler: event.NewIgnoreUnimplementedDomainEventHandler(),
		},
	}
}
