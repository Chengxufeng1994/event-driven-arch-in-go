// Code generated by mockery v2.47.0. DO NOT EDIT.

package query

import mock "github.com/stretchr/testify/mock"

// MockQueries is an autogenerated mock type for the Queries type
type MockQueries struct {
	mock.Mock
}

// NewMockQueries creates a new instance of MockQueries. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockQueries(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockQueries {
	mock := &MockQueries{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}