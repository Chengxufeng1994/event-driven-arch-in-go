// Code generated by mockery v2.47.0. DO NOT EDIT.

package es

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockAggregateStore is an autogenerated mock type for the AggregateStore type
type MockAggregateStore struct {
	mock.Mock
}

// Load provides a mock function with given fields: ctx, aggregate
func (_m *MockAggregateStore) Load(ctx context.Context, aggregate EventSourcedAggregate) error {
	ret := _m.Called(ctx, aggregate)

	if len(ret) == 0 {
		panic("no return value specified for Load")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, EventSourcedAggregate) error); ok {
		r0 = rf(ctx, aggregate)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: ctx, aggregate
func (_m *MockAggregateStore) Save(ctx context.Context, aggregate EventSourcedAggregate) error {
	ret := _m.Called(ctx, aggregate)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, EventSourcedAggregate) error); ok {
		r0 = rf(ctx, aggregate)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockAggregateStore creates a new instance of MockAggregateStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAggregateStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAggregateStore {
	mock := &MockAggregateStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
