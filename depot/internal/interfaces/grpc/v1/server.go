package v1

import (
	"context"

	depotv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/depot/api/depot/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/application/port/in/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/application/usecase"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/valueobject"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type server struct {
	app usecase.ShoppingListUseCase
	depotv1.UnimplementedDepotServiceServer
}

var _ depotv1.DepotServiceServer = (*server)(nil)

func RegisterServer(app usecase.ShoppingListUseCase, registrar grpc.ServiceRegistrar) error {
	depotv1.RegisterDepotServiceServer(registrar, server{app: app})
	return nil
}

func (s server) CreateShoppingList(ctx context.Context, request *depotv1.CreateShoppingListRequest) (*depotv1.CreateShoppingListResponse, error) {
	id := uuid.New().String()

	items := make([]valueobject.OrderItem, 0, len(request.GetItems()))
	for _, item := range request.GetItems() {
		items = append(items, s.itemToDomain(item))
	}

	err := s.app.CreateShoppingList(ctx, command.CreateShoppingList{
		ID:      id,
		OrderID: request.GetOrderId(),
		Items:   items,
	})

	return &depotv1.CreateShoppingListResponse{Id: id}, err
}

func (s server) CancelShoppingList(ctx context.Context, request *depotv1.CancelShoppingListRequest) (*depotv1.CancelShoppingListResponse, error) {
	err := s.app.CancelShoppingList(ctx, command.CancelShoppingList{
		ID: request.GetId(),
	})

	return &depotv1.CancelShoppingListResponse{}, err
}

func (s server) AssignShoppingList(ctx context.Context, request *depotv1.AssignShoppingListRequest) (*depotv1.AssignShoppingListResponse, error) {
	err := s.app.AssignShoppingList(ctx, command.AssignShoppingList{
		ID:    request.GetId(),
		BotID: request.GetBotId(),
	})
	return &depotv1.AssignShoppingListResponse{}, err
}

func (s server) CompleteShoppingList(ctx context.Context, request *depotv1.CompleteShoppingListRequest) (*depotv1.CompleteShoppingListResponse, error) {
	err := s.app.CompleteShoppingList(ctx, command.CompleteShoppingList{ID: request.GetId()})
	return &depotv1.CompleteShoppingListResponse{}, err
}

func (s server) itemToDomain(item *depotv1.OrderItem) valueobject.OrderItem {
	return valueobject.NewOrderItem(
		item.GetStoreId(),
		item.GetProductId(),
		int(item.GetQuantity()),
	)
}
