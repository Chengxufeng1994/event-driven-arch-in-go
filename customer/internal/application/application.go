package application

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/application/port/in/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/application/port/in/query"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/application/usecase"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
)

type (
	CustomerApplication struct {
		appCommands
		appQueries
	}

	appCommands struct {
		command.RegisterCustomerHandler
		command.ChangeSmsNumberHandler
		command.AuthorizeCustomerHandler
		command.EnableCustomerHandler
		command.DisableCustomerHandler
	}

	appQueries struct {
		query.GetCustomerHandler
	}
)

var _ usecase.CustomerUsecase = (*CustomerApplication)(nil)

func New(
	customerRepository repository.CustomerRepository,
	domainEventPublisher ddd.EventPublisher[ddd.AggregateEvent],
) *CustomerApplication {
	return &CustomerApplication{
		appCommands: appCommands{
			RegisterCustomerHandler:  command.NewRegisterCustomerHandler(customerRepository, domainEventPublisher),
			ChangeSmsNumberHandler:   command.NewChangeSmsNumberHandler(customerRepository, domainEventPublisher),
			AuthorizeCustomerHandler: command.NewAuthorizeCustomerHandler(customerRepository, domainEventPublisher),
			EnableCustomerHandler:    command.NewEnableCustomerHandler(customerRepository, domainEventPublisher),
			DisableCustomerHandler:   command.NewDisableCustomerHandler(customerRepository, domainEventPublisher),
		},
		appQueries: appQueries{
			GetCustomerHandler: query.NewGetCustomerHandler(customerRepository),
		},
	}
}
