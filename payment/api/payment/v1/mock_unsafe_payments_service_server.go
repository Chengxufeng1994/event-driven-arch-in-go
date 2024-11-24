// Code generated by mockery v2.47.0. DO NOT EDIT.

package paymentv1

import mock "github.com/stretchr/testify/mock"

// MockUnsafePaymentsServiceServer is an autogenerated mock type for the UnsafePaymentsServiceServer type
type MockUnsafePaymentsServiceServer struct {
	mock.Mock
}

// mustEmbedUnimplementedPaymentsServiceServer provides a mock function with given fields:
func (_m *MockUnsafePaymentsServiceServer) mustEmbedUnimplementedPaymentsServiceServer() {
	_m.Called()
}

// NewMockUnsafePaymentsServiceServer creates a new instance of MockUnsafePaymentsServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUnsafePaymentsServiceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUnsafePaymentsServiceServer {
	mock := &MockUnsafePaymentsServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
