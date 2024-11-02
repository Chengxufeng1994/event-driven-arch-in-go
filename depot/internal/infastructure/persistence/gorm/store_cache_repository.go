package gorm

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/application/port/out/client"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/valueobject"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/infastructure/persistence/gorm/po"
	"github.com/jackc/pgerrcode"
	"github.com/lib/pq"
	"github.com/stackus/errors"
	"gorm.io/gorm"
)

type GormStoreCacheRepository struct {
	db       *gorm.DB
	fallback client.StoreClient
}

var _ repository.StoreCacheRepository = (*GormStoreCacheRepository)(nil)

func NewGormStoreCacheRepository(db *gorm.DB, fallback client.StoreClient) *GormStoreCacheRepository {
	return &GormStoreCacheRepository{
		db:       db,
		fallback: fallback,
	}
}

// Add implements repository.StoreCacheRepository.
func (r *GormStoreCacheRepository) Add(ctx context.Context, storeID, name, location string) error {
	result := r.db.WithContext(ctx).
		Model(&po.StoreCache{}).
		Create(&po.StoreCache{
			ID:       storeID,
			Name:     name,
			Location: location,
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

// Rename implements repository.StoreCacheRepository.
func (r *GormStoreCacheRepository) Rename(ctx context.Context, storeID string, name string) error {
	result := r.db.WithContext(ctx).
		Model(&po.StoreCache{}).
		Where("id = ?", storeID).
		Update("name", name)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Find implements repository.StoreCacheRepository.
func (r *GormStoreCacheRepository) Find(ctx context.Context, storeID string) (*valueobject.Store, error) {
	var storePO po.StoreCache

	result := r.db.WithContext(ctx).
		Model(&po.StoreCache{}).
		Where("id = ?", storeID).
		First(&storePO)

	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.Wrap(result.Error, "querying store")
		}

		store, err := r.fallback.Find(ctx, storeID)
		if err != nil {
			return nil, errors.Wrap(err, "store fallback failed")
		}

		// attempt to add it to the cache
		return &store, r.Add(ctx, store.ID, store.Name, store.Location)
	}

	store := valueobject.NewStore(storePO.ID, storePO.Name)
	return &store, nil
}
