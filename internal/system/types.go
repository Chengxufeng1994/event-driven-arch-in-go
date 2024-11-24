package system

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/config"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/logger"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/rpc"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/waiter"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"gorm.io/gorm"
)

type Service interface {
	Name() string
	Basename() string
	Logger() logger.Logger
	Config() *config.Config
	Database() *gorm.DB
	Gin() *gin.Engine
	RPC() *rpc.RPCServer
	JetStream() nats.JetStreamContext
	Waiter() waiter.Waiter
}

type Module interface {
	Startup(context.Context, Service) error
}
