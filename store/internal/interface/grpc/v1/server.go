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

type handler struct {
	app usecase.StoreUseCase
	storev1.UnimplementedStoresServiceServer
}

var _ storev1.StoresServiceServer = (*handler)(nil)

func RegisterServer(_ context.Context, application usecase.StoreUseCase, registrar grpc.ServiceRegistrar) error {
	handler := &handler{
		app: application,
	}
	storev1.RegisterStoresServiceServer(registrar, handler)
	return nil
}

// AddProduct implements storev1.StoresServiceServer.
func (s *handler) AddProduct(ctx context.Context, request *storev1.AddProductRequest) (*storev1.AddProductResponse, error) {
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
func (s *handler) CreateStore(ctx context.Context, request *storev1.CreateStoreRequest) (*storev1.CreateStoreResponse, error) {
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
func (s *handler) DisableParticipation(ctx context.Context, request *storev1.DisableParticipationRequest) (*storev1.DisableParticipationResponse, error) {
	err := s.app.DisableParticipation(ctx, command.DisableParticipation{
		ID: request.GetId(),
	})
	if err != nil {
		return nil, err
	}

	return &storev1.DisableParticipationResponse{}, nil
}

// EnableParticipation implements storev1.StoresServiceServer.
func (s *handler) EnableParticipation(ctx context.Context, request *storev1.EnableParticipationRequest) (*storev1.EnableParticipationResponse, error) {
	err := s.app.EnableParticipation(ctx, command.EnableParticipation{
		ID: request.GetId(),
	})
	if err != nil {
		return nil, err
	}

	return &storev1.EnableParticipationResponse{}, nil
}

// GetCatalog implements storev1.StoresServiceServer.
func (s *handler) GetCatalog(ctx context.Context, request *storev1.GetCatalogRequest) (*storev1.GetCatalogResponse, error) {
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

// GetParticipatingStores implements storev1.StoresServiceServer.
func (s *handler) GetParticipatingStores(ctx context.Context, request *storev1.GetParticipatingStoresRequest) (*storev1.GetParticipatingStoresResponse, error) {
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

// GetProduct implements storev1.StoresServiceServer.
func (s *handler) GetProduct(ctx context.Context, request *storev1.GetProductRequest) (*storev1.GetProductResponse, error) {
	product, err := s.app.GetProduct(ctx, query.GetProduct{
		ID: request.GetId(),
	})
	if err != nil {
		return nil, err
	}

	return &storev1.GetProductResponse{Product: s.productFromDomain(product)}, nil
}

// GetStore implements storev1.StoresServiceServer.
func (s *handler) GetStore(ctx context.Context, request *storev1.GetStoreRequest) (*storev1.GetStoreResponse, error) {
	store, err := s.app.GetStore(ctx, query.GetStore{ID: request.GetId()})
	if err != nil {
		return nil, err
	}

	return &storev1.GetStoreResponse{Store: s.storeFromDomain(store)}, nil
}

// GetStores implements storev1.StoresServiceServer.
func (s *handler) GetStores(ctx context.Context, request *storev1.GetStoresRequest) (*storev1.GetStoresResponse, error) {
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

// RemoveProduct implements storev1.StoresServiceServer.
func (s *handler) RemoveProduct(ctx context.Context, request *storev1.RemoveProductRequest) (*storev1.RemoveProductResponse, error) {
	err := s.app.RemoveProduct(ctx, command.RemoveProduct{
		ID: request.GetId(),
	})

	return &storev1.RemoveProductResponse{}, err
}

func (s *handler) storeFromDomain(store *aggregate.StoreAgg) *storev1.Store {
	return &storev1.Store{
		Id:            store.ID,
		Name:          store.Name,
		Location:      store.Location,
		Participating: store.Participating,
	}
}

func (s *handler) productFromDomain(product *aggregate.ProductAgg) *storev1.Product {
	return &storev1.Product{
		Id:          product.ID,
		StoreId:     product.StoreID,
		Name:        product.Name,
		Description: product.Description,
		Sku:         product.SKU,
		Price:       product.Price,
	}
}
