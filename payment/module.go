package payment

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/broker/nats"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/di"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/logger"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/monolith"
	outboxstoregorm "github.com/Chengxufeng1994/event-driven-arch-in-go/internal/outboxstore/gorm"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/tm"
	orderv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/api/order/v1"
	paymentv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/payment/api/payment/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/docs"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/application"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/application/handler"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/application/usecase"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/infrastructure/logging"
	persistencegorm "github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/infrastructure/persistence/gorm"
	v1 "github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/interface/grpc/v1"
	restv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/interface/rest/v1"
)

type Module struct{}

var _ monolith.Module = (*Module)(nil)

func NewModule() *Module { return &Module{} }

func (m *Module) PrepareRun(ctx context.Context, mono monolith.Monolith) error {
	container := di.New()
	// setup Driver adapters
	endpoint := fmt.Sprintf("%s:%d", mono.Config().Server.GPPC.Host, mono.Config().Server.GPPC.Port)
	container.AddSingleton("registry", func(c di.Container) (any, error) {
		reg := registry.New()
		if err := orderv1.Registrations(reg); err != nil {
			return nil, err
		}
		if err := paymentv1.Registrations(reg); err != nil {
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
			outboxstoregorm.NewOutboxStore("payments.outbox", c.Get("db").(*gorm.DB)),
		), nil
	})
	container.AddScoped("tx", func(c di.Container) (any, error) {
		db := c.Get("db").(*gorm.DB)
		return db.Begin(), nil
	})
	container.AddScoped("txStream", func(c di.Container) (any, error) {
		tx := c.Get("tx").(*gorm.DB)
		outboxStore := outboxstoregorm.NewOutboxStore("payments.outbox", tx)
		return am.RawMessageStreamWithMiddleware(
			c.Get("stream").(am.RawMessageStream),
			tm.NewOutboxStreamMiddleware(outboxStore),
		), nil
	})
	container.AddScoped("eventStream", func(c di.Container) (any, error) {
		return am.NewEventStream(c.Get("registry").(registry.Registry), c.Get("txStream").(am.RawMessageStream)), nil
	})
	container.AddScoped("replyStream", func(c di.Container) (any, error) {
		return am.NewReplyStream(c.Get("registry").(registry.Registry), c.Get("txStream").(am.RawMessageStream)), nil
	})
	container.AddScoped("inboxMiddleware", func(c di.Container) (any, error) {
		tx := c.Get("tx").(*gorm.DB)
		inboxStore := outboxstoregorm.NewInboxStore("payments.inbox", tx)
		return tm.NewInboxHandlerMiddleware(inboxStore), nil
	})
	container.AddScoped("invoices", func(c di.Container) (any, error) {
		return persistencegorm.NewGormInvoiceRepository(c.Get("tx").(*gorm.DB)), nil
	})
	container.AddScoped("payments", func(c di.Container) (any, error) {
		return persistencegorm.NewGormPaymentRepository(c.Get("tx").(*gorm.DB)), nil
	})

	// setup app
	container.AddScoped("app", func(c di.Container) (any, error) {
		return logging.NewLogApplicationAccess(
			application.NewPaymentApplication(
				c.Get("invoices").(repository.InvoiceRepository),
				c.Get("payments").(repository.PaymentRepository),
				c.Get("domainEventDispatcher").(ddd.EventDispatcher[ddd.Event]),
			),
			c.Get("logger").(logger.Logger)), nil
	})
	container.AddScoped("domainEventHandlers", func(c di.Container) (any, error) {
		return logging.NewLogEventHandlerAccess[ddd.Event](
			handler.NewDomainEventHandlers(c.Get("eventStream").(am.EventStream)),
			"DomainEvents", c.Get("logger").(logger.Logger),
		), nil
	})
	container.AddScoped("integrationEventHandlers", func(c di.Container) (any, error) {
		return logging.NewLogEventHandlerAccess[ddd.Event](
			handler.NewIntegrationEventHandlers(c.Get("app").(usecase.PaymentUseCase)),
			"IntegrationEvents", c.Get("logger").(logger.Logger),
		), nil
	})
	container.AddScoped("commandHandlers", func(c di.Container) (any, error) {
		return logging.NewLogCommandHandlerAccess(
			handler.NewCommandHandler(c.Get("app").(usecase.PaymentUseCase)),
			"Commands", c.Get("logger").(logger.Logger),
		), nil
	})

	// setup Driver adapters
	if err := v1.RegisterServerTx(container, mono.RPC().GRPCServer()); err != nil {
		return err
	}
	if err := restv1.RegisterGateway(ctx, mono.Gin(), endpoint); err != nil {
		return err
	}
	if err := docs.RegisterSwagger(mono.Gin()); err != nil {
		return err
	}
	if err := handler.RegisterIntegrationEventHandlersTx(container); err != nil {
		return err
	}
	handler.RegisterDomainEventHandlersTx(container)
	if err := handler.RegisterCommandHandlersTx(container); err != nil {
		return err
	}
	go m.startOutboxProcessor(ctx, container)

	return nil
}

func (m *Module) Name() string {
	return "payment"
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
		logger.WithError(err).Error("payments outbox processor encountered an error")
	}
}
