// Code generated by mockery v2.47.0. DO NOT EDIT.

package am

import (
	context "context"

	ddd "github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	mock "github.com/stretchr/testify/mock"
)

// MockEventPublisher is an autogenerated mock type for the EventPublisher type
type MockEventPublisher struct {
	mock.Mock
}

// Publish provides a mock function with given fields: ctx, topicName, event
func (_m *MockEventPublisher) Publish(ctx context.Context, topicName string, event ddd.Event) error {
	ret := _m.Called(ctx, topicName, event)

	if len(ret) == 0 {
		panic("no return value specified for Publish")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ddd.Event) error); ok {
		r0 = rf(ctx, topicName, event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockEventPublisher creates a new instance of MockEventPublisher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockEventPublisher(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockEventPublisher {
	mock := &MockEventPublisher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
