package gorm

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/repository"
	"gorm.io/gorm"
)

type GormBasketRepository struct {
	db *gorm.DB
}

var _ repository.BasketRepository = (*GormBasketRepository)(nil)

func NewGormBasketRepository(db *gorm.DB) *GormBasketRepository {
	return &GormBasketRepository{db: db}
}

func (b *GormBasketRepository) Save(ctx context.Context, basket *aggregate.BasketAgg) error {
	panic("unimplemented")
}

func (b *GormBasketRepository) Update(ctx context.Context, basket *aggregate.BasketAgg) error {
	panic("unimplemented")
}

func (b *GormBasketRepository) Find(ctx context.Context, basketID string) (*aggregate.BasketAgg, error) {
	panic("unimplemented")
}
