// Code generated by mockery v2.47.0. DO NOT EDIT.

package mapper

import (
	aggregate "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/aggregate"
	mock "github.com/stretchr/testify/mock"

	po "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/infrastructure/persistence/gorm/po"
)

// MockOrderMapperIntf is an autogenerated mock type for the OrderMapperIntf type
type MockOrderMapperIntf struct {
	mock.Mock
}

// ToDomain provides a mock function with given fields: order
func (_m *MockOrderMapperIntf) ToDomain(order *po.Order) (*aggregate.Order, error) {
	ret := _m.Called(order)

	if len(ret) == 0 {
		panic("no return value specified for ToDomain")
	}

	var r0 *aggregate.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(*po.Order) (*aggregate.Order, error)); ok {
		return rf(order)
	}
	if rf, ok := ret.Get(0).(func(*po.Order) *aggregate.Order); ok {
		r0 = rf(order)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*aggregate.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(*po.Order) error); ok {
		r1 = rf(order)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ToPersistence provides a mock function with given fields: order
func (_m *MockOrderMapperIntf) ToPersistence(order *aggregate.Order) (*po.Order, error) {
	ret := _m.Called(order)

	if len(ret) == 0 {
		panic("no return value specified for ToPersistence")
	}

	var r0 *po.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(*aggregate.Order) (*po.Order, error)); ok {
		return rf(order)
	}
	if rf, ok := ret.Get(0).(func(*aggregate.Order) *po.Order); ok {
		r0 = rf(order)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*po.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(*aggregate.Order) error); ok {
		r1 = rf(order)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockOrderMapperIntf creates a new instance of MockOrderMapperIntf. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockOrderMapperIntf(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockOrderMapperIntf {
	mock := &MockOrderMapperIntf{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
