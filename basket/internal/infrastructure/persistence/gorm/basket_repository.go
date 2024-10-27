package gorm

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/infrastructure/persistence/gorm/mapper"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/infrastructure/persistence/gorm/po"
	"github.com/stackus/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GormBasketRepository struct {
	db           *gorm.DB
	basketMapper mapper.BasketMapperIntf
}

var _ repository.BasketRepository = (*GormBasketRepository)(nil)

func NewGormBasketRepository(db *gorm.DB) *GormBasketRepository {
	return &GormBasketRepository{
		db:           db,
		basketMapper: mapper.NewBasketMapper(),
	}
}

func (r *GormBasketRepository) Save(ctx context.Context, basket *aggregate.Basket) error {
	basketPO, err := r.basketMapper.ToPersistent(basket)
	if err != nil {
		return errors.Wrap(err, "mapping basket")
	}

	result := r.db.WithContext(ctx).
		Model(&po.Basket{}).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{
				"customer_id", "payment_id", "items", "status",
			}),
		}).
		Create(&basketPO)

	return errors.Wrap(result.Error, "inserting basket")
}

func (r *GormBasketRepository) Update(ctx context.Context, basket *aggregate.Basket) error {
	basketPO, err := r.basketMapper.ToPersistent(basket)
	if err != nil {
		return errors.Wrap(err, "mapping basket")
	}

	result := r.db.WithContext(ctx).
		Model(&po.Basket{}).
		Where("id = ?", basket.ID).
		Select("customer_id", "payment_id", "items", "status").
		Updates(&po.Basket{
			CustomerID: basketPO.CustomerID,
			PaymentID:  basketPO.PaymentID,
			Items:      basketPO.Items,
			Status:     basketPO.Status,
		})
	if result.Error != nil {
		return errors.Wrap(result.Error, "updating basket")
	}

	return nil
}

func (r *GormBasketRepository) Find(ctx context.Context, basketID string) (*aggregate.Basket, error) {
	var basket *po.Basket
	result := r.db.WithContext(ctx).
		Where("id = ?", basketID).
		First(&basket)

	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "finding basket")
	}

	return mapper.NewBasketMapper().ToDomain(basket)
}
