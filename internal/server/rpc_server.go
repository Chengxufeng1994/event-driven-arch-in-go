package server

import (
	"context"
	"fmt"
	"time"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/config"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/logger"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type RpcServer struct {
	config     *config.Server
	grpcServer *grpc.Server
}

func GrpcLogger(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	startTime := time.Now()
	result, err := handler(ctx, req)
	duration := time.Since(startTime)

	statusCode := codes.Unknown
	if st, ok := status.FromError(err); ok {
		statusCode = st.Code()
	}

	logger := logger.ContextUnavailable()
	if err != nil {
		logger.Errorf("failed to handle a gRPC request: %v", err)
	}

	logger.WithField("protocol", "grpc").
		WithField("method", info.FullMethod).
		WithField("status_code", int(statusCode)).
		WithField("status_text", statusCode.String()).
		WithField("duration", duration).
		Info("received a gRPC request")

	return result, err
}
func grpcPanicRecoveryHandler(p any) (err error) {
	return status.Errorf(codes.Internal, "%s", p)
}

// interceptorLogger adapts go-kit logger to interceptor logger.
// This code is simple enough to be copied and not imported.
func interceptorLogger(logger logger.Logger) logging.Logger {
	return logging.LoggerFunc(func(_ context.Context, lvl logging.Level, msg string, fields ...any) {
		largs := append([]any{"msg", msg}, fields...)
		switch lvl {
		case logging.LevelDebug:
			logger.Debug(largs...)
		case logging.LevelInfo:
			logger.Info(largs...)
		case logging.LevelWarn:
			logger.Warn(largs...)
		case logging.LevelError:
			logger.Error(largs...)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}
	})
}

func NewGrpcServer(logger logger.Logger, config *config.Server) *RpcServer {
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(interceptorLogger(logger)),
			recovery.UnaryServerInterceptor(recovery.WithRecoveryHandler(grpcPanicRecoveryHandler)),
		),
	)

	reflection.Register(grpcServer)

	return &RpcServer{
		config:     config,
		grpcServer: grpcServer,
	}
}

func (s *RpcServer) GRPCServer() *grpc.Server {
	return s.grpcServer
}
