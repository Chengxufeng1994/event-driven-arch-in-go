// Code generated by mockery v2.47.0. DO NOT EDIT.

package am

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockMessagePublisher is an autogenerated mock type for the MessagePublisher type
type MockMessagePublisher[O any] struct {
	mock.Mock
}

// Publish provides a mock function with given fields: ctx, topicName, v
func (_m *MockMessagePublisher[O]) Publish(ctx context.Context, topicName string, v O) error {
	ret := _m.Called(ctx, topicName, v)

	if len(ret) == 0 {
		panic("no return value specified for Publish")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, O) error); ok {
		r0 = rf(ctx, topicName, v)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockMessagePublisher creates a new instance of MockMessagePublisher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMessagePublisher[O any](t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMessagePublisher[O] {
	mock := &MockMessagePublisher[O]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
