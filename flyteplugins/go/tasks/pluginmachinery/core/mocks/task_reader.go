// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	flyteidlcore "github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core"

	mock "github.com/stretchr/testify/mock"

	storage "github.com/flyteorg/flytestdlib/storage"
)

// TaskReader is an autogenerated mock type for the TaskReader type
type TaskReader struct {
	mock.Mock
}

type TaskReader_Path struct {
	*mock.Call
}

func (_m TaskReader_Path) Return(_a0 storage.DataReference, _a1 error) *TaskReader_Path {
	return &TaskReader_Path{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *TaskReader) OnPath(ctx context.Context) *TaskReader_Path {
	c_call := _m.On("Path", ctx)
	return &TaskReader_Path{Call: c_call}
}

func (_m *TaskReader) OnPathMatch(matchers ...interface{}) *TaskReader_Path {
	c_call := _m.On("Path", matchers...)
	return &TaskReader_Path{Call: c_call}
}

// Path provides a mock function with given fields: ctx
func (_m *TaskReader) Path(ctx context.Context) (storage.DataReference, error) {
	ret := _m.Called(ctx)

	var r0 storage.DataReference
	if rf, ok := ret.Get(0).(func(context.Context) storage.DataReference); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(storage.DataReference)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type TaskReader_Read struct {
	*mock.Call
}

func (_m TaskReader_Read) Return(_a0 *flyteidlcore.TaskTemplate, _a1 error) *TaskReader_Read {
	return &TaskReader_Read{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *TaskReader) OnRead(ctx context.Context) *TaskReader_Read {
	c_call := _m.On("Read", ctx)
	return &TaskReader_Read{Call: c_call}
}

func (_m *TaskReader) OnReadMatch(matchers ...interface{}) *TaskReader_Read {
	c_call := _m.On("Read", matchers...)
	return &TaskReader_Read{Call: c_call}
}

// Read provides a mock function with given fields: ctx
func (_m *TaskReader) Read(ctx context.Context) (*flyteidlcore.TaskTemplate, error) {
	ret := _m.Called(ctx)

	var r0 *flyteidlcore.TaskTemplate
	if rf, ok := ret.Get(0).(func(context.Context) *flyteidlcore.TaskTemplate); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flyteidlcore.TaskTemplate)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
