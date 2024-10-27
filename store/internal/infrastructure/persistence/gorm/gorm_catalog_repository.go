package gorm

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/infrastructure/persistence/gorm/po"
	"github.com/shopspring/decimal"
	"github.com/stackus/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GormCatalogRepository struct {
	db *gorm.DB
}

var _ repository.CatalogRepository = (*GormCatalogRepository)(nil)

func NewGormCatalogRepository(db *gorm.DB) *GormCatalogRepository {
	return &GormCatalogRepository{db: db}
}

// AddProduct implements repository.CatalogRepository.
func (r *GormCatalogRepository) AddProduct(ctx context.Context, productID string, storeID string, name string, description string, sku string, price float64) error {

	result := r.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{
				"store_id", "name", "description", "sku", "price",
			}),
		}).
		Create(&po.Product{
			ID:          productID,
			StoreID:     storeID,
			Name:        name,
			Description: description,
			SKU:         sku,
			Price:       decimal.NewFromFloat(price),
		})

	return errors.Wrap(result.Error, "inserting store")
}

// Rebrand implements repository.CatalogRepository.
func (r *GormCatalogRepository) Rebrand(ctx context.Context, productID string, name string, description string) error {
	result := r.db.WithContext(ctx).
		Model(&po.Product{}).
		Where("id = ?", productID).
		Updates(&po.Product{
			Name:        name,
			Description: description,
		})

	return result.Error
}

// RemoveProduct implements repository.CatalogRepository.
func (r *GormCatalogRepository) RemoveProduct(ctx context.Context, productID string) error {
	result := r.db.WithContext(ctx).
		Unscoped().
		Where("id = ?", productID).
		Delete(&po.Product{})

	return result.Error
}

// UpdatePrice implements repository.CatalogRepository.
func (r *GormCatalogRepository) UpdatePrice(ctx context.Context, productID string, delta float64) error {
	result := r.db.WithContext(ctx).
		Model(&po.Product{}).
		Where("id = ?", productID).Update("price", gorm.Expr("price + ?", delta))

	return result.Error
}

// Find implements repository.CatalogRepository.
func (r *GormCatalogRepository) Find(ctx context.Context, productID string) (*aggregate.CatalogProduct, error) {
	var product po.Product
	result := r.db.WithContext(ctx).
		Model(&po.Product{}).
		Where("id = ?", productID).
		First(&product)

	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "scanning product")
	}

	price, _ := product.Price.Float64()
	return &aggregate.CatalogProduct{
		ID:          product.ID,
		StoreID:     product.StoreID,
		Name:        product.Name,
		Description: product.Description,
		SKU:         product.SKU,
		Price:       price,
	}, nil
}

// GetCatalog implements repository.CatalogRepository.
func (r *GormCatalogRepository) GetCatalog(ctx context.Context, storeID string) ([]*aggregate.CatalogProduct, error) {
	var products []*po.Product

	result := r.db.WithContext(ctx).
		Where("store_id = ?", storeID).
		Find(&products)

	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "scanning products")
	}

	productsDomain := make([]*aggregate.CatalogProduct, 0, len(products))
	for _, product := range products {
		price, _ := product.Price.Float64()
		productsDomain = append(productsDomain, &aggregate.CatalogProduct{
			ID:          product.ID,
			StoreID:     product.StoreID,
			Name:        product.Name,
			Description: product.Description,
			SKU:         product.SKU,
			Price:       price,
		})
	}

	return productsDomain, nil
}
