package customer

import (
	"context"

	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"

	customerv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/customer/api/customer/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/docs"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/application"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/application/handler"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/application/usecase"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/infrastructure/logging"
	persistencegorm "github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/infrastructure/persistence/gorm"
	grpcv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/interface/grpc/v1"
	restv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/interface/rest/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/broker/nats"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/di"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/logger"
	outboxstoregorm "github.com/Chengxufeng1994/event-driven-arch-in-go/internal/outboxstore/gorm"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/system"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/tm"
)

type Module struct{}

var _ system.Module = (*Module)(nil)

func NewModule() *Module { return &Module{} }

func Root(ctx context.Context, svc system.Service) error {
	container := di.New()

	// setup Driver adapters
	container.AddSingleton("registry", func(c di.Container) (any, error) {
		reg := registry.New()
		if err := customerv1.Registrations(reg); err != nil {
			return nil, err
		}
		return reg, nil
	})
	container.AddSingleton("logger", func(c di.Container) (any, error) {
		return svc.Logger(), nil
	})
	container.AddSingleton("stream", func(c di.Container) (any, error) {
		return nats.NewStream(svc.Config().Infrastructure.Nats.Stream, svc.JetStream(), svc.Logger()), nil
	})
	container.AddSingleton("domainEventDispatcher", func(c di.Container) (any, error) {
		return ddd.NewEventDispatcher[ddd.AggregateEvent](), nil
	})
	container.AddSingleton("db", func(c di.Container) (any, error) {
		return svc.Database(), nil
	})
	container.AddSingleton("outboxProcessor", func(c di.Container) (any, error) {
		return tm.NewOutboxProcessor(
			c.Get("stream").(am.RawMessageStream),
			outboxstoregorm.NewOutboxStore("customers.outbox", c.Get("db").(*gorm.DB)),
		), nil
	})
	container.AddScoped("tx", func(c di.Container) (any, error) {
		db := c.Get("db").(*gorm.DB)
		return db.Begin(), nil
	})
	container.AddScoped("customers", func(c di.Container) (any, error) {
		return persistencegorm.NewGormCustomerRepository(c.Get("tx").(*gorm.DB)), nil
	})
	container.AddScoped("txStream", func(c di.Container) (any, error) {
		tx := c.Get("db").(*gorm.DB)
		outboxStore := outboxstoregorm.NewOutboxStore("customers.outbox", tx)
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
		inboxStore := outboxstoregorm.NewInboxStore("customers.inbox", tx)
		return tm.NewInboxHandlerMiddleware(inboxStore), nil
	})

	// setup application
	container.AddScoped("app", func(c di.Container) (any, error) {
		return logging.NewLogApplicationAccess(
			application.NewCustomerApplication(
				c.Get("customers").(repository.CustomerRepository),
				c.Get("domainEventDispatcher").(ddd.EventDispatcher[ddd.AggregateEvent]),
			),
			c.Get("logger").(logger.Logger),
		), nil
	})
	container.AddScoped("domainEventHandlers", func(c di.Container) (any, error) {
		return logging.NewLogEventHandlerAccess[ddd.AggregateEvent](
			handler.NewDomainEventHandler(c.Get("eventStream").(am.EventStream)),
			"DomainEvents", c.Get("logger").(logger.Logger),
		), nil
	})
	container.AddScoped("commandHandlers", func(c di.Container) (any, error) {
		return logging.NewLogCommandHandlerAccess[ddd.Command](
			handler.NewCommandHandler(c.Get("app").(usecase.CustomerUsecase)),
			"Commands", c.Get("logger").(logger.Logger),
		), nil
	})

	// setup Driver adapters
	if err := grpcv1.RegisterServerTx(container, svc.RPC().GRPCServer()); err != nil {
		return err
	}
	if err := restv1.RegisterGateway(ctx, svc.Gin(), svc.Config().Server.GPPC.Address()); err != nil {
		return err
	}
	if err := docs.RegisterSwagger(svc.Gin()); err != nil {
		return err
	}
	handler.RegisterDomainEventHandlersTx(container)
	if err := handler.RegisterCommandHandlersTx(container); err != nil {
		return err
	}

	go startOutboxProcessor(ctx, container)

	return nil
}

func (m *Module) Startup(ctx context.Context, svc system.Service) error {
	return Root(ctx, svc)
}

func (m *Module) Name() string {
	return "customers"
}

func startOutboxProcessor(ctx context.Context, container di.Container) {
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
		logger.WithError(err).Error("customers outbox processor encountered an error")
	}
}
