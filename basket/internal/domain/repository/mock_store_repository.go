// Code generated by mockery v2.47.0. DO NOT EDIT.

package repository

import (
	context "context"

	entity "github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/entity"
	mock "github.com/stretchr/testify/mock"
)

// MockStoreRepository is an autogenerated mock type for the StoreRepository type
type MockStoreRepository struct {
	mock.Mock
}

// Find provides a mock function with given fields: ctx, storeID
func (_m *MockStoreRepository) Find(ctx context.Context, storeID string) (*entity.Store, error) {
	ret := _m.Called(ctx, storeID)

	if len(ret) == 0 {
		panic("no return value specified for Find")
	}

	var r0 *entity.Store
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.Store, error)); ok {
		return rf(ctx, storeID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.Store); ok {
		r0 = rf(ctx, storeID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Store)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, storeID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockStoreRepository creates a new instance of MockStoreRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockStoreRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockStoreRepository {
	mock := &MockStoreRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
