// Code generated by mockery v2.47.0. DO NOT EDIT.

package repository

import (
	context "context"

	entity "github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/entity"
	mock "github.com/stretchr/testify/mock"
)

// MockProductRepository is an autogenerated mock type for the ProductRepository type
type MockProductRepository struct {
	mock.Mock
}

// Find provides a mock function with given fields: ctx, productID
func (_m *MockProductRepository) Find(ctx context.Context, productID string) (*entity.Product, error) {
	ret := _m.Called(ctx, productID)

	if len(ret) == 0 {
		panic("no return value specified for Find")
	}

	var r0 *entity.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.Product, error)); ok {
		return rf(ctx, productID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.Product); ok {
		r0 = rf(ctx, productID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, productID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockProductRepository creates a new instance of MockProductRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockProductRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockProductRepository {
	mock := &MockProductRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
