package gorm

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/repository"
	"gorm.io/gorm"
)

type GormShoppingListRepository struct {
	db *gorm.DB
}

var _ repository.ShoppingListRepository = (*GormShoppingListRepository)(nil)

func NewGormShoppingListRepository(db *gorm.DB) *GormShoppingListRepository {
	return &GormShoppingListRepository{db: db}
}

// Save implements repository.ShoppingListRepository.
func (g *GormShoppingListRepository) Save(ctx context.Context, list *aggregate.ShoppingListAgg) error {
	panic("unimplemented")
}

// Update implements repository.ShoppingListRepository.
func (g *GormShoppingListRepository) Update(ctx context.Context, list *aggregate.ShoppingListAgg) error {
	panic("unimplemented")
}

// Find implements repository.ShoppingListRepository.
func (g *GormShoppingListRepository) Find(ctx context.Context, shoppingListID string) (*aggregate.ShoppingListAgg, error) {
	panic("unimplemented")
}

// FindByOrderID implements repository.ShoppingListRepository.
func (g *GormShoppingListRepository) FindByOrderID(ctx context.Context, orderID string) (*aggregate.ShoppingListAgg, error) {
	panic("unimplemented")
}
