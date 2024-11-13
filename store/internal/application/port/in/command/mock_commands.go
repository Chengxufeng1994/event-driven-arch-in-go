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

// AddProduct provides a mock function with given fields: ctx, cmd
func (_m *MockCommands) AddProduct(ctx context.Context, cmd AddProduct) error {
	ret := _m.Called(ctx, cmd)

	if len(ret) == 0 {
		panic("no return value specified for AddProduct")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, AddProduct) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateStore provides a mock function with given fields: ctx, cmd
func (_m *MockCommands) CreateStore(ctx context.Context, cmd CreateStore) error {
	ret := _m.Called(ctx, cmd)

	if len(ret) == 0 {
		panic("no return value specified for CreateStore")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, CreateStore) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DecreaseProductPrice provides a mock function with given fields: ctx, cmd
func (_m *MockCommands) DecreaseProductPrice(ctx context.Context, cmd DecreaseProductPrice) error {
	ret := _m.Called(ctx, cmd)

	if len(ret) == 0 {
		panic("no return value specified for DecreaseProductPrice")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, DecreaseProductPrice) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DisableParticipation provides a mock function with given fields: ctx, cmd
func (_m *MockCommands) DisableParticipation(ctx context.Context, cmd DisableParticipation) error {
	ret := _m.Called(ctx, cmd)

	if len(ret) == 0 {
		panic("no return value specified for DisableParticipation")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, DisableParticipation) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EnableParticipation provides a mock function with given fields: ctx, cmd
func (_m *MockCommands) EnableParticipation(ctx context.Context, cmd EnableParticipation) error {
	ret := _m.Called(ctx, cmd)

	if len(ret) == 0 {
		panic("no return value specified for EnableParticipation")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, EnableParticipation) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IncreaseProductPrice provides a mock function with given fields: ctx, cmd
func (_m *MockCommands) IncreaseProductPrice(ctx context.Context, cmd IncreaseProductPrice) error {
	ret := _m.Called(ctx, cmd)

	if len(ret) == 0 {
		panic("no return value specified for IncreaseProductPrice")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, IncreaseProductPrice) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RebrandProduct provides a mock function with given fields: ctx, cmd
func (_m *MockCommands) RebrandProduct(ctx context.Context, cmd RebrandProduct) error {
	ret := _m.Called(ctx, cmd)

	if len(ret) == 0 {
		panic("no return value specified for RebrandProduct")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, RebrandProduct) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RebrandStore provides a mock function with given fields: ctx, cmd
func (_m *MockCommands) RebrandStore(ctx context.Context, cmd RebrandStore) error {
	ret := _m.Called(ctx, cmd)

	if len(ret) == 0 {
		panic("no return value specified for RebrandStore")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, RebrandStore) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveProduct provides a mock function with given fields: ctx, cmd
func (_m *MockCommands) RemoveProduct(ctx context.Context, cmd RemoveProduct) error {
	ret := _m.Called(ctx, cmd)

	if len(ret) == 0 {
		panic("no return value specified for RemoveProduct")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, RemoveProduct) error); ok {
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