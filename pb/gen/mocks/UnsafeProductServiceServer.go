// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UnsafeProductServiceServer is an autogenerated mock type for the UnsafeProductServiceServer type
type UnsafeProductServiceServer struct {
	mock.Mock
}

// mustEmbedUnimplementedProductServiceServer provides a mock function with given fields:
func (_m *UnsafeProductServiceServer) mustEmbedUnimplementedProductServiceServer() {
	_m.Called()
}

// NewUnsafeProductServiceServer creates a new instance of UnsafeProductServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUnsafeProductServiceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *UnsafeProductServiceServer {
	mock := &UnsafeProductServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
