// Code generated by mockery v2.47.0. DO NOT EDIT.

package repository

import (
	context "context"

	aggregate "github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/domain/aggregate"

	mock "github.com/stretchr/testify/mock"
)

// MockInvoiceRepository is an autogenerated mock type for the InvoiceRepository type
type MockInvoiceRepository struct {
	mock.Mock
}

// Find provides a mock function with given fields: ctx, invoiceID
func (_m *MockInvoiceRepository) Find(ctx context.Context, invoiceID string) (*aggregate.InvoiceAgg, error) {
	ret := _m.Called(ctx, invoiceID)

	if len(ret) == 0 {
		panic("no return value specified for Find")
	}

	var r0 *aggregate.InvoiceAgg
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*aggregate.InvoiceAgg, error)); ok {
		return rf(ctx, invoiceID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *aggregate.InvoiceAgg); ok {
		r0 = rf(ctx, invoiceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*aggregate.InvoiceAgg)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, invoiceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, invoice
func (_m *MockInvoiceRepository) Save(ctx context.Context, invoice *aggregate.InvoiceAgg) error {
	ret := _m.Called(ctx, invoice)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *aggregate.InvoiceAgg) error); ok {
		r0 = rf(ctx, invoice)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, invoice
func (_m *MockInvoiceRepository) Update(ctx context.Context, invoice *aggregate.InvoiceAgg) error {
	ret := _m.Called(ctx, invoice)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *aggregate.InvoiceAgg) error); ok {
		r0 = rf(ctx, invoice)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockInvoiceRepository creates a new instance of MockInvoiceRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockInvoiceRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockInvoiceRepository {
	mock := &MockInvoiceRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}