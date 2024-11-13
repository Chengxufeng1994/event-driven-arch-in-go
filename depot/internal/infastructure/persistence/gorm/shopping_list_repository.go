package gorm

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/infastructure/persistence/gorm/dao"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/infastructure/persistence/gorm/mapper"
	"github.com/stackus/errors"
	"gorm.io/gorm"
)

type GormShoppingListRepository struct {
	db                 *gorm.DB
	shoppingListDao    dao.ShoppingListDaoIntf
	shoppingListMapper mapper.ShoppingListMapperIntf
}

var _ repository.ShoppingListRepository = (*GormShoppingListRepository)(nil)

func NewGormShoppingListRepository(db *gorm.DB) *GormShoppingListRepository {
	return &GormShoppingListRepository{
		db:                 db,
		shoppingListDao:    dao.NewShoppingListDao(db),
		shoppingListMapper: mapper.NewShoppingListMapper(),
	}
}

// Save implements repository.ShoppingListRepository.
func (r *GormShoppingListRepository) Save(ctx context.Context, shoppingList *aggregate.ShoppingList) error {
	shoppingListPO, err := r.shoppingListMapper.ToPersistent(shoppingList)
	if err != nil {
		return errors.Wrap(err, "mapping shopping list")
	}
	return r.shoppingListDao.Save(ctx, shoppingListPO)
}

// Update implements repository.ShoppingListRepository.
func (r *GormShoppingListRepository) Update(ctx context.Context, list *aggregate.ShoppingList) error {
	shoppingListPO, err := r.shoppingListMapper.ToPersistent(list)
	if err != nil {
		return errors.Wrap(err, "mapping shopping list")
	}
	return r.shoppingListDao.Update(ctx, shoppingListPO)
}

// Find implements repository.ShoppingListRepository.
func (r *GormShoppingListRepository) Find(ctx context.Context, shoppingListID string) (*aggregate.ShoppingList, error) {
	shoppingList, err := r.shoppingListDao.Find(ctx, shoppingListID)
	if err != nil {
		return nil, err
	}

	return r.shoppingListMapper.ToDomain(shoppingList)
}

// FindByOrderID implements repository.ShoppingListRepository.
func (r *GormShoppingListRepository) FindByOrderID(ctx context.Context, orderID string) (*aggregate.ShoppingList, error) {
	shoppingList, err := r.shoppingListDao.FindByOrderID(ctx, orderID)
	if err != nil {
		return nil, err
	}

	return r.shoppingListMapper.ToDomain(shoppingList)
}
