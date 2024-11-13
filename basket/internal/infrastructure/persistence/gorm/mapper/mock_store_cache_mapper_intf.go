// Code generated by mockery v2.47.0. DO NOT EDIT.

package mapper

import (
	entity "github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/entity"
	mock "github.com/stretchr/testify/mock"

	po "github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/infrastructure/persistence/gorm/po"
)

// MockStoreCacheMapperIntf is an autogenerated mock type for the StoreCacheMapperIntf type
type MockStoreCacheMapperIntf struct {
	mock.Mock
}

// ToDomain provides a mock function with given fields: store
func (_m *MockStoreCacheMapperIntf) ToDomain(store po.StoreCache) (entity.Store, error) {
	ret := _m.Called(store)

	if len(ret) == 0 {
		panic("no return value specified for ToDomain")
	}

	var r0 entity.Store
	var r1 error
	if rf, ok := ret.Get(0).(func(po.StoreCache) (entity.Store, error)); ok {
		return rf(store)
	}
	if rf, ok := ret.Get(0).(func(po.StoreCache) entity.Store); ok {
		r0 = rf(store)
	} else {
		r0 = ret.Get(0).(entity.Store)
	}

	if rf, ok := ret.Get(1).(func(po.StoreCache) error); ok {
		r1 = rf(store)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ToPersistence provides a mock function with given fields: store
func (_m *MockStoreCacheMapperIntf) ToPersistence(store entity.Store) (po.StoreCache, error) {
	ret := _m.Called(store)

	if len(ret) == 0 {
		panic("no return value specified for ToPersistence")
	}

	var r0 po.StoreCache
	var r1 error
	if rf, ok := ret.Get(0).(func(entity.Store) (po.StoreCache, error)); ok {
		return rf(store)
	}
	if rf, ok := ret.Get(0).(func(entity.Store) po.StoreCache); ok {
		r0 = rf(store)
	} else {
		r0 = ret.Get(0).(po.StoreCache)
	}

	if rf, ok := ret.Get(1).(func(entity.Store) error); ok {
		r1 = rf(store)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockStoreCacheMapperIntf creates a new instance of MockStoreCacheMapperIntf. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockStoreCacheMapperIntf(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockStoreCacheMapperIntf {
	mock := &MockStoreCacheMapperIntf{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}