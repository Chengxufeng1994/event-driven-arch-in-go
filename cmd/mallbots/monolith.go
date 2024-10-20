package main

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/config"
	pkggorm "github.com/Chengxufeng1994/event-driven-arch-in-go/internal/gorm"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/logger"
	logging "github.com/Chengxufeng1994/event-driven-arch-in-go/internal/logger"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/monolith"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/server"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/notification"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store"
	"github.com/gin-gonic/gin"
)

func NewApp() *monolith.MonolithApplication {
	logger := logging.ContextUnavailable()

	cfg, err := config.LoadConfig("")
	if err != nil {
		logger.Errorf("failed to load config: %v", err)
		return nil
	}

	gormDB, err := pkggorm.NewGormDB(cfg.Infrastructure)
	if err != nil {
		logger.Errorf("failed to connect to database: %v", err)
		return nil
	}

	ginEngine := initGinEngine()

	grpcServer := initGrpcServer(logger, cfg.Server)
	// 	&ordering.Module{},
	// 	&payments.Module{},

	basketModule := basket.NewModule()
	customerModule := customer.NewModule()
	depotModule := depot.NewModule()
	notificationModule := notification.NewModule()
	storeModule := store.NewModule()

	monolithApplication := monolith.NewMonolithApplication(
		"MALL BOTS",
		"mallbots-monolith-application",
		cfg,
		logger,
		monolith.WithDatabase(gormDB),
		monolith.WithGinEngine(ginEngine),
		monolith.WithGrpcServer(grpcServer),
		monolith.WithModules(
			basketModule,
			customerModule,
			depotModule,
			notificationModule,
			storeModule),
	)

	return monolithApplication
}

func initGinEngine() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	return router
}

func initRouter(g *gin.Engine) {
}

func initGrpcServer(logger logger.Logger, cfg *config.Server) *server.RpcServer {
	return server.NewGrpcServer(logger, cfg)
}
