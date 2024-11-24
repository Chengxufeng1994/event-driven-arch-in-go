package system

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"net"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"github.com/pressly/goose/v3"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/config"
	pkggorm "github.com/Chengxufeng1994/event-driven-arch-in-go/internal/gorm"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/jetstream"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/logger"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/rpc"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/server"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/waiter"
)

type RunFunc func(basename string) error

type System struct {
	name              string
	basename          string
	logger            logger.Logger
	cfg               *config.Config
	db                *gorm.DB
	nc                *nats.Conn
	js                nats.JetStreamContext
	gin               *gin.Engine
	genericHTTPServer *server.GenericHTTPServer
	grpcServer        *rpc.RPCServer
	waiter            waiter.Waiter
	tp                *sdktrace.TracerProvider
}

var _ Service = (*System)(nil)

func NewSystem(name, basename string, cfg *config.Config, logger logger.Logger) (s *System, err error) {
	s = &System{
		name:     name,
		basename: basename,
		cfg:      cfg,
		logger:   logger,
	}

	s.initWaiter()

	s.db, err = pkggorm.NewGormDB(cfg.Infrastructure)
	if err != nil {
		return nil, errors.New("failed to connect to database")
	}

	s.nc, err = nats.Connect(cfg.Infrastructure.Nats.URL)
	if err != nil {
		return nil, errors.New("failed to connect to nats")
	}
	s.js, err = jetstream.NewJetStream(cfg.Infrastructure, s.nc)
	if err != nil {
		return nil, errors.New("failed to connect to jetstream")
	}

	if err = s.initOpenTelemetry(); err != nil {
		return nil, err
	}

	s.initGinEngine()
	s.initGrpcServer(cfg.Server)

	return s, nil
}

func (s System) Name() string                     { return s.name }
func (s System) Basename() string                 { return s.basename }
func (s System) Config() *config.Config           { return s.cfg }
func (s System) Database() *gorm.DB               { return s.db }
func (s System) Gin() *gin.Engine                 { return s.gin }
func (s System) JetStream() nats.JetStreamContext { return s.js }
func (s System) Logger() logger.Logger            { return s.logger }
func (s System) RPC() *rpc.RPCServer              { return s.grpcServer }
func (s System) Waiter() waiter.Waiter            { return s.waiter }

func (s *System) initOpenTelemetry() error {
	exporter, err := otlptracegrpc.New(context.Background(),
		otlptracegrpc.WithEndpoint(s.cfg.Infrastructure.Otel.ExporterEndpoint),
		otlptracegrpc.WithInsecure(),
	)
	if err != nil {
		return err
	}

	res, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			semconv.ServiceName(s.cfg.Infrastructure.Otel.ServiceName)),
	)
	if err != nil {
		return err
	}

	s.tp = sdktrace.NewTracerProvider(
		sdktrace.WithResource(res),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
	)

	otel.SetTracerProvider(s.tp)

	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	s.waiter.Cleanup(func() {
		if err := s.tp.Shutdown(context.Background()); err != nil {
			s.logger.WithError(err).Error("ran into an issue shutting down the tracer provider")
		}
	})

	return nil
}

func (s *System) initGinEngine() {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(cors.Default())
	router.Use(requestid.New())
	router.Use(gin.Logger())
	router.GET("/liveness", HealthCheck)
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	s.gin = router
}

func HealthCheck(c *gin.Context) {
	message := "OK"
	c.String(http.StatusOK, "\n"+message)
}

func (s *System) initGrpcServer(cfg *config.Server) {
	s.grpcServer = rpc.NewGrpcServer(s.logger, cfg)
}

func (s *System) initWaiter() {
	s.waiter = waiter.New(waiter.CatchSignals())
}

func (s *System) MigrateDB(fs fs.FS) error {
	goose.SetBaseFS(fs)
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}
	if err := goose.Up(pkggorm.SqlDB, "."); err != nil {
		return err
	}
	return nil
}

func (s *System) WaitForWeb(ctx context.Context) error {
	s.genericHTTPServer = server.NewGenericHttpServer(s.gin, s.cfg.Server)

	eg, gCtx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		s.logger.Infof("web server started: %d", s.cfg.Server.HTTP.Port)
		defer s.logger.Infof("web server shutdown")
		err := s.genericHTTPServer.ListenAndServe(gCtx)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}

		return nil
	})

	return eg.Wait()
}

func (s *System) WaitForRPC(ctx context.Context) error {
	address := fmt.Sprintf("%s:%d", s.cfg.Server.GRPC.Host, s.cfg.Server.GRPC.Port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		s.logger.Fatalf("failed to listen: %v", err)
		return err
	}

	eg, gCtx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		s.logger.Infof("rpc server started: %d", s.cfg.Server.GRPC.Port)
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
		timeout := time.NewTimer(30 * time.Second)
		select {
		case <-timeout.C:
			s.RPC().GRPCServer().Stop()
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
