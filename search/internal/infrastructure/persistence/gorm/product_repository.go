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

type ProductCacheRepository struct {
	db       *gorm.DB
	fallback out.ProductRepository
}

var _ out.ProductCacheRepository = (*ProductCacheRepository)(nil)

func NewGormProductCacheRepository(db *gorm.DB, fallback out.ProductRepository) *ProductCacheRepository {
	return &ProductCacheRepository{
		db:       db,
		fallback: fallback,
	}
}

// Add implements out.ProductCacheRepository.
func (r *ProductCacheRepository) Add(ctx context.Context, productID string, storeID string, name string) error {
	result := r.db.WithContext(ctx).
		Model(&model.ProductCache{}).
		Create(&model.ProductCache{
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

// Rebrand implements out.ProductCacheRepository.
func (r *ProductCacheRepository) Rebrand(ctx context.Context, productID string, name string) error {
	result := r.db.WithContext(ctx).
		Model(&model.ProductCache{}).
		Where("id = ?", productID).
		Update("name", name)

	return result.Error
}

// Remove implements out.ProductCacheRepository.
func (r *ProductCacheRepository) Remove(ctx context.Context, productID string) error {
	result := r.db.WithContext(ctx).
		Model(&model.ProductCache{}).
		Delete(&model.ProductCache{}, productID)

	return result.Error
}

// Find implements out.ProductCacheRepository.
func (r *ProductCacheRepository) Find(ctx context.Context, productID string) (*domain.Product, error) {
	var productPO model.ProductCache

	result := r.db.WithContext(ctx).
		Model(&model.ProductCache{}).
		Where("id = ?", productID).
		First(&productPO)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			product, err := r.fallback.Find(ctx, productID)
			if err != nil {
				return nil, err
			}
			return product, r.Add(ctx, product.ID, product.StoreID, product.Name)
		}
		return nil, result.Error
	}

	return &domain.Product{
		ID:      productPO.ID,
		StoreID: productPO.StoreID,
		Name:    productPO.Name,
	}, nil
}
