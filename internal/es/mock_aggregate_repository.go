// Code generated by mockery v2.47.0. DO NOT EDIT.

package es

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockAggregateRepository is an autogenerated mock type for the AggregateRepository type
type MockAggregateRepository[T EventSourcedAggregate] struct {
	mock.Mock
}

// Load provides a mock function with given fields: ctx, aggregateID
func (_m *MockAggregateRepository[T]) Load(ctx context.Context, aggregateID string) (T, error) {
	ret := _m.Called(ctx, aggregateID)

	if len(ret) == 0 {
		panic("no return value specified for Load")
	}

	var r0 T
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (T, error)); ok {
		return rf(ctx, aggregateID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) T); ok {
		r0 = rf(ctx, aggregateID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(T)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, aggregateID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, aggregate
func (_m *MockAggregateRepository[T]) Save(ctx context.Context, aggregate T) error {
	ret := _m.Called(ctx, aggregate)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, T) error); ok {
		r0 = rf(ctx, aggregate)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockAggregateRepository creates a new instance of MockAggregateRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAggregateRepository[T EventSourcedAggregate](t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAggregateRepository[T] {
	mock := &MockAggregateRepository[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
