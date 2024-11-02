package main

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/config"
	pkggorm "github.com/Chengxufeng1994/event-driven-arch-in-go/internal/gorm"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/jetstream"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/logger"
	logging "github.com/Chengxufeng1994/event-driven-arch-in-go/internal/logger"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/monolith"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/server"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/notification"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"github.com/spf13/viper"
)

func NewApp() *monolith.MonolithApplication {
	logger := logging.ContextUnavailable()

	cfg, err := config.LoadConfig("")
	if err != nil {
		logger.Errorf("failed to load config: %v", err)
		return nil
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(_ fsnotify.Event) {
		logger.Info("reloading config")
		if err := viper.Unmarshal(&cfg); err != nil {
			logger.Errorf("failed to unmarshal config: %v", err)
		}
	})

	gormDB, err := pkggorm.NewGormDB(cfg.Infrastructure)
	if err != nil {
		logger.Errorf("failed to connect to database: %v", err)
		return nil
	}

	ginEngine := initGinEngine()

	grpcServer := initGrpcServer(logger, cfg.Server)

	nc, err := nats.Connect(cfg.Infrastructure.Nats.URL)
	if err != nil {
		logger.Errorf("failed to connect to nats: %v", err)
		return nil
	}
	js, err := jetstream.NewJetStream(cfg.Infrastructure, nc)
	if err != nil {
		return nil
	}

	basketModule := basket.NewModule()
	customerModule := customer.NewModule()
	depotModule := depot.NewModule()
	notificationModule := notification.NewModule()
	orderModule := ordering.NewModule()
	paymentModule := payment.NewModule()
	searchModule := search.NewModule()
	storeModule := store.NewModule()

	monolithApplication := monolith.NewMonolithApplication(
		"MALL BOTS",
		"mallbots-monolith-application",
		cfg,
		logger,
		monolith.WithDatabase(gormDB),
		monolith.WithGinEngine(ginEngine),
		monolith.WithGrpcServer(grpcServer),
		monolith.WithNatsConn(nc),
		monolith.WithJetStreamContext(js),
		monolith.WithModules(
			basketModule,
			customerModule,
			depotModule,
			notificationModule,
			orderModule,
			paymentModule,
			searchModule,
			storeModule),
	)

	return monolithApplication
}

func initGinEngine() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(cors.Default())
	router.Use(requestid.New())
	router.Use(gin.Logger())
	return router
}

func initGrpcServer(logger logger.Logger, cfg *config.Server) *server.RpcServer {
	return server.NewGrpcServer(logger, cfg)
}
