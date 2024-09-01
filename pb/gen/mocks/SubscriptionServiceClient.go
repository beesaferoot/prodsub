// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	context "context"

	gen "github.com/prodsub/pb/gen"
	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// SubscriptionServiceClient is an autogenerated mock type for the SubscriptionServiceClient type
type SubscriptionServiceClient struct {
	mock.Mock
}

// CreateSubscription provides a mock function with given fields: ctx, in, opts
func (_m *SubscriptionServiceClient) CreateSubscription(ctx context.Context, in *gen.SubscriptionCreateRequest, opts ...grpc.CallOption) (*gen.SubscriptionCreateResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateSubscription")
	}

	var r0 *gen.SubscriptionCreateResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *gen.SubscriptionCreateRequest, ...grpc.CallOption) (*gen.SubscriptionCreateResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *gen.SubscriptionCreateRequest, ...grpc.CallOption) *gen.SubscriptionCreateResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.SubscriptionCreateResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *gen.SubscriptionCreateRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteSubscription provides a mock function with given fields: ctx, in, opts
func (_m *SubscriptionServiceClient) DeleteSubscription(ctx context.Context, in *gen.SubscriptionDeleteRequest, opts ...grpc.CallOption) (*gen.SubscriptionDeleteResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteSubscription")
	}

	var r0 *gen.SubscriptionDeleteResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *gen.SubscriptionDeleteRequest, ...grpc.CallOption) (*gen.SubscriptionDeleteResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *gen.SubscriptionDeleteRequest, ...grpc.CallOption) *gen.SubscriptionDeleteResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.SubscriptionDeleteResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *gen.SubscriptionDeleteRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSubscription provides a mock function with given fields: ctx, in, opts
func (_m *SubscriptionServiceClient) GetSubscription(ctx context.Context, in *gen.SubscriptionGetRequest, opts ...grpc.CallOption) (*gen.SubscriptionGetResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetSubscription")
	}

	var r0 *gen.SubscriptionGetResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *gen.SubscriptionGetRequest, ...grpc.CallOption) (*gen.SubscriptionGetResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *gen.SubscriptionGetRequest, ...grpc.CallOption) *gen.SubscriptionGetResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.SubscriptionGetResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *gen.SubscriptionGetRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSubscription provides a mock function with given fields: ctx, in, opts
func (_m *SubscriptionServiceClient) ListSubscription(ctx context.Context, in *gen.SubscriptionListRequest, opts ...grpc.CallOption) (*gen.SubscriptionListResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListSubscription")
	}

	var r0 *gen.SubscriptionListResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *gen.SubscriptionListRequest, ...grpc.CallOption) (*gen.SubscriptionListResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *gen.SubscriptionListRequest, ...grpc.CallOption) *gen.SubscriptionListResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.SubscriptionListResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *gen.SubscriptionListRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateSubscription provides a mock function with given fields: ctx, in, opts
func (_m *SubscriptionServiceClient) UpdateSubscription(ctx context.Context, in *gen.SubscriptionUpdateRequest, opts ...grpc.CallOption) (*gen.SubscriptionUpdateResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateSubscription")
	}

	var r0 *gen.SubscriptionUpdateResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *gen.SubscriptionUpdateRequest, ...grpc.CallOption) (*gen.SubscriptionUpdateResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *gen.SubscriptionUpdateRequest, ...grpc.CallOption) *gen.SubscriptionUpdateResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.SubscriptionUpdateResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *gen.SubscriptionUpdateRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSubscriptionServiceClient creates a new instance of SubscriptionServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSubscriptionServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *SubscriptionServiceClient {
	mock := &SubscriptionServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
