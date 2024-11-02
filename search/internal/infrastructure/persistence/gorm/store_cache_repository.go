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

type StoreCacheRepository struct {
	db       *gorm.DB
	fallback out.StoreRepository
}

var _ out.StoreCacheRepository = (*StoreCacheRepository)(nil)

func NewGormStoreCacheRepository(db *gorm.DB, fallback out.StoreRepository) *StoreCacheRepository {
	return &StoreCacheRepository{
		db:       db,
		fallback: fallback,
	}
}

// Add implements out.StoreCacheRepository.
func (r *StoreCacheRepository) Add(ctx context.Context, storeID string, name string) error {
	result := r.db.WithContext(ctx).
		Model(&model.StoreCache{}).
		Create(&model.StoreCache{
			ID:   storeID,
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

// Rename implements out.StoreCacheRepository.
func (r *StoreCacheRepository) Rename(ctx context.Context, storeID string, name string) error {
	result := r.db.WithContext(ctx).
		Model(&model.StoreCache{}).
		Where("id = ?", storeID).
		Update("name", name)

	return result.Error
}

// Find implements out.StoreCacheRepository.
func (r *StoreCacheRepository) Find(ctx context.Context, storeID string) (*domain.Store, error) {
	var storePO model.StoreCache
	result := r.db.WithContext(ctx).
		Model(&model.StoreCache{}).
		Where("id = ?", storeID).
		First(&storePO)

	err := result.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			store, err := r.fallback.Find(ctx, storeID)
			if err != nil {
				return nil, err
			}
			return store, r.Add(ctx, store.ID, store.Name)
		}
		return nil, err
	}

	return &domain.Store{
		ID:   storePO.ID,
		Name: storePO.Name,
	}, nil
}
