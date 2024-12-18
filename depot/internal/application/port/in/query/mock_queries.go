// Code generated by mockery v2.47.0. DO NOT EDIT.

package query

import (
	context "context"

	aggregate "github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/aggregate"

	mock "github.com/stretchr/testify/mock"
)

// MockQueries is an autogenerated mock type for the Queries type
type MockQueries struct {
	mock.Mock
}

// GetShoppingList provides a mock function with given fields: ctx, query
func (_m *MockQueries) GetShoppingList(ctx context.Context, query GetShoppingList) (*aggregate.ShoppingList, error) {
	ret := _m.Called(ctx, query)

	if len(ret) == 0 {
		panic("no return value specified for GetShoppingList")
	}

	var r0 *aggregate.ShoppingList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, GetShoppingList) (*aggregate.ShoppingList, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, GetShoppingList) *aggregate.ShoppingList); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*aggregate.ShoppingList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, GetShoppingList) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockQueries creates a new instance of MockQueries. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockQueries(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockQueries {
	mock := &MockQueries{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
