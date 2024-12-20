// Code generated by mockery v2.47.0. DO NOT EDIT.

package usecase

import (
	command "github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/application/port/in/command"
	aggregate "github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/domain/aggregate"

	context "context"

	mock "github.com/stretchr/testify/mock"

	query "github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/application/port/in/query"
)

// MockCustomerUsecase is an autogenerated mock type for the CustomerUsecase type
type MockCustomerUsecase struct {
	mock.Mock
}

// AuthorizeCustomer provides a mock function with given fields: ctx, authorize
func (_m *MockCustomerUsecase) AuthorizeCustomer(ctx context.Context, authorize command.AuthorizeCustomer) error {
	ret := _m.Called(ctx, authorize)

	if len(ret) == 0 {
		panic("no return value specified for AuthorizeCustomer")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, command.AuthorizeCustomer) error); ok {
		r0 = rf(ctx, authorize)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ChangeSmsNumber provides a mock function with given fields: ctx, changeSmsNumber
func (_m *MockCustomerUsecase) ChangeSmsNumber(ctx context.Context, changeSmsNumber command.ChangeSmsNumber) error {
	ret := _m.Called(ctx, changeSmsNumber)

	if len(ret) == 0 {
		panic("no return value specified for ChangeSmsNumber")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, command.ChangeSmsNumber) error); ok {
		r0 = rf(ctx, changeSmsNumber)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DisableCustomer provides a mock function with given fields: ctx, disable
func (_m *MockCustomerUsecase) DisableCustomer(ctx context.Context, disable command.DisableCustomer) error {
	ret := _m.Called(ctx, disable)

	if len(ret) == 0 {
		panic("no return value specified for DisableCustomer")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, command.DisableCustomer) error); ok {
		r0 = rf(ctx, disable)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EnableCustomer provides a mock function with given fields: ctx, enable
func (_m *MockCustomerUsecase) EnableCustomer(ctx context.Context, enable command.EnableCustomer) error {
	ret := _m.Called(ctx, enable)

	if len(ret) == 0 {
		panic("no return value specified for EnableCustomer")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, command.EnableCustomer) error); ok {
		r0 = rf(ctx, enable)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCustomer provides a mock function with given fields: ctx, _a1
func (_m *MockCustomerUsecase) GetCustomer(ctx context.Context, _a1 query.GetCustomer) (*aggregate.Customer, error) {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetCustomer")
	}

	var r0 *aggregate.Customer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, query.GetCustomer) (*aggregate.Customer, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, query.GetCustomer) *aggregate.Customer); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*aggregate.Customer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, query.GetCustomer) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterCustomer provides a mock function with given fields: ctx, register
func (_m *MockCustomerUsecase) RegisterCustomer(ctx context.Context, register command.RegisterCustomer) error {
	ret := _m.Called(ctx, register)

	if len(ret) == 0 {
		panic("no return value specified for RegisterCustomer")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, command.RegisterCustomer) error); ok {
		r0 = rf(ctx, register)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockCustomerUsecase creates a new instance of MockCustomerUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCustomerUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCustomerUsecase {
	mock := &MockCustomerUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
