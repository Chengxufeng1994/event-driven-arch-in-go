package logging

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/application/usecase"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/logger"
)

type Application struct {
	usecase.ShoppingListUseCase
	logger.Logger
}

var _ usecase.ShoppingListUseCase = (*Application)(nil)

func NewLoggingApplicationAccess(shoppingListApplication usecase.ShoppingListUseCase, logger logger.Logger) *Application {
	return &Application{
		ShoppingListUseCase: shoppingListApplication,
		Logger:              logger,
	}
}
