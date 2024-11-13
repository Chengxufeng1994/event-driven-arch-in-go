// Code generated by mockery v2.47.0. DO NOT EDIT.

package repository

import (
	context "context"

	aggregate "github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/domain/aggregate"

	mock "github.com/stretchr/testify/mock"
)

// MockCustomerRepository is an autogenerated mock type for the CustomerRepository type
type MockCustomerRepository struct {
	mock.Mock
}

// Find provides a mock function with given fields: ctx, customerID
func (_m *MockCustomerRepository) Find(ctx context.Context, customerID string) (*aggregate.Customer, error) {
	ret := _m.Called(ctx, customerID)

	if len(ret) == 0 {
		panic("no return value specified for Find")
	}

	var r0 *aggregate.Customer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*aggregate.Customer, error)); ok {
		return rf(ctx, customerID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *aggregate.Customer); ok {
		r0 = rf(ctx, customerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*aggregate.Customer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, customerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, customer
func (_m *MockCustomerRepository) Save(ctx context.Context, customer *aggregate.Customer) error {
	ret := _m.Called(ctx, customer)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *aggregate.Customer) error); ok {
		r0 = rf(ctx, customer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, customer
func (_m *MockCustomerRepository) Update(ctx context.Context, customer *aggregate.Customer) error {
	ret := _m.Called(ctx, customer)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *aggregate.Customer) error); ok {
		r0 = rf(ctx, customer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockCustomerRepository creates a new instance of MockCustomerRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCustomerRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCustomerRepository {
	mock := &MockCustomerRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
