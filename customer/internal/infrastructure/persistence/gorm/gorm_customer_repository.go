package gorm

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/domain/repository"
	"gorm.io/gorm"
)

type GormCustomerRepository struct {
	db *gorm.DB
}

var _ repository.CustomerRepository = (*GormCustomerRepository)(nil)

func NewGormCustomerRepository(db *gorm.DB) *GormCustomerRepository {
	return &GormCustomerRepository{
		db: db,
	}
}

func (g *GormCustomerRepository) Save(ctx context.Context, customer *aggregate.CustomerAgg) error {
	panic("unimplemented")
}

func (g *GormCustomerRepository) Update(ctx context.Context, customer *aggregate.CustomerAgg) error {
	panic("unimplemented")
}

func (g *GormCustomerRepository) Find(ctx context.Context, customerID string) (*aggregate.CustomerAgg, error) {
	panic("unimplemented")
}
