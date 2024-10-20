package usecase

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/port/in/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/port/in/query"
)

type BasketUseCase interface {
	command.Commands
	query.Queries
}
