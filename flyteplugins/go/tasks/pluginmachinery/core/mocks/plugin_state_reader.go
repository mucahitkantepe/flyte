// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// PluginStateReader is an autogenerated mock type for the PluginStateReader type
type PluginStateReader struct {
	mock.Mock
}

type PluginStateReader_Get struct {
	*mock.Call
}

func (_m PluginStateReader_Get) Return(stateVersion uint8, err error) *PluginStateReader_Get {
	return &PluginStateReader_Get{Call: _m.Call.Return(stateVersion, err)}
}

func (_m *PluginStateReader) OnGet(t interface{}) *PluginStateReader_Get {
	c_call := _m.On("Get", t)
	return &PluginStateReader_Get{Call: c_call}
}

func (_m *PluginStateReader) OnGetMatch(matchers ...interface{}) *PluginStateReader_Get {
	c_call := _m.On("Get", matchers...)
	return &PluginStateReader_Get{Call: c_call}
}

// Get provides a mock function with given fields: t
func (_m *PluginStateReader) Get(t interface{}) (uint8, error) {
	ret := _m.Called(t)

	var r0 uint8
	if rf, ok := ret.Get(0).(func(interface{}) uint8); ok {
		r0 = rf(t)
	} else {
		r0 = ret.Get(0).(uint8)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(t)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type PluginStateReader_GetStateVersion struct {
	*mock.Call
}

func (_m PluginStateReader_GetStateVersion) Return(_a0 uint8) *PluginStateReader_GetStateVersion {
	return &PluginStateReader_GetStateVersion{Call: _m.Call.Return(_a0)}
}

func (_m *PluginStateReader) OnGetStateVersion() *PluginStateReader_GetStateVersion {
	c_call := _m.On("GetStateVersion")
	return &PluginStateReader_GetStateVersion{Call: c_call}
}

func (_m *PluginStateReader) OnGetStateVersionMatch(matchers ...interface{}) *PluginStateReader_GetStateVersion {
	c_call := _m.On("GetStateVersion", matchers...)
	return &PluginStateReader_GetStateVersion{Call: c_call}
}

// GetStateVersion provides a mock function with given fields:
func (_m *PluginStateReader) GetStateVersion() uint8 {
	ret := _m.Called()

	var r0 uint8
	if rf, ok := ret.Get(0).(func() uint8); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint8)
	}

	return r0
}
