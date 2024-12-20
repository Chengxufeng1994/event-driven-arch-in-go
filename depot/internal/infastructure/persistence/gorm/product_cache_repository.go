package gorm

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/entity"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/infastructure/persistence/gorm/po"
	"github.com/jackc/pgerrcode"
	"github.com/lib/pq"
	"github.com/stackus/errors"
	"gorm.io/gorm"
)

type GormProductCacheRepository struct {
	db       *gorm.DB
	fallback repository.ProductRepository
}

var _ repository.ProductCacheRepository = (*GormProductCacheRepository)(nil)

func NewGormProductCacheRepository(db *gorm.DB, fallback repository.ProductRepository) *GormProductCacheRepository {
	return &GormProductCacheRepository{
		db:       db,
		fallback: fallback,
	}
}

// Add implements repository.ProductCacheRepository.
func (r *GormProductCacheRepository) Add(ctx context.Context, productID string, storeID string, name string) error {
	result := r.db.WithContext(ctx).
		Model(&po.ProductCache{}).
		Create(&po.ProductCache{
			ID:      productID,
			StoreID: storeID,
			Name:    name,
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

// Rebrand implements repository.ProductCacheRepository.
func (r *GormProductCacheRepository) Rebrand(ctx context.Context, productID string, name string) error {
	result := r.db.WithContext(ctx).
		Model(&po.ProductCache{}).
		Where("id = ?", productID).
		Update("name", name)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Remove implements repository.ProductCacheRepository.
func (r *GormProductCacheRepository) Remove(ctx context.Context, productID string) error {
	result := r.db.WithContext(ctx).
		Unscoped().
		Delete(&po.ProductCache{}, productID)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// UpdatePrice implements repository.ProductCacheRepository.
func (r *GormProductCacheRepository) UpdatePrice(ctx context.Context, productID string, delta float64) error {
	result := r.db.WithContext(ctx).
		Model(&po.ProductCache{}).
		Where("id = ?", productID).
		Update("price", delta)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Find implements repository.ProductCacheRepository.
func (r *GormProductCacheRepository) Find(ctx context.Context, productID string) (*entity.Product, error) {
	var productPO po.ProductCache

	result := r.db.WithContext(ctx).
		Model(&po.ProductCache{}).
		Where("id = ?", productID).
		First(&productPO)

	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.Wrap(result.Error, "querying store")
		}

		product, err := r.fallback.Find(ctx, productID)
		if err != nil {
			return nil, errors.Wrap(err, "product fallback failed")
		}

		// attempt to add it to the cache
		return product, r.Add(ctx, product.ID, product.StoreID, product.Name)
	}

	return &entity.Product{
		ID:      productPO.ID,
		StoreID: productPO.StoreID,
		Name:    productPO.Name,
	}, nil
}
