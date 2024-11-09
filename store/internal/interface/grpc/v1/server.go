package v1

import (
	"context"

	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/application/port/in/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/application/port/in/query"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/application/usecase"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type server struct {
	app usecase.StoreUseCase
	storev1.UnimplementedStoresServiceServer
}

var _ storev1.StoresServiceServer = (*server)(nil)

func RegisterServer(app usecase.StoreUseCase, registrar grpc.ServiceRegistrar) error {
	handler := &server{app: app}
	storev1.RegisterStoresServiceServer(registrar, handler)
	return nil
}

func (s *server) AddProduct(ctx context.Context, request *storev1.AddProductRequest) (*storev1.AddProductResponse, error) {
	id := uuid.New().String()
	err := s.app.AddProduct(ctx, command.AddProduct{
		ID:          id,
		StoreID:     request.GetStoreId(),
		Name:        request.GetName(),
		Description: request.GetDescription(),
		SKU:         request.GetSku(),
		Price:       request.GetPrice(),
	})
	if err != nil {
		return nil, err
	}

	return &storev1.AddProductResponse{Id: id}, nil
}

// CreateStore implements storev1.StoresServiceServer.
func (s *server) CreateStore(ctx context.Context, request *storev1.CreateStoreRequest) (*storev1.CreateStoreResponse, error) {
	storeID := uuid.New().String()

	err := s.app.CreateStore(ctx, command.CreateStore{
		ID:       storeID,
		Name:     request.GetName(),
		Location: request.GetLocation(),
	})
	if err != nil {
		return nil, err
	}

	return &storev1.CreateStoreResponse{
		Id: storeID,
	}, nil
}

// DisableParticipation implements storev1.StoresServiceServer.
func (s *server) DisableParticipation(ctx context.Context, request *storev1.DisableParticipationRequest) (*storev1.DisableParticipationResponse, error) {
	err := s.app.DisableParticipation(ctx, command.DisableParticipation{
		ID: request.GetId(),
	})
	if err != nil {
		return nil, err
	}

	return &storev1.DisableParticipationResponse{}, nil
}

// EnableParticipation implements storev1.StoresServiceServer.
func (s *server) EnableParticipation(ctx context.Context, request *storev1.EnableParticipationRequest) (*storev1.EnableParticipationResponse, error) {
	err := s.app.EnableParticipation(ctx, command.EnableParticipation{
		ID: request.GetId(),
	})
	if err != nil {
		return nil, err
	}

	return &storev1.EnableParticipationResponse{}, nil
}

// GetParticipatingStores implements storev1.StoresServiceServer.
func (s *server) GetParticipatingStores(ctx context.Context, request *storev1.GetParticipatingStoresRequest) (*storev1.GetParticipatingStoresResponse, error) {
	stores, err := s.app.GetParticipatingStores(ctx, query.GetParticipatingStores{})
	if err != nil {
		return nil, err
	}

	protoStores := []*storev1.Store{}
	for _, store := range stores {
		protoStores = append(protoStores, s.storeFromDomain(store))
	}

	return &storev1.GetParticipatingStoresResponse{
		Stores: protoStores,
	}, nil
}

func (s *server) RebrandStore(ctx context.Context, request *storev1.RebrandStoreRequest) (*storev1.RebrandStoreResponse, error) {
	err := s.app.RebrandStore(ctx, command.RebrandStore{
		ID:   request.GetId(),
		Name: request.GetName(),
	})

	return &storev1.RebrandStoreResponse{}, err
}

func (s *server) GetStore(ctx context.Context, request *storev1.GetStoreRequest) (*storev1.GetStoreResponse, error) {
	store, err := s.app.GetStore(ctx, query.GetStore{ID: request.GetId()})
	if err != nil {
		return nil, err
	}

	return &storev1.GetStoreResponse{Store: s.storeFromDomain(store)}, nil
}

func (s *server) GetStores(ctx context.Context, request *storev1.GetStoresRequest) (*storev1.GetStoresResponse, error) {
	stores, err := s.app.GetStores(ctx, query.GetStores{})
	if err != nil {
		return nil, err
	}

	protoStores := []*storev1.Store{}
	for _, store := range stores {
		protoStores = append(protoStores, s.storeFromDomain(store))
	}

	return &storev1.GetStoresResponse{
		Stores: protoStores,
	}, nil
}

func (s *server) RebrandProduct(ctx context.Context, request *storev1.RebrandProductRequest) (*storev1.RebrandProductResponse, error) {
	err := s.app.RebrandProduct(ctx, command.RebrandProduct{
		ID:          request.GetId(),
		Name:        request.GetName(),
		Description: request.GetDescription(),
	})

	return &storev1.RebrandProductResponse{}, err
}

func (s *server) RemoveProduct(ctx context.Context, request *storev1.RemoveProductRequest) (*storev1.RemoveProductResponse, error) {
	err := s.app.RemoveProduct(ctx, command.RemoveProduct{
		ID: request.GetId(),
	})

	return &storev1.RemoveProductResponse{}, err
}

func (s *server) IncreaseProductPrice(ctx context.Context, request *storev1.IncreaseProductPriceRequest) (*storev1.IncreaseProductPriceResponse, error) {
	err := s.app.IncreaseProductPrice(ctx, command.IncreaseProductPrice{
		ID:    request.GetId(),
		Price: request.GetPrice(),
	})
	return &storev1.IncreaseProductPriceResponse{}, err
}

func (s *server) DecreaseProductPrice(ctx context.Context, request *storev1.DecreaseProductPriceRequest) (*storev1.DecreaseProductPriceResponse, error) {
	err := s.app.DecreaseProductPrice(ctx, command.DecreaseProductPrice{
		ID:    request.GetId(),
		Price: request.GetPrice(),
	})
	return &storev1.DecreaseProductPriceResponse{}, err
}

// GetProduct implements storev1.StoresServiceServer.
func (s *server) GetProduct(ctx context.Context, request *storev1.GetProductRequest) (*storev1.GetProductResponse, error) {
	product, err := s.app.GetProduct(ctx, query.GetProduct{
		ID: request.GetId(),
	})
	if err != nil {
		return nil, err
	}

	return &storev1.GetProductResponse{Product: s.productFromDomain(product)}, nil
}

// GetCatalog implements storev1.StoresServiceServer.
func (s *server) GetCatalog(ctx context.Context, request *storev1.GetCatalogRequest) (*storev1.GetCatalogResponse, error) {
	products, err := s.app.GetCatalog(ctx, query.GetCatalog{StoreID: request.GetStoreId()})
	if err != nil {
		return nil, err
	}

	protoProducts := []*storev1.Product{}
	for _, product := range products {
		protoProducts = append(protoProducts, s.productFromDomain(product))
	}

	return &storev1.GetCatalogResponse{
		Products: protoProducts,
	}, nil
}

func (s *server) storeFromDomain(store *aggregate.MallStore) *storev1.Store {
	return &storev1.Store{
		Id:            store.ID,
		Name:          store.Name,
		Location:      store.Location,
		Participating: store.Participating,
	}
}

func (s *server) productFromDomain(product *aggregate.CatalogProduct) *storev1.Product {
	return &storev1.Product{
		Id:          product.ID,
		StoreId:     product.StoreID,
		Name:        product.Name,
		Description: product.Description,
		Sku:         product.SKU,
		Price:       product.Price,
	}
}
