package gorm

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/repository"
	"gorm.io/gorm"
)

type GormStoreRepository struct {
	db *gorm.DB
}

var _ repository.StoreRepository = (*GormStoreRepository)(nil)

func NewGormStoreRepository(db *gorm.DB) *GormStoreRepository {
	return &GormStoreRepository{db: db}
}

// Save implements repository.StoreRepository.
func (g *GormStoreRepository) Save(ctx context.Context, store *aggregate.StoreAgg) error {
	panic("unimplemented")
}

// Update implements repository.StoreRepository.
func (g *GormStoreRepository) Update(ctx context.Context, store *aggregate.StoreAgg) error {
	panic("unimplemented")
}

// Delete implements repository.StoreRepository.
func (g *GormStoreRepository) Delete(ctx context.Context, storeID string) error {
	panic("unimplemented")
}

// Find implements repository.StoreRepository.
func (g *GormStoreRepository) Find(ctx context.Context, storeID string) (*aggregate.StoreAgg, error) {
	panic("unimplemented")
}

// FindAll implements repository.StoreRepository.
func (g *GormStoreRepository) FindAll(ctx context.Context) ([]aggregate.StoreAgg, error) {
	panic("unimplemented")
}
