package cosec

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/cosec/internal"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/cosec/internal/handlers"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/cosec/internal/logging"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/cosec/internal/models"
	customerv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/customer/api/customer/v1"
	depotv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/depot/api/depot/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/broker/nats"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/monolith"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry/serdes"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/sagastore/gorm"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/sec"
	orderv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/api/order/v1"
	paymentv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/payment/api/payment/v1"
)

type Module struct{}

var _ monolith.Module = (*Module)(nil)

func NewModule() *Module { return &Module{} }

func (m *Module) PrepareRun(ctx context.Context, mono monolith.Monolith) (err error) {
	// setup Driven adapters
	reg := registry.New()
	if err = registrations(reg); err != nil {
		return err
	}
	if err = orderv1.Registrations(reg); err != nil {
		return err
	}
	if err = customerv1.Registrations(reg); err != nil {
		return err
	}
	if err = depotv1.Registrations(reg); err != nil {
		return err
	}
	if err = paymentv1.Registrations(reg); err != nil {
		return err
	}
	stream := nats.NewStream(mono.Config().Infrastructure.Nats.Stream, mono.JetStream(), mono.Logger())
	eventStream := am.NewEventStream(reg, stream)
	commandStream := am.NewCommandStream(reg, stream)
	replyStream := am.NewReplyStream(reg, stream)
	sagaStore := gorm.NewSagaStore("cosec.sagas", mono.Database(), reg)
	sagaRepo := sec.NewSagaRepository[*models.CreateOrderData](reg, sagaStore)

	// setup application
	orchestrator := logging.NewLogReplyHandlerAccess[*models.CreateOrderData](
		sec.NewOrchestrator[*models.CreateOrderData](internal.NewCreateOrderSaga(), sagaRepo, commandStream),
		"CreateOrderSaga", mono.Logger(),
	)
	integrationEventHandlers := logging.NewLogEventHandlerAccess[ddd.Event](
		handlers.NewIntegrationEventHandlers(orchestrator),
		"IntegrationEvents", mono.Logger(),
	)

	// setup Driver adapters
	if err = handlers.RegisterIntegrationEventHandlers(eventStream, integrationEventHandlers); err != nil {
		return err
	}
	if err = handlers.RegisterReplyHandlers(replyStream, orchestrator); err != nil {
		return err
	}

	return
}

func (m *Module) Name() string {
	return "cosec"
}

func registrations(reg registry.Registry) (err error) {
	serde := serdes.NewJSONSerde(reg)

	// Saga data
	if err = serde.RegisterKey(internal.CreateOrderSagaName, models.CreateOrderData{}); err != nil {
		return err
	}

	return nil
}
