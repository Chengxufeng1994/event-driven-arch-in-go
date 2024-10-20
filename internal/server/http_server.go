package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/config"
	"golang.org/x/sync/errgroup"
)

type GenericHttpServer struct {
	config  *config.Server
	handler http.Handler
}

const shutdownTimeout = 5

func NewGenericHttpServer(handler http.Handler, config *config.Server) *GenericHttpServer {
	return &GenericHttpServer{
		config:  config,
		handler: handler,
	}
}

func (s *GenericHttpServer) ListenAndServe(ctx context.Context) error {
	return s.listenAndServe(ctx)
}

func (s *GenericHttpServer) listenAndServe(ctx context.Context) error {
	addr := fmt.Sprintf("%s:%d", s.config.HTTP.Host, s.config.HTTP.Port)
	srv := &http.Server{
		Addr:              addr,
		Handler:           s.handler,
		ReadHeaderTimeout: 10 * time.Second,
	}

	eg, gCtx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		<-gCtx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout*time.Second)
		defer cancel()
		return srv.Shutdown(ctx)
	})

	eg.Go(func() error {
		err := srv.ListenAndServe()
		if err != nil || !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil
	})

	return eg.Wait()
}
