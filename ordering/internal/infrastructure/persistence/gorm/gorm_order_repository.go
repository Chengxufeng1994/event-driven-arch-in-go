package gorm

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/infrastructure/persistence/gorm/mapper"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/infrastructure/persistence/gorm/po"
	"github.com/stackus/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// deprecated: use event sourcing
type GormOrderRepository struct {
	db     *gorm.DB
	mapper mapper.OrderMapperIntf
}

func NewGormOrderRepository(db *gorm.DB) *GormOrderRepository {
	return &GormOrderRepository{
		db:     db,
		mapper: mapper.NewOrderMapper(),
	}
}

func (r *GormOrderRepository) Save(ctx context.Context, order *aggregate.Order) error {
	orderPO, err := r.mapper.ToPersistence(order)
	if err != nil {
		return errors.Wrap(err, "mapping order")
	}

	result := r.db.WithContext(ctx).
		Model(&po.Order{}).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{
				"customer_id", "payment_id", "shopping_id", "invoice_id", "items", "status",
			}),
		}).
		Create(&orderPO)

	return errors.Wrap(result.Error, "inserting order")
}

func (r *GormOrderRepository) Update(ctx context.Context, order *aggregate.Order) error {
	orderPO, err := r.mapper.ToPersistence(order)
	if err != nil {
		return errors.Wrap(err, "mapping order")
	}

	result := r.db.WithContext(ctx).
		Model(&po.Order{}).
		Where("id = ?", order.ID).
		Select("customer_id", "payment_id", "shopping_id", "invoice_id", "items", "status").
		Updates(&po.Order{
			CustomerID: orderPO.CustomerID,
			PaymentID:  orderPO.PaymentID,
			ShoppingID: orderPO.ShoppingID,
			InvoiceID:  orderPO.InvoiceID,
			Items:      orderPO.Items,
			Status:     orderPO.Status,
		})

	return errors.Wrap(result.Error, "updating order")
}

func (r *GormOrderRepository) Find(ctx context.Context, orderID string) (*aggregate.Order, error) {
	var order *po.Order

	result := r.db.WithContext(ctx).
		Model(&po.Order{}).
		Where("id = ?", orderID).
		First(&order)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "finding order")
	}

	return r.mapper.ToDomain(order)
}
