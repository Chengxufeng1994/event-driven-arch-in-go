package usecase

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application/port/in/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application/port/in/query"
)

type OrderUseCase interface {
	command.Commands
	query.Queries
}
