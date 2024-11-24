package v1

import (
	"context"

	basketv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/basket/api/basket/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/port/in/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/port/in/query"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/usecase"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/errorsotel"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

type server struct {
	app usecase.BasketUseCase
	basketv1.UnimplementedBasketServiceServer
}

var _ basketv1.BasketServiceServer = (*server)(nil)

func RegisterServer(app usecase.BasketUseCase, registrar grpc.ServiceRegistrar) error {
	basketv1.RegisterBasketServiceServer(registrar, server{app: app})
	return nil
}

func (s server) StartBasket(ctx context.Context, request *basketv1.StartBasketRequest) (*basketv1.StartBasketResponse, error) {
	span := trace.SpanFromContext(ctx)

	basketID := uuid.New().String()

	span.SetAttributes(
		attribute.String("BasketID", basketID),
		attribute.String("CustomerID", request.GetCustomerId()),
	)

	err := s.app.StartBasket(ctx, command.StartBasket{
		ID:         basketID,
		CustomerID: request.GetCustomerId(),
	})
	if err != nil {
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
	}

	return &basketv1.StartBasketResponse{Id: basketID}, err
}

func (s server) CancelBasket(ctx context.Context, request *basketv1.CancelBasketRequest) (*basketv1.CancelBasketResponse, error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("BasketID", request.GetId()),
	)

	err := s.app.CancelBasket(ctx, command.CancelBasket{
		ID: request.GetId(),
	})
	if err != nil {
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
	}

	return &basketv1.CancelBasketResponse{}, err
}

func (s server) CheckoutBasket(ctx context.Context, request *basketv1.CheckoutBasketRequest) (*basketv1.CheckoutBasketResponse, error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("BasketID", request.GetId()),
		attribute.String("PaymentID", request.GetPaymentId()),
	)

	err := s.app.CheckoutBasket(ctx, command.CheckoutBasket{
		ID:        request.GetId(),
		PaymentID: request.GetPaymentId(),
	})
	if err != nil {
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
	}

	return &basketv1.CheckoutBasketResponse{}, err
}

func (s server) AddItem(ctx context.Context, request *basketv1.AddItemRequest) (*basketv1.AddItemResponse, error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("BasketID", request.GetId()),
		attribute.String("ProductID", request.GetProductId()),
	)

	err := s.app.AddItem(ctx, command.AddItem{
		ID:        request.GetId(),
		ProductID: request.GetProductId(),
		Quantity:  int(request.GetQuantity()),
	})
	if err != nil {
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
	}

	return &basketv1.AddItemResponse{}, err
}

func (s server) RemoveItem(ctx context.Context, request *basketv1.RemoveItemRequest) (*basketv1.RemoveItemResponse, error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("BasketID", request.GetId()),
		attribute.String("ProductID", request.GetProductId()),
	)

	err := s.app.RemoveItem(ctx, command.RemoveItem{
		ID:        request.GetId(),
		ProductID: request.GetProductId(),
		Quantity:  int(request.GetQuantity()),
	})
	if err != nil {
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
	}

	return &basketv1.RemoveItemResponse{}, err
}

func (s server) GetBasket(ctx context.Context, request *basketv1.GetBasketRequest) (*basketv1.GetBasketResponse, error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("BasketID", request.GetId()),
	)

	basket, err := s.app.GetBasket(ctx, query.GetBasket{
		ID: request.GetId(),
	})
	if err != nil {
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	return &basketv1.GetBasketResponse{
		Basket: s.basketFromDomain(basket),
	}, nil
}

func (s server) basketFromDomain(basket *aggregate.Basket) *basketv1.Basket {
	protoBasket := &basketv1.Basket{
		Id: basket.ID(),
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
