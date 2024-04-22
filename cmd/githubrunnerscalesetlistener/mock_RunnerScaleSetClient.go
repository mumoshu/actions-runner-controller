// Code generated by mockery v2.36.1. DO NOT EDIT.

package main

import (
	context "context"

	actions "github.com/actions/actions-runner-controller/github/actions"
	"go.opentelemetry.io/otel"

	mock "github.com/stretchr/testify/mock"
)

// MockRunnerScaleSetClient is an autogenerated mock type for the RunnerScaleSetClient type
type MockRunnerScaleSetClient struct {
	mock.Mock
}

// AcquireJobsForRunnerScaleSet provides a mock function with given fields: ctx, requestIds
func (_m *MockRunnerScaleSetClient) AcquireJobsForRunnerScaleSet(ctx context.Context, requestIds []int64) error {
	ctx, span := otel.Tracer("arc").Start(ctx, "MockRunnerScaleSetClient.AcquireJobsForRunnerScaleSet")
	defer span.End()

	ret := _m.Called(ctx, requestIds)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []int64) error); ok {
		r0 = rf(ctx, requestIds)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetRunnerScaleSetMessage provides a mock function with given fields: ctx, handler, maxCapacity
func (_m *MockRunnerScaleSetClient) GetRunnerScaleSetMessage(ctx context.Context, handler func(*actions.RunnerScaleSetMessage) error, maxCapacity int) error {
	ctx, span := otel.Tracer("arc").Start(ctx, "MockRunnerScaleSetClient.GetRunnerScaleSetMessage")
	defer span.End()

	ret := _m.Called(ctx, handler, maxCapacity)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func(*actions.RunnerScaleSetMessage) error, int) error); ok {
		r0 = rf(ctx, handler, maxCapacity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockRunnerScaleSetClient creates a new instance of MockRunnerScaleSetClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRunnerScaleSetClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRunnerScaleSetClient {
	mock := &MockRunnerScaleSetClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
