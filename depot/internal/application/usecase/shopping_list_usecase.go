package usecase

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/application/port/in/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/application/port/in/query"
)

type ShoppingListUseCase interface {
	command.Commands
	query.Queries
}
