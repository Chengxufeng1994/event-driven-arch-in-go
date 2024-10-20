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
		f := make(map[string]any, len(fields)/2)
		i := logging.Fields(fields).Iterator()
		for i.Next() {
			k, v := i.At()
			f[k] = v
		}
		logger := logger.WithFields(f)

		switch lvl {
		case logging.LevelDebug:
			logger.Debug(msg)
		case logging.LevelInfo:
			logger.Info(msg)
		case logging.LevelWarn:
			logger.Warn(msg)
		case logging.LevelError:
			logger.Error(msg)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}
	})
}

func NewGrpcServer(logger logger.Logger, config *config.Server) *RpcServer {

	opts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
		// Add any other option (check functions starting with logging.With).
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(interceptorLogger(logger), opts...),
			recovery.UnaryServerInterceptor(recovery.WithRecoveryHandler(grpcPanicRecoveryHandler)),
		),
		grpc.ChainStreamInterceptor(
			logging.StreamServerInterceptor(interceptorLogger(logger), opts...),
			// Add any other interceptor you want.
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