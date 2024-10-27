package mapper

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/infrastructure/persistence/gorm/po"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
)

type CustomerMapperIntf interface {
	ToPersistent(customer *aggregate.Customer) *po.Customer
	ToDomain(customer *po.Customer) *aggregate.Customer
	ToDomainList(customers []*po.Customer) []*aggregate.Customer
}

type CustomerMapper struct{}

var _ CustomerMapperIntf = (*CustomerMapper)(nil)

func NewCustomerMapper() *CustomerMapper {
	return &CustomerMapper{}
}

func (c *CustomerMapper) ToPersistent(customer *aggregate.Customer) *po.Customer {
	return &po.Customer{
		ID:        customer.ID,
		Name:      customer.Name,
		SmsNumber: customer.SmsNumber,
		Enabled:   customer.Enabled,
	}
}

func (c *CustomerMapper) ToDomain(customer *po.Customer) *aggregate.Customer {
	return &aggregate.Customer{
		AggregateBase: ddd.NewAggregateBase(customer.ID),
		Name:          customer.Name,
		SmsNumber:     customer.SmsNumber,
		Enabled:       customer.Enabled,
	}
}

func (c *CustomerMapper) ToDomainList(customers []*po.Customer) []*aggregate.Customer {
	rlt := make([]*aggregate.Customer, 0, len(customers))
	for i := 0; i < len(customers); i++ {
		rlt = append(rlt, c.ToDomain(customers[i]))
	}
	return rlt
}
