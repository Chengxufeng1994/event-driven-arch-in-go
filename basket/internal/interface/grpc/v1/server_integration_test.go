//go:build integration

package v1

import (
	"context"
	"net"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	basketv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/basket/api/basket/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/entity"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/valueobject"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/es"
)

type serverSuite struct {
	mocks struct {
		baskets   *repository.MockBasketRepository
		stores    *repository.MockStoreRepository
		products  *repository.MockProductRepository
		publisher *ddd.MockEventPublisher[ddd.Event]
	}
	server *grpc.Server
	client basketv1.BasketServiceClient
	suite.Suite
}

func TestServer(t *testing.T) {
	suite.Run(t, &serverSuite{})
}

func (s *serverSuite) SetupSuite()    {}
func (s *serverSuite) TearDownSuite() {}

func (s *serverSuite) SetupTest() {
	const grpcTestPort = ":10912"

	var err error
	// create server
	s.server = grpc.NewServer()
	var listener net.Listener
	listener, err = net.Listen("tcp", grpcTestPort)
	if err != nil {
		s.T().Fatal(err)
	}

	// create mocks
	s.mocks = struct {
		baskets   *repository.MockBasketRepository
		stores    *repository.MockStoreRepository
		products  *repository.MockProductRepository
		publisher *ddd.MockEventPublisher[ddd.Event]
	}{
		baskets:   repository.NewMockBasketRepository(s.T()),
		stores:    repository.NewMockStoreRepository(s.T()),
		products:  repository.NewMockProductRepository(s.T()),
		publisher: ddd.NewMockEventPublisher[ddd.Event](s.T()),
	}

	// create app
	app := application.NewBasketApplication(s.mocks.baskets, s.mocks.stores, s.mocks.products, s.mocks.publisher)

	// register app with server
	if err = RegisterServer(app, s.server); err != nil {
		s.T().Fatal(err)
	}
	go func(listener net.Listener) {
		err := s.server.Serve(listener)
		if err != nil {
			s.T().Fatal(err)
		}
	}(listener)

	// create client
	var conn *grpc.ClientConn
	conn, err = grpc.Dial(grpcTestPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		s.T().Fatal(err)
	}
	s.client = basketv1.NewBasketServiceClient(conn)
}
func (s *serverSuite) TearDownTest() {
	s.server.GracefulStop()
}

func (s *serverSuite) TestBasketService_StartBasket() {
	s.mocks.baskets.On("Load", mock.Anything, mock.AnythingOfType("string")).Return(&aggregate.Basket{
		Aggregate: es.NewAggregate("basket-id", aggregate.BasketAggregate),
	}, nil)
	s.mocks.baskets.On("Save", mock.Anything, mock.AnythingOfType("*aggregate.Basket")).Return(nil)
	s.mocks.publisher.On("Publish", mock.Anything, mock.AnythingOfType("ddd.event")).Return(nil)

	_, err := s.client.StartBasket(context.Background(), &basketv1.StartBasketRequest{CustomerId: "customer-id"})
	s.Assert().NoError(err)
}

func (s *serverSuite) TestBasketService_CancelBasket() {
	s.mocks.baskets.On("Load", mock.Anything, "basket-id").Return(&aggregate.Basket{
		Aggregate:  es.NewAggregate("basket-id", aggregate.BasketAggregate),
		CustomerID: "customer-id",
		Status:     valueobject.BasketIsOpen,
	}, nil)
	s.mocks.baskets.On("Save", mock.Anything, mock.AnythingOfType("*aggregate.Basket")).Return(nil)
	s.mocks.publisher.On("Publish", mock.Anything, mock.AnythingOfType("ddd.event")).Return(nil)

	_, err := s.client.CancelBasket(context.Background(), &basketv1.CancelBasketRequest{Id: "basket-id"})
	s.Assert().NoError(err)
}

func (s *serverSuite) TestBasketService_CheckoutBasket() {
	s.mocks.baskets.On("Load", mock.Anything, "basket-id").Return(&aggregate.Basket{
		Aggregate:  es.NewAggregate("basket-id", aggregate.BasketAggregate),
		CustomerID: "customer-id",
		Items: map[string]*entity.Item{
			"product-id": {
				StoreID:      "store-id",
				ProductID:    "product-id",
				StoreName:    "store-name",
				ProductName:  "product-name",
				ProductPrice: 1.00,
				Quantity:     1,
			},
		},
		Status: valueobject.BasketIsOpen,
	}, nil)
	s.mocks.baskets.On("Save", mock.Anything, mock.AnythingOfType("*aggregate.Basket")).Return(nil)
	s.mocks.publisher.On("Publish", mock.Anything, mock.AnythingOfType("ddd.event")).Return(nil)

	_, err := s.client.CheckoutBasket(context.Background(), &basketv1.CheckoutBasketRequest{
		Id:        "basket-id",
		PaymentId: "payment-id",
	})
	s.Assert().NoError(err)
}

func (s *serverSuite) TestBasketService_AddItem() {
	product := &entity.Product{
		ID:      "product-id",
		StoreID: "store-id",
		Name:    "product-name",
		Price:   10.00,
	}
	store := &entity.Store{
		ID:   "store-id",
		Name: "store-name",
	}
	s.mocks.baskets.On("Load", mock.Anything, "basket-id").Return(&aggregate.Basket{
		Aggregate:  es.NewAggregate("basket-id", aggregate.BasketAggregate),
		CustomerID: "customer-id",
		Items: map[string]*entity.Item{
			"product-id": {
				StoreID:      "store-id",
				ProductID:    "product-id",
				StoreName:    "store-name",
				ProductName:  "product-name",
				ProductPrice: 1.00,
				Quantity:     1,
			},
		},
		Status: valueobject.BasketIsOpen,
	}, nil)
	s.mocks.baskets.On("Save", mock.Anything, mock.AnythingOfType("*aggregate.Basket")).Return(nil)
	s.mocks.products.On("Find", mock.Anything, "product-id").Return(product, nil)
	s.mocks.stores.On("Find", mock.Anything, "store-id").Return(store, nil)

	_, err := s.client.AddItem(context.Background(), &basketv1.AddItemRequest{
		Id:        "basket-id",
		ProductId: "product-id",
		Quantity:  1,
	})
	s.Assert().NoError(err)
}

func (s *serverSuite) TestBasketService_RemoveItem() {
	product := &entity.Product{
		ID:      "product-id",
		StoreID: "store-id",
		Name:    "product-name",
		Price:   10.00,
	}

	s.mocks.baskets.On("Load", mock.Anything, "basket-id").Return(&aggregate.Basket{
		Aggregate:  es.NewAggregate("basket-id", aggregate.BasketAggregate),
		CustomerID: "customer-id",
		Items: map[string]*entity.Item{
			"product-id": {
				StoreID:      "store-id",
				ProductID:    "product-id",
				StoreName:    "store-name",
				ProductName:  "product-name",
				ProductPrice: 1.00,
				Quantity:     1,
			},
		},
		Status: valueobject.BasketIsOpen,
	}, nil)
	s.mocks.baskets.On("Save", mock.Anything, mock.AnythingOfType("*aggregate.Basket")).Return(nil)
	s.mocks.products.On("Find", mock.Anything, "product-id").Return(product, nil)

	_, err := s.client.RemoveItem(context.Background(), &basketv1.RemoveItemRequest{
		Id:        "basket-id",
		ProductId: "product-id",
		Quantity:  1,
	})
	s.Assert().NoError(err)
}

func (s *serverSuite) TestBasketService_GetBasket() {
	basket := &aggregate.Basket{
		Aggregate:  es.NewAggregate("basket-id", aggregate.BasketAggregate),
		CustomerID: "customer-id",
		Items: map[string]*entity.Item{
			"product-id": {
				StoreID:      "store-id",
				ProductID:    "product-id",
				StoreName:    "store-name",
				ProductName:  "product-name",
				ProductPrice: 1.00,
				Quantity:     1,
			},
		},
		Status: valueobject.BasketIsOpen,
	}
	s.mocks.baskets.On("Load", mock.Anything, "basket-id").Return(basket, nil)

	resp, err := s.client.GetBasket(context.Background(), &basketv1.GetBasketRequest{Id: "basket-id"})
	if s.Assert().NoError(err) {
		s.Assert().Equal(basket.ID(), resp.Basket.GetId())
	}
}
