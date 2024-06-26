// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	internal "github.com/eddyvy/tfg-go-cli/internal"
	mock "github.com/stretchr/testify/mock"
)

// CliAddI is an autogenerated mock type for the CliAddI type
type CliAddI struct {
	mock.Mock
}

// ConfigureDatabaseForUpdate provides a mock function with given fields: conf
func (_m *CliAddI) ConfigureDatabaseForUpdate(conf *internal.GlobalConfig) error {
	ret := _m.Called(conf)

	if len(ret) == 0 {
		panic("no return value specified for ConfigureDatabaseForUpdate")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*internal.GlobalConfig) error); ok {
		r0 = rf(conf)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FormatProject provides a mock function with given fields: conf
func (_m *CliAddI) FormatProject(conf *internal.GlobalConfig) error {
	ret := _m.Called(conf)

	if len(ret) == 0 {
		panic("no return value specified for FormatProject")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*internal.GlobalConfig) error); ok {
		r0 = rf(conf)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReadYamlConfig provides a mock function with given fields: projectRelPath
func (_m *CliAddI) ReadYamlConfig(projectRelPath string) (*internal.GlobalConfig, error) {
	ret := _m.Called(projectRelPath)

	if len(ret) == 0 {
		panic("no return value specified for ReadYamlConfig")
	}

	var r0 *internal.GlobalConfig
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*internal.GlobalConfig, error)); ok {
		return rf(projectRelPath)
	}
	if rf, ok := ret.Get(0).(func(string) *internal.GlobalConfig); ok {
		r0 = rf(projectRelPath)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internal.GlobalConfig)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(projectRelPath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TidyProject provides a mock function with given fields: conf
func (_m *CliAddI) TidyProject(conf *internal.GlobalConfig) error {
	ret := _m.Called(conf)

	if len(ret) == 0 {
		panic("no return value specified for TidyProject")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*internal.GlobalConfig) error); ok {
		r0 = rf(conf)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateProject provides a mock function with given fields: conf
func (_m *CliAddI) UpdateProject(conf *internal.GlobalConfig) error {
	ret := _m.Called(conf)

	if len(ret) == 0 {
		panic("no return value specified for UpdateProject")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*internal.GlobalConfig) error); ok {
		r0 = rf(conf)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewCliAddI creates a new instance of CliAddI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCliAddI(t interface {
	mock.TestingT
	Cleanup(func())
}) *CliAddI {
	mock := &CliAddI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
