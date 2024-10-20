package usecase

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/application/port/int/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/application/port/int/query"
)

type PaymentUseCase interface {
	command.Commands
	query.Queries
}
