package gorm

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/repository"
	"gorm.io/gorm"
)

type GormProductRepository struct {
	db *gorm.DB
}

var _ repository.ProductRepository = (*GormProductRepository)(nil)

func NewGormProductRepository(db *gorm.DB) *GormProductRepository {
	return &GormProductRepository{db: db}
}

func (g *GormProductRepository) AddProduct(ctx context.Context, product *aggregate.ProductAgg) error {
	panic("unimplemented")
}

func (g *GormProductRepository) FindProduct(ctx context.Context, id string) (*aggregate.ProductAgg, error) {
	panic("unimplemented")
}

func (g *GormProductRepository) GetCatalog(ctx context.Context, storeID string) ([]aggregate.ProductAgg, error) {
	panic("unimplemented")
}

func (g *GormProductRepository) RemoveProduct(ctx context.Context, id string) error {
	panic("unimplemented")
}
