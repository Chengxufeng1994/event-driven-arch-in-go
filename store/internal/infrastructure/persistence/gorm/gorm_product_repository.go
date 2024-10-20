package gorm

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/infrastructure/persistence/gorm/mapper"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/infrastructure/persistence/gorm/po"
	"github.com/stackus/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GormProductRepository struct {
	db            *gorm.DB
	productMapper mapper.ProductMapperIntf
}

var _ repository.ProductRepository = (*GormProductRepository)(nil)

func NewGormProductRepository(db *gorm.DB) *GormProductRepository {
	return &GormProductRepository{
		db:            db,
		productMapper: mapper.NewProductMapper(),
	}
}

func (r *GormProductRepository) AddProduct(ctx context.Context, product *aggregate.ProductAgg) error {
	productPO := r.productMapper.ToPersistent(product)

	result := r.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{
				"store_id", "name", "description", "sku", "price",
			}),
		}).
		Create(&productPO)

	return errors.Wrap(result.Error, "inserting product")
}

func (r *GormProductRepository) RemoveProduct(ctx context.Context, prodcutID string) error {
	result := r.db.WithContext(ctx).
		Unscoped().
		Where("id = ?", prodcutID).
		Delete(&po.Product{})

	return errors.Wrap(result.Error, "deleting product")
}

func (r *GormProductRepository) FindProduct(ctx context.Context, productID string) (*aggregate.ProductAgg, error) {
	var product po.Product

	result := r.db.WithContext(ctx).
		Where("id = ?", productID).
		First(&product)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "scanning product")
	}

	return r.productMapper.ToDomain(&product), nil
}

func (r *GormProductRepository) GetCatalog(ctx context.Context, storeID string) ([]*aggregate.ProductAgg, error) {
	var products []*po.Product

	result := r.db.WithContext(ctx).
		Where("store_id = ?", storeID).
		Find(&products)

	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "scanning products")
	}

	return r.productMapper.ToDomainList(products), nil
}
