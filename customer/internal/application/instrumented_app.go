package application

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/application/port/in/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/application/usecase"
	"github.com/prometheus/client_golang/prometheus"
)

type instrumentedApp struct {
	usecase.CustomerUsecase
	customersRegistered prometheus.Counter
}

var _ usecase.CustomerUsecase = (*instrumentedApp)(nil)

func NewInstrumentedApp(app usecase.CustomerUsecase, customersRegistered prometheus.Counter) usecase.CustomerUsecase {
	return instrumentedApp{
		CustomerUsecase:     app,
		customersRegistered: customersRegistered,
	}
}

func (a instrumentedApp) RegisterCustomer(ctx context.Context, register command.RegisterCustomer) error {
	err := a.CustomerUsecase.RegisterCustomer(ctx, register)
	if err != nil {
		return err
	}
	a.customersRegistered.Inc()
	return nil
}
