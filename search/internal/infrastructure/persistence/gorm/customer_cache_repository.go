package gorm

import (
	"context"
	"errors"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/application/port/out"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/domain"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/infrastructure/persistence/gorm/model"
	"github.com/jackc/pgerrcode"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type CustomerCacheRepository struct {
	db       *gorm.DB
	fallback out.CustomerRepository
}

var _ out.CustomerCacheRepository = (*CustomerCacheRepository)(nil)

func NewGormCustomerCacheRepository(db *gorm.DB, fallback out.CustomerRepository) *CustomerCacheRepository {
	return &CustomerCacheRepository{
		db:       db,
		fallback: fallback,
	}
}

// Add implements out.CustomerCacheRepository.
func (r *CustomerCacheRepository) Add(ctx context.Context, customerID string, name string) error {
	result := r.db.WithContext(ctx).
		Model(&model.CustomerCache{}).
		Create(&model.CustomerCache{
			ID:   customerID,
			Name: name,
		})

	if result.Error != nil {
		var pgErr *pq.Error
		if errors.As(result.Error, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				return nil
			}
		}
	}

	return nil
}

// Find implements out.CustomerCacheRepository.
func (r *CustomerCacheRepository) Find(ctx context.Context, customerID string) (*domain.Customer, error) {
	var customerPO model.CustomerCache
	result := r.db.WithContext(ctx).
		Model(&model.CustomerCache{}).
		Where("id = ?", customerID).
		First(&customerPO)

	err := result.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			customer, err := r.fallback.Find(ctx, customerID)
			if err != nil {
				return nil, err
			}
			return customer, r.Add(ctx, customer.ID, customer.Name)
		}
		return nil, err
	}

	return &domain.Customer{
		ID:   customerPO.ID,
		Name: customerPO.Name,
	}, nil
}
