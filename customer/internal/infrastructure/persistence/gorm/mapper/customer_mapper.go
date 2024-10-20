package mapper

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/infrastructure/persistence/gorm/po"
)

type CustomerMapperIntf interface {
	ToPersistent(customer *aggregate.CustomerAgg) *po.Customer
	ToDomain(customer *po.Customer) *aggregate.CustomerAgg
	ToDomainList(customers []*po.Customer) []*aggregate.CustomerAgg
}

type CustomerMapper struct{}

var _ CustomerMapperIntf = (*CustomerMapper)(nil)

func NewCustomerMapper() *CustomerMapper {
	return &CustomerMapper{}
}

func (c *CustomerMapper) ToPersistent(customer *aggregate.CustomerAgg) *po.Customer {
	return &po.Customer{
		ID:        customer.ID,
		Name:      customer.Name,
		SmsNumber: customer.SmsNumber,
		Enabled:   customer.Enabled,
	}
}

func (c *CustomerMapper) ToDomain(customer *po.Customer) *aggregate.CustomerAgg {
	return &aggregate.CustomerAgg{
		ID:        customer.ID,
		Name:      customer.Name,
		SmsNumber: customer.SmsNumber,
		Enabled:   customer.Enabled,
	}
}

func (c *CustomerMapper) ToDomainList(customers []*po.Customer) []*aggregate.CustomerAgg {
	rlt := make([]*aggregate.CustomerAgg, 0, len(customers))
	for i := 0; i < len(customers); i++ {
		rlt = append(rlt, c.ToDomain(customers[i]))
	}
	return rlt
}
