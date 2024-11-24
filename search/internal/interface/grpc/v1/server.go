package v1

import (
	"context"

	searchv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/search/api/search/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/application/usecase"
	"google.golang.org/grpc"
)

type server struct {
	app usecase.SearchUseCase
	searchv1.UnimplementedSearchServiceServer
}

var _ searchv1.SearchServiceServer = (*server)(nil)

func RegisterServer(app usecase.SearchUseCase, register grpc.ServiceRegistrar) error {
	searchv1.RegisterSearchServiceServer(register, server{app: app})
	return nil
}

// GetOrder implements v1.SearchServiceServer.
func (s server) GetOrder(ctx context.Context, request *searchv1.GetOrderRequest) (resp *searchv1.GetOrderResponse, err error) {
	// TODO: implement me
	panic("unimplemented")
}

// SearchOrders implements v1.SearchServiceServer.
func (s server) SearchOrders(ctx context.Context, request *searchv1.SearchOrdersRequest) (resp *searchv1.SearchOrdersResponse, err error) {
	// TODO: implement me
	panic("unimplemented")
}
