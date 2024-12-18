// Code generated by mockery v2.47.0. DO NOT EDIT.

package command

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockCommands is an autogenerated mock type for the Commands type
type MockCommands struct {
	mock.Mock
}

// AssignShoppingList provides a mock function with given fields: ctx, cmd
func (_m *MockCommands) AssignShoppingList(ctx context.Context, cmd AssignShoppingList) error {
	ret := _m.Called(ctx, cmd)

	if len(ret) == 0 {
		panic("no return value specified for AssignShoppingList")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, AssignShoppingList) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CancelShoppingList provides a mock function with given fields: ctx, cmd
func (_m *MockCommands) CancelShoppingList(ctx context.Context, cmd CancelShoppingList) error {
	ret := _m.Called(ctx, cmd)

	if len(ret) == 0 {
		panic("no return value specified for CancelShoppingList")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, CancelShoppingList) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CompleteShoppingList provides a mock function with given fields: ctx, cmd
func (_m *MockCommands) CompleteShoppingList(ctx context.Context, cmd CompleteShoppingList) error {
	ret := _m.Called(ctx, cmd)

	if len(ret) == 0 {
		panic("no return value specified for CompleteShoppingList")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, CompleteShoppingList) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateShoppingList provides a mock function with given fields: ctx, cmd
func (_m *MockCommands) CreateShoppingList(ctx context.Context, cmd CreateShoppingList) error {
	ret := _m.Called(ctx, cmd)

	if len(ret) == 0 {
		panic("no return value specified for CreateShoppingList")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, CreateShoppingList) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InitiateShopping provides a mock function with given fields: ctx, cmd
func (_m *MockCommands) InitiateShopping(ctx context.Context, cmd InitiateShopping) error {
	ret := _m.Called(ctx, cmd)

	if len(ret) == 0 {
		panic("no return value specified for InitiateShopping")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, InitiateShopping) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockCommands creates a new instance of MockCommands. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCommands(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCommands {
	mock := &MockCommands{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
