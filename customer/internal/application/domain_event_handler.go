package application

import "github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/application/port/in/event"

type (
	CustomerDomainEventHandler struct {
		appDomainEvents
	}

	appDomainEvents struct {
		event.IgnoreUnimplementedDomainEventHandler
	}
)

var _ event.DomainEventHandlers = (*CustomerDomainEventHandler)(nil)

func NewCustomerDomainEventHandler() *CustomerDomainEventHandler {
	return &CustomerDomainEventHandler{
		appDomainEvents: appDomainEvents{
			IgnoreUnimplementedDomainEventHandler: event.NewIgnoreUnimplementedDomainEventHandler(),
		},
	}
}
