package dao

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/infastructure/persistence/gorm/po"
	"github.com/stackus/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ShoppingListDaoIntf interface {
	Save(ctx context.Context, shoppingList *po.ShoppingList) error
	Update(ctx context.Context, shoppingList *po.ShoppingList) error
	Find(ctx context.Context, shoppingListID string) (*po.ShoppingList, error)
	FindByOrderID(ctx context.Context, orderID string) (*po.ShoppingList, error)
}

type ShoppingListDao struct {
	db *gorm.DB
}

var _ ShoppingListDaoIntf = (*ShoppingListDao)(nil)

func NewShoppingListDao(db *gorm.DB) *ShoppingListDao {
	return &ShoppingListDao{
		db: db,
	}
}

func (d *ShoppingListDao) Save(ctx context.Context, shoppingList *po.ShoppingList) error {
	result := d.db.WithContext(ctx).
		Model(&po.ShoppingList{}).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"order_id", "stops", "assigned_bot_id", "status"}),
		}).
		Create(&shoppingList)

	if result.Error != nil {
		return errors.Wrap(result.Error, "saving shopping list")
	}

	return nil
}

func (d *ShoppingListDao) Update(ctx context.Context, shoppingList *po.ShoppingList) error {
	result := d.db.WithContext(ctx).
		Model(&po.ShoppingList{}).
		Where("id = ?", shoppingList.ID).
		Select("order_id", "stops", "assigned_bot_id", "status").
		Updates(&shoppingList)

	if result.Error != nil {
		return errors.Wrap(result.Error, "updating shopping list")
	}

	return nil
}

func (d *ShoppingListDao) Find(ctx context.Context, shoppingListID string) (*po.ShoppingList, error) {
	var shoppingList po.ShoppingList
	result := d.db.WithContext(ctx).
		Model(&po.ShoppingList{}).
		Where("id = ?", shoppingListID).
		First(&shoppingList)

	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "finding shopping list")
	}

	return &shoppingList, nil
}

func (d *ShoppingListDao) FindByOrderID(ctx context.Context, orderID string) (*po.ShoppingList, error) {
	var shoppingList po.ShoppingList
	result := d.db.WithContext(ctx).
		Model(&po.ShoppingList{}).
		Where("order_id = ?", orderID).
		First(&shoppingList)

	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "finding shopping list")
	}

	return &shoppingList, nil
}
