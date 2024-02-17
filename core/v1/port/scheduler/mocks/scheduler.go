// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	core "codebase/core"

	mock "github.com/stretchr/testify/mock"
)

// Scheduler is an autogenerated mock type for the Scheduler type
type Scheduler struct {
	mock.Mock
}

// Start provides a mock function with given fields: ic
func (_m *Scheduler) Start(ic *core.InternalContext) *core.CustomError {
	ret := _m.Called(ic)

	var r0 *core.CustomError
	if rf, ok := ret.Get(0).(func(*core.InternalContext) *core.CustomError); ok {
		r0 = rf(ic)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.CustomError)
		}
	}

	return r0
}

// NewScheduler creates a new instance of Scheduler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewScheduler(t interface {
	mock.TestingT
	Cleanup(func())
}) *Scheduler {
	mock := &Scheduler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
