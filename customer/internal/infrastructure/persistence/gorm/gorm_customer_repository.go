package gorm

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/infrastructure/persistence/gorm/mapper"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/infrastructure/persistence/gorm/po"
	"github.com/stackus/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GormCustomerRepository struct {
	db     *gorm.DB
	mapper mapper.CustomerMapperIntf
}

var _ repository.CustomerRepository = (*GormCustomerRepository)(nil)

func NewGormCustomerRepository(db *gorm.DB) GormCustomerRepository {
	return GormCustomerRepository{
		db:     db,
		mapper: mapper.NewCustomerMapper(),
	}
}

func (r GormCustomerRepository) Save(ctx context.Context, customer *aggregate.Customer) error {
	customerPO := r.mapper.ToPersistent(customer)

	result := r.db.WithContext(ctx).
		Model(&po.Customer{}).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{
				"name", "sms_number", "enabled",
			}),
		}).
		Create(&customerPO)

	if result.Error != nil {
		return errors.Wrap(result.Error, "inserting customer")
	}

	return nil
}

func (r GormCustomerRepository) Update(ctx context.Context, customer *aggregate.Customer) error {
	customerPO := r.mapper.ToPersistent(customer)
	result := r.db.WithContext(ctx).
		Model(&po.Customer{}).
		Where("id = ?", customer.ID()).
		Select("name", "sms_number", "enabled").
		Updates(&po.Customer{
			Name:      customerPO.Name,
			SmsNumber: customerPO.SmsNumber,
			Enabled:   customerPO.Enabled,
		})

	if result.Error != nil {
		return errors.Wrap(result.Error, "updating customer")
	}

	return nil
}

func (r GormCustomerRepository) Find(ctx context.Context, customerID string) (*aggregate.Customer, error) {
	var customerPO *po.Customer
	result := r.db.WithContext(ctx).
		Model(&po.Customer{}).
		Where("id = ?", customerID).
		First(&customerPO)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "failed to find customer")
	}

	return r.mapper.ToDomain(customerPO), nil
}
