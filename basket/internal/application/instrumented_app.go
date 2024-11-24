package application

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/port/in/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/usecase"
	"github.com/prometheus/client_golang/prometheus"
)

type instrumentedApp struct {
	usecase.BasketUseCase
	basketsStarted    prometheus.Counter
	basketsCheckedOut prometheus.Counter
	basketsCanceled   prometheus.Counter
}

var _ usecase.BasketUseCase = (*instrumentedApp)(nil)

func NewInstrumentedApp(app usecase.BasketUseCase, basketsStarted, basketsCheckedOut, baksetsCanceled prometheus.Counter) usecase.BasketUseCase {
	return instrumentedApp{
		BasketUseCase:     app,
		basketsStarted:    basketsStarted,
		basketsCheckedOut: basketsCheckedOut,
		basketsCanceled:   baksetsCanceled,
	}
}

func (a instrumentedApp) StartBasket(ctx context.Context, start command.StartBasket) error {
	err := a.BasketUseCase.StartBasket(ctx, start)
	if err != nil {
		return err
	}
	a.basketsStarted.Inc()
	return nil
}

func (a instrumentedApp) CheckoutBasket(ctx context.Context, checkout command.CheckoutBasket) error {
	err := a.BasketUseCase.CheckoutBasket(ctx, checkout)
	if err != nil {
		return err
	}
	a.basketsCheckedOut.Inc()
	return nil
}

func (a instrumentedApp) CancelBasket(ctx context.Context, cancel command.CancelBasket) error {
	err := a.BasketUseCase.CancelBasket(ctx, cancel)
	if err != nil {
		return err
	}
	a.basketsCanceled.Inc()
	return nil
}
