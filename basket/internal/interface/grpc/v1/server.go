package v1

import (
	"context"

	basketv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/basket/api/basket/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/port/in/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/port/in/query"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/usecase"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/aggregate"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type server struct {
	app usecase.BasketUseCase
	basketv1.UnimplementedBasketServiceServer
}

var _ basketv1.BasketServiceServer = (*server)(nil)

func RegisterServer(_ context.Context, app usecase.BasketUseCase, registrar grpc.ServiceRegistrar) error {
	basketv1.RegisterBasketServiceServer(registrar, server{app: app})
	return nil
}

func (s server) StartBasket(ctx context.Context, request *basketv1.StartBasketRequest) (*basketv1.StartBasketResponse,
	error,
) {
	basketID := uuid.New().String()
	err := s.app.StartBasket(ctx, command.StartBasket{
		ID:         basketID,
		CustomerID: request.GetCustomerId(),
	})

	return &basketv1.StartBasketResponse{Id: basketID}, err
}

func (s server) CancelBasket(ctx context.Context, request *basketv1.CancelBasketRequest,
) (*basketv1.CancelBasketResponse, error) {
	err := s.app.CancelBasket(ctx, command.CancelBasket{
		ID: request.GetId(),
	})

	return &basketv1.CancelBasketResponse{}, err
}

func (s server) CheckoutBasket(ctx context.Context, request *basketv1.CheckoutBasketRequest,
) (*basketv1.CheckoutBasketResponse, error) {
	err := s.app.CheckoutBasket(ctx, command.CheckoutBasket{
		ID:        request.GetId(),
		PaymentID: request.GetPaymentId(),
	})

	return &basketv1.CheckoutBasketResponse{}, err
}

func (s server) AddItem(ctx context.Context, request *basketv1.AddItemRequest) (*basketv1.AddItemResponse, error) {
	err := s.app.AddItem(ctx, command.AddItem{
		ID:        request.GetId(),
		ProductID: request.GetProductId(),
		Quantity:  int(request.GetQuantity()),
	})

	return &basketv1.AddItemResponse{}, err
}

func (s server) RemoveItem(ctx context.Context, request *basketv1.RemoveItemRequest) (*basketv1.RemoveItemResponse,
	error,
) {
	err := s.app.RemoveItem(ctx, command.RemoveItem{
		ID:        request.GetId(),
		ProductID: request.GetProductId(),
		Quantity:  int(request.GetQuantity()),
	})

	return &basketv1.RemoveItemResponse{}, err
}

func (s server) GetBasket(ctx context.Context, request *basketv1.GetBasketRequest) (*basketv1.GetBasketResponse,
	error,
) {
	basket, err := s.app.GetBasket(ctx, query.GetBasket{
		ID: request.GetId(),
	})
	if err != nil {
		return nil, err
	}

	return &basketv1.GetBasketResponse{
		Basket: s.basketFromDomain(basket),
	}, nil
}

func (s server) basketFromDomain(basket *aggregate.Basket) *basketv1.Basket {
	protoBasket := &basketv1.Basket{
		Id: basket.ID,
	}

	protoBasket.Items = make([]*basketv1.Item, 0, len(basket.Items))

	for _, item := range basket.Items {
		protoBasket.Items = append(protoBasket.Items, &basketv1.Item{
			StoreId:      item.StoreID,
			StoreName:    item.StoreName,
			ProductId:    item.ProductID,
			ProductName:  item.ProductName,
			ProductPrice: item.ProductPrice,
			Quantity:     int32(item.Quantity),
		})
	}

	return protoBasket
}
