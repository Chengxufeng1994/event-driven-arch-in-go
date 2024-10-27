package usecase

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/application/port/in/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/application/port/in/event"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/application/port/in/query"
)

type CustomerUsecase interface {
	command.Commands
	query.Queries
	event.DomainEventHandlers
}
