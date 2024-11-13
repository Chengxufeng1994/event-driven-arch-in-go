package mapper

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/entity"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/infrastructure/persistence/gorm/po"
	"github.com/shopspring/decimal"
)

type ProductCacheMapperIntf interface {
	ToPersistence(store entity.Product) (po.ProductCache, error)
	ToDomain(store po.ProductCache) (entity.Product, error)
}

type ProductCacheMapper struct{}

var _ ProductCacheMapperIntf = (*ProductCacheMapper)(nil)

func NewProductCacheMapper() *ProductCacheMapper {
	return &ProductCacheMapper{}
}

func (p *ProductCacheMapper) ToPersistence(product entity.Product) (po.ProductCache, error) {
	return po.ProductCache{
		ID:      product.ID,
		StoreID: product.StoreID,
		Name:    product.Name,
		Price:   decimal.NewFromFloat(product.Price),
	}, nil
}

func (p *ProductCacheMapper) ToDomain(store po.ProductCache) (entity.Product, error) {
	price, _ := store.Price.Float64()
	return entity.NewProduct(store.ID, store.StoreID, store.Name, price), nil
}
