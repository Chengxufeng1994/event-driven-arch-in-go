// Code generated by mockery v2.47.0. DO NOT EDIT.

package ddd

import mock "github.com/stretchr/testify/mock"

// MockEventSubscriber is an autogenerated mock type for the EventSubscriber type
type MockEventSubscriber[T Event] struct {
	mock.Mock
}

// Subscribe provides a mock function with given fields: eventHandler, events
func (_m *MockEventSubscriber[T]) Subscribe(eventHandler EventHandler[T], events ...string) {
	_va := make([]interface{}, len(events))
	for _i := range events {
		_va[_i] = events[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, eventHandler)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// NewMockEventSubscriber creates a new instance of MockEventSubscriber. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockEventSubscriber[T Event](t interface {
	mock.TestingT
	Cleanup(func())
}) *MockEventSubscriber[T] {
	mock := &MockEventSubscriber[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
