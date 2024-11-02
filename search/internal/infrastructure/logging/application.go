package logging

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/logger"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/application/port/in/query"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/application/usecase"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/domain"
)

type Application struct {
	usecase.SearchUseCase
	logger logger.Logger
}

var _ usecase.SearchUseCase = (*Application)(nil)

func NewLogApplicationAccess(app usecase.SearchUseCase, logger logger.Logger) Application {
	return Application{
		SearchUseCase: app,
		logger:        logger,
	}
}

// GetOrder implements usecase.SearchUseCase.
func (a Application) GetOrder(ctx context.Context, get query.GetOrder) (order *domain.Order, err error) {
	a.logger.Info("--> Search.GetOrder")
	defer func() { a.logger.WithError(err).Info("<-- Search.GetOrder") }()
	return a.SearchUseCase.GetOrder(ctx, get)
}

// SearchOrders implements usecase.SearchUseCase.
func (a Application) SearchOrders(ctx context.Context, search query.SearchOrders) (orders []*domain.Order, err error) {
	a.logger.Info("--> Search.SearchOrders")
	defer func() { a.logger.WithError(err).Info("<-- Search.SearchOrders") }()
	return a.SearchUseCase.SearchOrders(ctx, search)
}
