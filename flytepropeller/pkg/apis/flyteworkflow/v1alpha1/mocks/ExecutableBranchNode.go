// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	core "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/core"
	mock "github.com/stretchr/testify/mock"

	v1alpha1 "github.com/flyteorg/flyte/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"
)

// ExecutableBranchNode is an autogenerated mock type for the ExecutableBranchNode type
type ExecutableBranchNode struct {
	mock.Mock
}

type ExecutableBranchNode_GetElse struct {
	*mock.Call
}

func (_m ExecutableBranchNode_GetElse) Return(_a0 *string) *ExecutableBranchNode_GetElse {
	return &ExecutableBranchNode_GetElse{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableBranchNode) OnGetElse() *ExecutableBranchNode_GetElse {
	c_call := _m.On("GetElse")
	return &ExecutableBranchNode_GetElse{Call: c_call}
}

func (_m *ExecutableBranchNode) OnGetElseMatch(matchers ...interface{}) *ExecutableBranchNode_GetElse {
	c_call := _m.On("GetElse", matchers...)
	return &ExecutableBranchNode_GetElse{Call: c_call}
}

// GetElse provides a mock function with given fields:
func (_m *ExecutableBranchNode) GetElse() *string {
	ret := _m.Called()

	var r0 *string
	if rf, ok := ret.Get(0).(func() *string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	return r0
}

type ExecutableBranchNode_GetElseFail struct {
	*mock.Call
}

func (_m ExecutableBranchNode_GetElseFail) Return(_a0 *core.Error) *ExecutableBranchNode_GetElseFail {
	return &ExecutableBranchNode_GetElseFail{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableBranchNode) OnGetElseFail() *ExecutableBranchNode_GetElseFail {
	c_call := _m.On("GetElseFail")
	return &ExecutableBranchNode_GetElseFail{Call: c_call}
}

func (_m *ExecutableBranchNode) OnGetElseFailMatch(matchers ...interface{}) *ExecutableBranchNode_GetElseFail {
	c_call := _m.On("GetElseFail", matchers...)
	return &ExecutableBranchNode_GetElseFail{Call: c_call}
}

// GetElseFail provides a mock function with given fields:
func (_m *ExecutableBranchNode) GetElseFail() *core.Error {
	ret := _m.Called()

	var r0 *core.Error
	if rf, ok := ret.Get(0).(func() *core.Error); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Error)
		}
	}

	return r0
}

type ExecutableBranchNode_GetElseIf struct {
	*mock.Call
}

func (_m ExecutableBranchNode_GetElseIf) Return(_a0 []v1alpha1.ExecutableIfBlock) *ExecutableBranchNode_GetElseIf {
	return &ExecutableBranchNode_GetElseIf{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableBranchNode) OnGetElseIf() *ExecutableBranchNode_GetElseIf {
	c_call := _m.On("GetElseIf")
	return &ExecutableBranchNode_GetElseIf{Call: c_call}
}

func (_m *ExecutableBranchNode) OnGetElseIfMatch(matchers ...interface{}) *ExecutableBranchNode_GetElseIf {
	c_call := _m.On("GetElseIf", matchers...)
	return &ExecutableBranchNode_GetElseIf{Call: c_call}
}

// GetElseIf provides a mock function with given fields:
func (_m *ExecutableBranchNode) GetElseIf() []v1alpha1.ExecutableIfBlock {
	ret := _m.Called()

	var r0 []v1alpha1.ExecutableIfBlock
	if rf, ok := ret.Get(0).(func() []v1alpha1.ExecutableIfBlock); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]v1alpha1.ExecutableIfBlock)
		}
	}

	return r0
}

type ExecutableBranchNode_GetIf struct {
	*mock.Call
}

func (_m ExecutableBranchNode_GetIf) Return(_a0 v1alpha1.ExecutableIfBlock) *ExecutableBranchNode_GetIf {
	return &ExecutableBranchNode_GetIf{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableBranchNode) OnGetIf() *ExecutableBranchNode_GetIf {
	c_call := _m.On("GetIf")
	return &ExecutableBranchNode_GetIf{Call: c_call}
}

func (_m *ExecutableBranchNode) OnGetIfMatch(matchers ...interface{}) *ExecutableBranchNode_GetIf {
	c_call := _m.On("GetIf", matchers...)
	return &ExecutableBranchNode_GetIf{Call: c_call}
}

// GetIf provides a mock function with given fields:
func (_m *ExecutableBranchNode) GetIf() v1alpha1.ExecutableIfBlock {
	ret := _m.Called()

	var r0 v1alpha1.ExecutableIfBlock
	if rf, ok := ret.Get(0).(func() v1alpha1.ExecutableIfBlock); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableIfBlock)
		}
	}

	return r0
}