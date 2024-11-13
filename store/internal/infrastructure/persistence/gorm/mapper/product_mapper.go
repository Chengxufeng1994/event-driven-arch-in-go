package mapper

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/es"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/infrastructure/persistence/gorm/po"
	"github.com/shopspring/decimal"
)

type ProductMapperIntf interface {
	ToPersistent(product *aggregate.Product) *po.Product
	ToDomain(product *po.Product) *aggregate.Product
	ToDomainList(products []*po.Product) []*aggregate.Product
}

type ProductMapper struct {
}

var _ ProductMapperIntf = (*ProductMapper)(nil)

func NewProductMapper() *ProductMapper {
	return &ProductMapper{}
}

func (p *ProductMapper) ToPersistent(product *aggregate.Product) *po.Product {

	return &po.Product{
		ID:          product.ID(),
		StoreID:     product.StoreID,
		Name:        product.Name,
		Description: product.Description,
		SKU:         product.SKU,
		Price:       decimal.NewFromFloat(product.Price),
	}
}

func (p *ProductMapper) ToDomain(product *po.Product) *aggregate.Product {
	price, _ := product.Price.Float64()
	return &aggregate.Product{
		Aggregate:   es.NewAggregate(product.ID, aggregate.ProductAggregate),
		StoreID:     product.StoreID,
		Name:        product.Name,
		Description: product.Description,
		SKU:         product.SKU,
		Price:       price,
	}
}

func (p *ProductMapper) ToDomainList(products []*po.Product) []*aggregate.Product {
	rlt := make([]*aggregate.Product, 0, len(products))

	for _, product := range products {
		rlt = append(rlt, p.ToDomain(product))
	}

	return rlt
}
