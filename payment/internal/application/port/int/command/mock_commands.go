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

// AdjustInvoice provides a mock function with given fields: ctx, adjust
func (_m *MockCommands) AdjustInvoice(ctx context.Context, adjust AdjustInvoice) error {
	ret := _m.Called(ctx, adjust)

	if len(ret) == 0 {
		panic("no return value specified for AdjustInvoice")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, AdjustInvoice) error); ok {
		r0 = rf(ctx, adjust)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AuthorizePayment provides a mock function with given fields: ctx, authorize
func (_m *MockCommands) AuthorizePayment(ctx context.Context, authorize AuthorizePayment) error {
	ret := _m.Called(ctx, authorize)

	if len(ret) == 0 {
		panic("no return value specified for AuthorizePayment")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, AuthorizePayment) error); ok {
		r0 = rf(ctx, authorize)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CancelInvoice provides a mock function with given fields: ctx, cancel
func (_m *MockCommands) CancelInvoice(ctx context.Context, cancel CancelInvoice) error {
	ret := _m.Called(ctx, cancel)

	if len(ret) == 0 {
		panic("no return value specified for CancelInvoice")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, CancelInvoice) error); ok {
		r0 = rf(ctx, cancel)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ConfirmPayment provides a mock function with given fields: ctx, confirm
func (_m *MockCommands) ConfirmPayment(ctx context.Context, confirm ConfirmPayment) error {
	ret := _m.Called(ctx, confirm)

	if len(ret) == 0 {
		panic("no return value specified for ConfirmPayment")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ConfirmPayment) error); ok {
		r0 = rf(ctx, confirm)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateInvoice provides a mock function with given fields: ctx, create
func (_m *MockCommands) CreateInvoice(ctx context.Context, create CreateInvoice) error {
	ret := _m.Called(ctx, create)

	if len(ret) == 0 {
		panic("no return value specified for CreateInvoice")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, CreateInvoice) error); ok {
		r0 = rf(ctx, create)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PayInvoice provides a mock function with given fields: ctx, pay
func (_m *MockCommands) PayInvoice(ctx context.Context, pay PayInvoice) error {
	ret := _m.Called(ctx, pay)

	if len(ret) == 0 {
		panic("no return value specified for PayInvoice")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, PayInvoice) error); ok {
		r0 = rf(ctx, pay)
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
