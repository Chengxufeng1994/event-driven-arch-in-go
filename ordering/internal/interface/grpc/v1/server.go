package v1

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/errorsotel"
	orderv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/api/order/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application/port/in/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application/port/in/query"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application/usecase"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/valueobject"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

type server struct {
	app usecase.OrderUseCase
	orderv1.UnimplementedOrderingServiceServer
}

var _ orderv1.OrderingServiceServer = (*server)(nil)

func RegisterServer(application usecase.OrderUseCase, registrar grpc.ServiceRegistrar) error {
	server := &server{app: application}
	orderv1.RegisterOrderingServiceServer(registrar, server)
	return nil
}

func (s *server) CreateOrder(ctx context.Context, request *orderv1.CreateOrderRequest) (*orderv1.CreateOrderResponse, error) {
	span := trace.SpanFromContext(ctx)

	id := uuid.New().String()

	span.SetAttributes(
		attribute.String("OrderID", id),
		attribute.String("CustomerID", request.GetCustomerId()),
		attribute.String("PaymentID", request.GetPaymentId()),
	)

	items := make([]valueobject.Item, 0, len(request.GetItems()))
	for _, item := range request.GetItems() {
		items = append(items, s.itemToDomainItem(item))
	}

	cmd := command.CreateOrder{
		ID:         id,
		CustomerID: request.GetCustomerId(),
		PaymentID:  request.GetPaymentId(),
		Items:      items,
	}
	err := s.app.CreateOrder(ctx, cmd)
	if err != nil {
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
	}

	return &orderv1.CreateOrderResponse{Id: id}, err
}

func (s *server) CancelOrder(ctx context.Context, request *orderv1.CancelOrderRequest) (*orderv1.CancelOrderResponse, error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("OrderID", request.GetId()),
	)

	err := s.app.CancelOrder(ctx, command.CancelOrder{ID: request.GetId()})
	if err != nil {
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
	}

	return &orderv1.CancelOrderResponse{}, err
}

func (s *server) ReadyOrder(ctx context.Context, request *orderv1.ReadyOrderRequest) (*orderv1.ReadyOrderResponse, error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("OrderID", request.GetId()),
	)

	err := s.app.ReadyOrder(ctx, command.ReadyOrder{ID: request.GetId()})
	if err != nil {
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
	}

	return &orderv1.ReadyOrderResponse{}, err
}

func (s *server) CompleteOrder(ctx context.Context, request *orderv1.CompleteOrderRequest) (*orderv1.CompleteOrderResponse, error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("OrderID", request.GetId()),
	)

	err := s.app.CompleteOrder(ctx, command.CompleteOrder{ID: request.GetId()})
	if err != nil {
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
	}

	return &orderv1.CompleteOrderResponse{}, err
}

func (s *server) GetOrder(ctx context.Context, request *orderv1.GetOrderRequest) (*orderv1.GetOrderResponse, error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("OrderID", request.GetId()),
	)

	orderAgg, err := s.app.GetOrder(ctx, query.GetOrder{OrderID: request.GetId()})
	if err != nil {
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	return &orderv1.GetOrderResponse{Order: s.orderFromDomain(orderAgg)}, nil
}

func (s *server) itemToDomainItem(item *orderv1.Item) valueobject.Item {
	itemVO := valueobject.NewItem(
		item.GetProductId(),
		item.GetStoreId(),
		item.GetStoreName(),
		item.GetProductName(),
		item.GetPrice(),
		int(item.GetQuantity()),
	)

	return itemVO
}

func (s *server) orderFromDomain(orderAgg *aggregate.Order) *orderv1.Order {
	return &orderv1.Order{
		Id:         orderAgg.ID(),
		CustomerId: orderAgg.CustomerID,
		PaymentId:  orderAgg.PaymentID,
		Items:      s.itemsFromDomainItems(orderAgg.Items),
		Status:     orderAgg.Status.String(),
	}
}

func (s *server) itemFromDomain(item valueobject.Item) *orderv1.Item {
	return &orderv1.Item{
		StoreId:     item.StoreID,
		ProductId:   item.ProductID,
		StoreName:   item.StoreName,
		ProductName: item.ProductName,
		Price:       item.Price,
		Quantity:    int32(item.Quantity),
	}
}

func (s *server) itemsFromDomainItems(itemsDO []valueobject.Item) []*orderv1.Item {
	items := make([]*orderv1.Item, 0, len(itemsDO))

	for _, item := range itemsDO {
		items = append(items, s.itemFromDomain(item))
	}

	return items
}
