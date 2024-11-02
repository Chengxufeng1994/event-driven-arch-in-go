package gorm

import (
	"context"
	"errors"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/application/port/out/client"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/application/port/out/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/domain/valueobject"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/infrastructure/persistence/gorm/po"
	"github.com/jackc/pgerrcode"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type GormCustomerCacheRepository struct {
	db       *gorm.DB
	fallback client.CustomerClient
}

var _ repository.CustomerCacheRepository = (*GormCustomerCacheRepository)(nil)

func NewGormCustomerCacheRepository(db *gorm.DB, fallback client.CustomerClient) *GormCustomerCacheRepository {
	return &GormCustomerCacheRepository{
		db:       db,
		fallback: fallback,
	}
}

// Add implements repository.CustomerCacheRepository.
func (r *GormCustomerCacheRepository) Add(ctx context.Context, customerID string, name string, smsNumber string) error {
	result := r.db.WithContext(ctx).
		Model(&po.CustomerCache{}).
		Create(&po.CustomerCache{
			ID:        customerID,
			Name:      name,
			SmsNumber: smsNumber,
		})

	if result.Error != nil {
		var pgErr *pq.Error
		if errors.As(result.Error, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				return nil
			}
		}
		return result.Error
	}

	return nil
}

// UpdateSmsNumber implements repository.CustomerCacheRepository.
func (r *GormCustomerCacheRepository) UpdateSmsNumber(ctx context.Context, customerID string, smsNumber string) error {
	result := r.db.WithContext(ctx).
		Model(&po.CustomerCache{}).
		Where("id = ?", customerID).
		Update("sms_number", smsNumber)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Find implements repository.CustomerCacheRepository.
func (r *GormCustomerCacheRepository) Find(ctx context.Context, customerID string) (*valueobject.Customer, error) {
	var customerPO *po.CustomerCache

	result := r.db.WithContext(ctx).
		Model(&po.CustomerCache{}).
		Where("id = ?", customerID).
		First(&customerPO)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			customer, err := r.fallback.Find(ctx, customerID)
			if err != nil {
				return nil, err
			}
			return &customer, r.Add(ctx, customer.ID, customer.Name, customer.SmsNumber)
		}
		return nil, result.Error
	}

	return &valueobject.Customer{
		ID:        customerPO.ID,
		Name:      customerPO.Name,
		SmsNumber: customerPO.SmsNumber,
	}, nil
}
