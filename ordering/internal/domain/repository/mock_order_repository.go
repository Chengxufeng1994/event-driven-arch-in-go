// Code generated by mockery v2.47.0. DO NOT EDIT.

package repository

import (
	context "context"

	aggregate "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/aggregate"

	mock "github.com/stretchr/testify/mock"
)

// MockOrderRepository is an autogenerated mock type for the OrderRepository type
type MockOrderRepository struct {
	mock.Mock
}

// Load provides a mock function with given fields: ctx, orderID
func (_m *MockOrderRepository) Load(ctx context.Context, orderID string) (*aggregate.Order, error) {
	ret := _m.Called(ctx, orderID)

	if len(ret) == 0 {
		panic("no return value specified for Load")
	}

	var r0 *aggregate.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*aggregate.Order, error)); ok {
		return rf(ctx, orderID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *aggregate.Order); ok {
		r0 = rf(ctx, orderID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*aggregate.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, orderID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, order
func (_m *MockOrderRepository) Save(ctx context.Context, order *aggregate.Order) error {
	ret := _m.Called(ctx, order)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *aggregate.Order) error); ok {
		r0 = rf(ctx, order)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockOrderRepository creates a new instance of MockOrderRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockOrderRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockOrderRepository {
	mock := &MockOrderRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
