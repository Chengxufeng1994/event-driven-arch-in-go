package depot

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	depotv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/depot/api/depot/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/docs"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/application"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/application/handler"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/application/usecase"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/repository"
	infragrpc "github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/infastructure/client/grpc"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/infastructure/logging"
	persistencegorm "github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/infastructure/persistence/gorm"
	grpcv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/interfaces/grpc/v1"
	restv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/interfaces/rest/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/broker/nats"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/di"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/logger"
	outboxstoregorm "github.com/Chengxufeng1994/event-driven-arch-in-go/internal/outboxstore/gorm"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/system"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/tm"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
)

type Module struct{}

var _ system.Module = (*Module)(nil)

func NewModule() *Module { return &Module{} }

func (m *Module) Startup(ctx context.Context, mono system.Service) error {
	container := di.New()

	// setup Driver adapters
	endpoint := fmt.Sprintf("%s:%d", mono.Config().Server.GPPC.Host, mono.Config().Server.GPPC.Port)
	container.AddSingleton("registry", func(c di.Container) (any, error) {
		reg := registry.New()
		if err := storev1.Registrations(reg); err != nil {
			return nil, err
		}
		if err := depotv1.Registrations(reg); err != nil {
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
		return ddd.NewEventDispatcher[ddd.AggregateEvent](), nil
	})
	container.AddSingleton("db", func(c di.Container) (any, error) {
		return mono.Database(), nil
	})
	container.AddSingleton("conn", func(c di.Container) (any, error) {
		return infragrpc.Dial(ctx, endpoint)
	})
	container.AddSingleton("outboxProcessor", func(c di.Container) (any, error) {
		return tm.NewOutboxProcessor(
			c.Get("stream").(am.RawMessageStream),
			outboxstoregorm.NewOutboxStore("depot.outbox", c.Get("db").(*gorm.DB)),
		), nil
	})
	container.AddScoped("tx", func(c di.Container) (any, error) {
		db := c.Get("db").(*gorm.DB)
		return db.Begin(), nil
	})
	container.AddScoped("txStream", func(c di.Container) (any, error) {
		tx := c.Get("tx").(*gorm.DB)
		outboxStore := outboxstoregorm.NewOutboxStore("depot.outbox", tx)
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
		inboxStore := outboxstoregorm.NewInboxStore("depot.inbox", tx)
		return tm.NewInboxHandlerMiddleware(inboxStore), nil
	})
	container.AddScoped("shoppingLists", func(c di.Container) (any, error) {
		tx := c.Get("tx").(*gorm.DB)
		return persistencegorm.NewGormShoppingListRepository(tx), nil
	})
	container.AddScoped("stores", func(c di.Container) (any, error) {
		return persistencegorm.NewGormStoreCacheRepository(
				c.Get("tx").(*gorm.DB),
				infragrpc.NewGrpcStoreClient(c.Get("conn").(*grpc.ClientConn))),
			nil
	})
	container.AddScoped("products", func(c di.Container) (any, error) {
		return persistencegorm.NewGormProductCacheRepository(
				c.Get("tx").(*gorm.DB),
				infragrpc.NewGrpcProductClient(c.Get("conn").(*grpc.ClientConn))),
			nil
	})

	// setup application
	container.AddScoped("app", func(c di.Container) (any, error) {
		return logging.NewLogApplicationAccess(
			application.NewShoppingListApplication(
				c.Get("shoppingLists").(repository.ShoppingListRepository),
				c.Get("stores").(repository.StoreCacheRepository),
				c.Get("products").(repository.ProductCacheRepository),
				c.Get("domainEventDispatcher").(ddd.EventDispatcher[ddd.AggregateEvent]),
			),
			c.Get("logger").(logger.Logger)), nil
	})
	container.AddScoped("domainEventHandlers", func(c di.Container) (any, error) {
		return logging.NewLogEventHandlerAccess[ddd.AggregateEvent](
			handler.NewDomainEventHandlers(c.Get("eventStream").(am.EventStream)),
			"DomainEvents", c.Get("logger").(logger.Logger),
		), nil
	})
	container.AddScoped("integrationEventHandlers", func(c di.Container) (any, error) {
		return logging.NewLogEventHandlerAccess[ddd.Event](
			handler.NewIntegrationEventHandlers(
				c.Get("products").(repository.ProductCacheRepository),
				c.Get("stores").(repository.StoreCacheRepository),
			),
			"IntegrationEvents", c.Get("logger").(logger.Logger),
		), nil
	})
	container.AddScoped("commandHandlers", func(c di.Container) (any, error) {
		return logging.NewLogCommandHandlerAccess[ddd.Command](
			handler.NewCommandHandlers(c.Get("app").(usecase.ShoppingListUseCase)),
			"Commands", c.Get("logger").(logger.Logger),
		), nil
	})

	// setup Driver adapters
	if err := grpcv1.RegisterServerTx(container, mono.RPC().GRPCServer()); err != nil {
		return err
	}
	if err := restv1.RegisterGateway(ctx, mono.Gin(), endpoint); err != nil {
		return err
	}
	if err := docs.RegisterSwagger(mono.Gin()); err != nil {
		return err
	}
	handler.RegisterDomainEventHandlersTx(container)
	if err := handler.RegisterIntegrationEventHandlersTx(container); err != nil {
		return err
	}
	if err := handler.RegisterCommandHandlersTx(container); err != nil {
		return err
	}

	go m.startOutboxProcessor(ctx, container)

	return nil
}

func (m *Module) Name() string {
	return "depot"
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
		logger.WithError(err).Error("depot outbox processor encountered an error")
	}
}
