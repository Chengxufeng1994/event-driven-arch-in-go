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
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/di"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/logger"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/monolith"
	outboxstoregorm "github.com/Chengxufeng1994/event-driven-arch-in-go/internal/outboxstore/gorm"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry/serdes"
	sagastoregorm "github.com/Chengxufeng1994/event-driven-arch-in-go/internal/sagastore/gorm"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/sec"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/tm"
	orderv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/api/order/v1"
	paymentv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/payment/api/payment/v1"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

type Module struct{}

var _ monolith.Module = (*Module)(nil)

func NewModule() *Module { return &Module{} }

func (m *Module) PrepareRun(ctx context.Context, mono monolith.Monolith) (err error) {
	container := di.New()
	// setup Driver adapters
	container.AddSingleton("registry", func(c di.Container) (any, error) {
		reg := registry.New()
		if err = registrations(reg); err != nil {
			return nil, err
		}
		if err = orderv1.Registrations(reg); err != nil {
			return nil, err
		}
		if err = customerv1.Registrations(reg); err != nil {
			return nil, err
		}
		if err = depotv1.Registrations(reg); err != nil {
			return nil, err
		}
		if err = paymentv1.Registrations(reg); err != nil {
			return nil, err
		}
		return reg, nil
	})
	container.AddSingleton("logger", func(c di.Container) (any, error) {
		return mono.Logger(), nil
	})
	container.AddSingleton("stream", func(c di.Container) (any, error) {
		return nats.NewStream(mono.Config().Infrastructure.Nats.Stream, mono.JetStream(), mono.Logger()), nil
	})
	container.AddSingleton("domainEventDispatcher", func(c di.Container) (any, error) {
		return ddd.NewEventDispatcher[ddd.Event](), nil
	})
	container.AddSingleton("db", func(c di.Container) (any, error) {
		return mono.Database(), nil
	})
	container.AddSingleton("outboxProcessor", func(c di.Container) (any, error) {
		return tm.NewOutboxProcessor(
			c.Get("stream").(am.RawMessageStream),
			outboxstoregorm.NewOutboxStore("cosec.outbox", c.Get("db").(*gorm.DB)),
		), nil
	})
	container.AddScoped("tx", func(c di.Container) (any, error) {
		db := c.Get("db").(*gorm.DB)
		return db.Begin(), nil
	})
	container.AddScoped("txStream", func(c di.Container) (any, error) {
		tx := c.Get("tx").(*gorm.DB)
		outboxStore := outboxstoregorm.NewOutboxStore("cosec.outbox", tx)
		return am.RawMessageStreamWithMiddleware(
			c.Get("stream").(am.RawMessageStream),
			tm.NewOutboxStreamMiddleware(outboxStore),
		), nil
	})
	container.AddScoped("eventStream", func(c di.Container) (any, error) {
		return am.NewEventStream(c.Get("registry").(registry.Registry), c.Get("txStream").(am.RawMessageStream)), nil
	})
	container.AddScoped("commandStream", func(c di.Container) (any, error) {
		return am.NewCommandStream(c.Get("registry").(registry.Registry), c.Get("txStream").(am.RawMessageStream)), nil
	})
	container.AddScoped("replyStream", func(c di.Container) (any, error) {
		return am.NewReplyStream(c.Get("registry").(registry.Registry), c.Get("txStream").(am.RawMessageStream)), nil
	})
	container.AddScoped("inboxMiddleware", func(c di.Container) (any, error) {
		tx := c.Get("tx").(*gorm.DB)
		inboxStore := outboxstoregorm.NewInboxStore("cosec.inbox", tx)
		return tm.NewInboxHandlerMiddleware(inboxStore), nil
	})
	container.AddScoped("sagaRepository", func(c di.Container) (any, error) {
		tx := c.Get("tx").(*gorm.DB)
		reg := c.Get("registry").(registry.Registry)
		return sec.NewSagaRepository[*models.CreateOrderData](
				reg,
				sagastoregorm.NewSagaStore("cosec.sagas", tx, reg),
			),
			nil
	})
	container.AddSingleton("saga", func(c di.Container) (any, error) {
		return internal.NewCreateOrderSaga(), nil
	})

	// setup application
	container.AddScoped("orchestrator", func(c di.Container) (any, error) {
		return logging.NewLogReplyHandlerAccess[*models.CreateOrderData](
			sec.NewOrchestrator[*models.CreateOrderData](
				c.Get("saga").(sec.Saga[*models.CreateOrderData]),
				c.Get("sagaRepository").(sec.SagaRepository[*models.CreateOrderData]),
				c.Get("commandStream").(am.CommandStream),
			),
			"CreateOrderSaga", mono.Logger(),
		), nil
	})
	container.AddScoped("integrationEventHandlers", func(c di.Container) (any, error) {
		return logging.NewLogEventHandlerAccess[ddd.Event](
			handlers.NewIntegrationEventHandlers(
				c.Get("orchestrator").(sec.Orchestrator[*models.CreateOrderData]),
			),
			"IntegrationEvents", mono.Logger(),
		), nil
	})

	// setup Driver adapters
	if err = handlers.RegisterIntegrationEventHandlersTx(container); err != nil {
		return err
	}
	if err = handlers.RegisterReplyHandlersTx(container); err != nil {
		return err
	}
	go m.startOutboxProcessor(ctx, container)

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

func (m *Module) startOutboxProcessor(ctx context.Context, container di.Container) {
	outboxProcessor := container.Get("outboxProcessor").(tm.OutboxProcessor)
	logger := container.Get("logger").(logger.Logger)

	eg := errgroup.Group{}
	eg.Go(func() error {
		return outboxProcessor.Start(ctx)
	})

	eg.Go(func() error {
		<-ctx.Done()
		return nil
	})

	if err := eg.Wait(); err != nil {
		logger.WithError(err).Error("cosec outbox processor encountered an error")
	}
}
