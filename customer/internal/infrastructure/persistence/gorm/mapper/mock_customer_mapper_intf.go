// Code generated by mockery v2.47.0. DO NOT EDIT.

package mapper

import (
	aggregate "github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/domain/aggregate"
	mock "github.com/stretchr/testify/mock"

	po "github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/infrastructure/persistence/gorm/po"
)

// MockCustomerMapperIntf is an autogenerated mock type for the CustomerMapperIntf type
type MockCustomerMapperIntf struct {
	mock.Mock
}

// ToDomain provides a mock function with given fields: customer
func (_m *MockCustomerMapperIntf) ToDomain(customer *po.Customer) *aggregate.Customer {
	ret := _m.Called(customer)

	if len(ret) == 0 {
		panic("no return value specified for ToDomain")
	}

	var r0 *aggregate.Customer
	if rf, ok := ret.Get(0).(func(*po.Customer) *aggregate.Customer); ok {
		r0 = rf(customer)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*aggregate.Customer)
		}
	}

	return r0
}

// ToDomainList provides a mock function with given fields: customers
func (_m *MockCustomerMapperIntf) ToDomainList(customers []*po.Customer) []*aggregate.Customer {
	ret := _m.Called(customers)

	if len(ret) == 0 {
		panic("no return value specified for ToDomainList")
	}

	var r0 []*aggregate.Customer
	if rf, ok := ret.Get(0).(func([]*po.Customer) []*aggregate.Customer); ok {
		r0 = rf(customers)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*aggregate.Customer)
		}
	}

	return r0
}

// ToPersistent provides a mock function with given fields: customer
func (_m *MockCustomerMapperIntf) ToPersistent(customer *aggregate.Customer) *po.Customer {
	ret := _m.Called(customer)

	if len(ret) == 0 {
		panic("no return value specified for ToPersistent")
	}

	var r0 *po.Customer
	if rf, ok := ret.Get(0).(func(*aggregate.Customer) *po.Customer); ok {
		r0 = rf(customer)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*po.Customer)
		}
	}

	return r0
}

// NewMockCustomerMapperIntf creates a new instance of MockCustomerMapperIntf. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCustomerMapperIntf(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCustomerMapperIntf {
	mock := &MockCustomerMapperIntf{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
