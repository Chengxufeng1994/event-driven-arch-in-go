// Code generated by mockery v2.47.0. DO NOT EDIT.

package am

import mock "github.com/stretchr/testify/mock"

// MockMessageSubscriber is an autogenerated mock type for the MessageSubscriber type
type MockMessageSubscriber[I IncomingMessage] struct {
	mock.Mock
}

// Subscribe provides a mock function with given fields: topicName, handler, options
func (_m *MockMessageSubscriber[I]) Subscribe(topicName string, handler MessageHandler[I], options ...SubscriberOption) (Subscription, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, topicName, handler)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Subscribe")
	}

	var r0 Subscription
	var r1 error
	if rf, ok := ret.Get(0).(func(string, MessageHandler[I], ...SubscriberOption) (Subscription, error)); ok {
		return rf(topicName, handler, options...)
	}
	if rf, ok := ret.Get(0).(func(string, MessageHandler[I], ...SubscriberOption) Subscription); ok {
		r0 = rf(topicName, handler, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(string, MessageHandler[I], ...SubscriberOption) error); ok {
		r1 = rf(topicName, handler, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Unsubscribe provides a mock function with given fields:
func (_m *MockMessageSubscriber[I]) Unsubscribe() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Unsubscribe")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockMessageSubscriber creates a new instance of MockMessageSubscriber. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMessageSubscriber[I IncomingMessage](t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMessageSubscriber[I] {
	mock := &MockMessageSubscriber[I]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
