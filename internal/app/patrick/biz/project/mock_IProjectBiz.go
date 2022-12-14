// Code generated by mockery v2.14.0. DO NOT EDIT.

package project

import (
	contextx "github.com/blackhorseya/gocommon/pkg/contextx"
	entityproject "github.com/blackhorseya/patrick/internal/pkg/entity/project"

	mock "github.com/stretchr/testify/mock"
)

// MockIProjectBiz is an autogenerated mock type for the IProjectBiz type
type MockIProjectBiz struct {
	mock.Mock
}

// InitProject provides a mock function with given fields: ctx, prj
func (_m *MockIProjectBiz) InitProject(ctx contextx.Contextx, prj *entityproject.Info) error {
	ret := _m.Called(ctx, prj)

	var r0 error
	if rf, ok := ret.Get(0).(func(contextx.Contextx, *entityproject.Info) error); ok {
		r0 = rf(ctx, prj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMockIProjectBiz interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockIProjectBiz creates a new instance of MockIProjectBiz. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockIProjectBiz(t mockConstructorTestingTNewMockIProjectBiz) *MockIProjectBiz {
	mock := &MockIProjectBiz{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
