package monolith

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/config"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/logger"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/server"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/waiter"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/web"
	"github.com/fatih/color"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Monolith interface {
	Name() string
	Basename() string
	Logger() logger.Logger
	Config() *config.Config
	Database() *gorm.DB
	Gin() *gin.Engine
	RPC() *server.RPCServer
	JetStream() nats.JetStreamContext
	Run() error
}

var progressMessage = color.GreenString("==>")

type Option func(app *MonolithApplication)

type RunFunc func(basename string) error

func WithAppConfig(appCfg *config.Config) Option {
	return func(app *MonolithApplication) {
		app.appCfg = appCfg
	}
}

func WithRunFunc(runFunc RunFunc) Option {
	return func(app *MonolithApplication) {
		app.runFunc = runFunc
	}
}

func WithModules(modules ...Module) Option {
	return func(app *MonolithApplication) {
		app.modules = modules
	}
}
func WithDatabase(gormDB *gorm.DB) Option {
	return func(app *MonolithApplication) {
		app.gormDB = gormDB
	}
}

func WithGinEngine(ginEngine *gin.Engine) Option {
	return func(app *MonolithApplication) {
		app.gin = ginEngine
	}
}

func WithGrpcServer(grpcServer *server.RPCServer) Option {
	return func(app *MonolithApplication) {
		app.grpcServer = grpcServer
	}
}

func WithNatsConn(nc *nats.Conn) Option {
	return func(app *MonolithApplication) {
		app.nc = nc
	}
}

func WithJetStreamContext(js nats.JetStreamContext) Option {
	return func(app *MonolithApplication) {
		app.js = js
	}
}

type MonolithApplication struct {
	name              string
	basename          string
	logger            logger.Logger
	appCfg            *config.Config
	gormDB            *gorm.DB
	gin               *gin.Engine
	grpcServer        *server.RPCServer
	genericHttpServer *server.GenericHTTPServer
	nc                *nats.Conn
	js                nats.JetStreamContext
	modules           []Module
	waiter            waiter.Waiter
	runFunc           RunFunc
}

var _ Monolith = (*MonolithApplication)(nil)

func NewMonolithApplication(
	name string,
	basename string,
	appCfg *config.Config,
	logger logger.Logger,
	opts ...Option,
) *MonolithApplication {
	app := &MonolithApplication{
		name:     name,
		basename: basename,
		appCfg:   appCfg,
		logger:   logger,
	}

	for _, opt := range opts {
		opt(app)
	}

	return app
}

func (app *MonolithApplication) Name() string {
	return app.name
}

func (app *MonolithApplication) Basename() string {
	return app.basename
}

func (app *MonolithApplication) Logger() logger.Logger {
	return app.logger
}

func (app *MonolithApplication) Config() *config.Config {
	return app.appCfg
}

func (app *MonolithApplication) Database() *gorm.DB {
	return app.gormDB
}

func (app *MonolithApplication) Gin() *gin.Engine {
	return app.gin
}

func (app *MonolithApplication) RPC() *server.RPCServer {
	return app.grpcServer
}

func (app *MonolithApplication) JetStream() nats.JetStreamContext {
	return app.js
}

// FIXME:
func (app *MonolithApplication) PrepareRunModules() error {
	for _, module := range app.modules {
		app.logger.Infof("%v prepare module: %v", progressMessage, module.Name())
		if err := module.PrepareRun(context.Background(), app); err != nil {
			return err
		}
	}
	return nil
}

func (app *MonolithApplication) waitForWeb(ctx context.Context) error {
	eg, gCtx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		app.logger.Infof("web server started: %d", app.appCfg.Server.HTTP.Port)
		defer app.logger.Infof("web server shutdown")
		return app.genericHttpServer.ListenAndServe(gCtx)
	})

	return eg.Wait()
}

func (app *MonolithApplication) waitForRPC(ctx context.Context) error {
	eg, gCtx := errgroup.WithContext(ctx)

	address := fmt.Sprintf("%s:%d", app.appCfg.Server.GPPC.Host, app.appCfg.Server.GPPC.Port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		app.logger.Fatalf("failed to listen: %v", err)
		return err
	}

	eg.Go(func() error {
		app.logger.Infof("rpc server started: %d", app.appCfg.Server.GPPC.Port)
		defer app.logger.Infof("rpc server shutdown")
		err := app.RPC().GRPCServer().Serve(lis)
		if err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			app.logger.Fatalf("failed to serve grpc: %v", err)
			return err
		}

		return nil
	})

	eg.Go(func() error {
		<-gCtx.Done()
		app.logger.Infof("rpc server to be shutdown")
		stopped := make(chan struct{})
		go func() {
			app.RPC().GRPCServer().GracefulStop()
			close(stopped)
		}()
		timeout := time.NewTimer(5 * time.Second)
		select {
		case <-timeout.C:
			return fmt.Errorf("rpc server failed to stop gracefully")
		case <-stopped:
			return nil
		}
	})

	return eg.Wait()
}

func (app *MonolithApplication) waitForStream(ctx context.Context) error {
	closed := make(chan struct{})
	app.nc.SetClosedHandler(func(*nats.Conn) {
		close(closed)
	})
	eg, gCtx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		app.logger.Info("message stream started")
		defer app.logger.Info("message stream stopped")
		<-closed
		return nil
	})

	eg.Go(func() error {
		<-gCtx.Done()
		return app.nc.Drain()
	})
	return eg.Wait()
}

func (app *MonolithApplication) Run() error {
	app.logger.Info("started mallbots application")
	defer app.logger.Info("stopped mallbots application")
	app.printWorkingDir()

	if err := app.PrepareRunModules(); err != nil {
		app.logger.Errorf("%v %v\n", "prepare run modules: ", err)
		return err
	}

	staticFiles := static.EmbedFolder(web.WebUI, ".")
	app.gin.Use(static.Serve("/", staticFiles))
	// static := http.FileServer(http.FS(web.WebUI))
	// app.gin.GET("/", func(ctx *gin.Context) {
	// 	static.ServeHTTP(ctx.Writer, ctx.Request)
	// })
	// app.gin.GET("/swagger-ui/*filepath", func(ctx *gin.Context) {
	// 	static.ServeHTTP(ctx.Writer, ctx.Request)
	// })

	app.genericHttpServer = server.NewGenericHttpServer(app.gin, app.appCfg.Server)

	app.waiter = waiter.New(waiter.CatchSignals())

	app.waiter.Add(app.waitForWeb, app.waitForRPC, app.waitForStream)

	return app.waiter.Wait()
}

func (app *MonolithApplication) printWorkingDir() {
	wd, _ := os.Getwd()
	app.logger.Infof("%v workingDir: %s", progressMessage, wd)
}
