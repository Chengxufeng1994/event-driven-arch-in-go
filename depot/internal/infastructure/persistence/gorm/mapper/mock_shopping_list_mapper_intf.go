// Code generated by mockery v2.47.0. DO NOT EDIT.

package mapper

import (
	aggregate "github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/aggregate"
	mock "github.com/stretchr/testify/mock"

	po "github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/infastructure/persistence/gorm/po"
)

// MockShoppingListMapperIntf is an autogenerated mock type for the ShoppingListMapperIntf type
type MockShoppingListMapperIntf struct {
	mock.Mock
}

// ToDomain provides a mock function with given fields: _a0
func (_m *MockShoppingListMapperIntf) ToDomain(_a0 *po.ShoppingList) (*aggregate.ShoppingList, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for ToDomain")
	}

	var r0 *aggregate.ShoppingList
	var r1 error
	if rf, ok := ret.Get(0).(func(*po.ShoppingList) (*aggregate.ShoppingList, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*po.ShoppingList) *aggregate.ShoppingList); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*aggregate.ShoppingList)
		}
	}

	if rf, ok := ret.Get(1).(func(*po.ShoppingList) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ToPersistent provides a mock function with given fields: _a0
func (_m *MockShoppingListMapperIntf) ToPersistent(_a0 *aggregate.ShoppingList) (*po.ShoppingList, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for ToPersistent")
	}

	var r0 *po.ShoppingList
	var r1 error
	if rf, ok := ret.Get(0).(func(*aggregate.ShoppingList) (*po.ShoppingList, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*aggregate.ShoppingList) *po.ShoppingList); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*po.ShoppingList)
		}
	}

	if rf, ok := ret.Get(1).(func(*aggregate.ShoppingList) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockShoppingListMapperIntf creates a new instance of MockShoppingListMapperIntf. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockShoppingListMapperIntf(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockShoppingListMapperIntf {
	mock := &MockShoppingListMapperIntf{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}