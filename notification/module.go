package notification

import (
	"context"
	"fmt"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/monolith"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/application"
	infragrpc "github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/infrastructure/client/grpc"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/infrastructure/logging"
	v1 "github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/interface/grpc/v1"
)

type Module struct{}

var _ monolith.Module = Module{}

func NewModule() *Module {
	return &Module{}
}

func (m Module) PrepareRun(ctx context.Context, mono monolith.Monolith) error {
	// setup Driven adapters
	endpoint := fmt.Sprintf("%s:%d", mono.Config().Server.GPPC.Host, mono.Config().Server.GPPC.Port)
	conn, err := infragrpc.Dial(ctx, endpoint)
	if err != nil {
		return err
	}
	customers := infragrpc.NewCustomerClient(conn)

	// setup application
	app := logging.NewLogApplicationAccess(application.New(customers), mono.Logger())

	// setup Driver adapters
	if err := v1.RegisterServer(ctx, app, mono.RPC().GRPCServer()); err != nil {
		return err
	}

	return nil
}

func (m Module) Name() string {
	return "notification"
}
