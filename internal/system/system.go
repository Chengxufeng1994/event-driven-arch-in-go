package system

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/config"
	pkggorm "github.com/Chengxufeng1994/event-driven-arch-in-go/internal/gorm"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/jetstream"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/logger"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/server"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/waiter"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/migrations"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"github.com/pressly/goose/v3"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type RunFunc func(basename string) error

type System struct {
	name              string
	basename          string
	logger            logger.Logger
	cfg               *config.Config
	gormDB            *gorm.DB
	gin               *gin.Engine
	grpcServer        *server.RPCServer
	genericHttpServer *server.GenericHTTPServer
	nc                *nats.Conn
	js                nats.JetStreamContext
	waiter            waiter.Waiter
}

var _ Service = (*System)(nil)

func NewSystem(name, basename string, cfg *config.Config, logger logger.Logger) (s *System, err error) {
	s = &System{
		name:     name,
		basename: basename,
		cfg:      cfg,
		logger:   logger,
	}

	s.gormDB, err = pkggorm.NewGormDB(cfg.Infrastructure)
	if err != nil {
		return nil, errors.New("failed to connect to database")
	}

	if err := migrateDB(pkggorm.SqlDB); err != nil {
		return nil, errors.New("failed to migrate to database")
	}

	s.nc, err = nats.Connect(cfg.Infrastructure.Nats.URL)
	if err != nil {
		return nil, errors.New("failed to connect to nats")
	}
	s.js, err = jetstream.NewJetStream(cfg.Infrastructure, s.nc)
	if err != nil {
		return nil, err
	}

	s.initGinEngine()
	s.initGrpcServer(logger, cfg.Server)
	s.initWaiter()

	return s, nil
}

func (s System) Name() string                     { return s.name }
func (s System) Basename() string                 { return s.basename }
func (s System) Config() *config.Config           { return s.cfg }
func (s System) Database() *gorm.DB               { return s.gormDB }
func (s System) Gin() *gin.Engine                 { return s.gin }
func (s System) JetStream() nats.JetStreamContext { return s.js }
func (s System) Logger() logger.Logger            { return s.logger }
func (s System) RPC() *server.RPCServer           { return s.grpcServer }
func (s System) Waiter() waiter.Waiter            { return s.waiter }

func (s *System) initGinEngine() {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(cors.Default())
	router.Use(requestid.New())
	router.Use(gin.Logger())
	s.gin = router
}

func (s *System) initGrpcServer(logger logger.Logger, cfg *config.Server) {
	s.grpcServer = server.NewGrpcServer(logger, cfg)
}
func (s *System) initWaiter() {
	s.waiter = waiter.New(waiter.CatchSignals())
}

func migrateDB(db *sql.DB) error {
	goose.SetBaseFS(migrations.FS)
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}
	if err := goose.Up(db, "."); err != nil {
		return err
	}
	return nil
}

func (s *System) WaitForWeb(ctx context.Context) error {
	s.genericHttpServer = server.NewGenericHttpServer(s.gin, s.cfg.Server)

	eg, gCtx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		s.logger.Infof("web server started: %d", s.cfg.Server.HTTP.Port)
		defer s.logger.Infof("web server shutdown")
		return s.genericHttpServer.ListenAndServe(gCtx)
	})

	return eg.Wait()
}

func (s *System) WaitForRPC(ctx context.Context) error {
	eg, gCtx := errgroup.WithContext(ctx)

	address := fmt.Sprintf("%s:%d", s.cfg.Server.GPPC.Host, s.cfg.Server.GPPC.Port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		s.logger.Fatalf("failed to listen: %v", err)
		return err
	}

	eg.Go(func() error {
		s.logger.Infof("rpc server started: %d", s.cfg.Server.GPPC.Port)
		defer s.logger.Infof("rpc server shutdown")
		err := s.RPC().GRPCServer().Serve(lis)
		if err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			s.logger.Fatalf("failed to serve grpc: %v", err)
			return err
		}

		return nil
	})

	eg.Go(func() error {
		<-gCtx.Done()
		s.logger.Infof("rpc server to be shutdown")
		stopped := make(chan struct{})
		go func() {
			s.RPC().GRPCServer().GracefulStop()
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

func (app *System) WaitForStream(ctx context.Context) error {
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
