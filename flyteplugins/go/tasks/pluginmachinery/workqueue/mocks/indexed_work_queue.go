// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	workqueue "github.com/flyteorg/flyteplugins/go/tasks/pluginmachinery/workqueue"
	mock "github.com/stretchr/testify/mock"
)

// IndexedWorkQueue is an autogenerated mock type for the IndexedWorkQueue type
type IndexedWorkQueue struct {
	mock.Mock
}

type IndexedWorkQueue_Get struct {
	*mock.Call
}

func (_m IndexedWorkQueue_Get) Return(info workqueue.WorkItemInfo, found bool, err error) *IndexedWorkQueue_Get {
	return &IndexedWorkQueue_Get{Call: _m.Call.Return(info, found, err)}
}

func (_m *IndexedWorkQueue) OnGet(id string) *IndexedWorkQueue_Get {
	c_call := _m.On("Get", id)
	return &IndexedWorkQueue_Get{Call: c_call}
}

func (_m *IndexedWorkQueue) OnGetMatch(matchers ...interface{}) *IndexedWorkQueue_Get {
	c_call := _m.On("Get", matchers...)
	return &IndexedWorkQueue_Get{Call: c_call}
}

// Get provides a mock function with given fields: id
func (_m *IndexedWorkQueue) Get(id string) (workqueue.WorkItemInfo, bool, error) {
	ret := _m.Called(id)

	var r0 workqueue.WorkItemInfo
	if rf, ok := ret.Get(0).(func(string) workqueue.WorkItemInfo); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(workqueue.WorkItemInfo)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(id)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type IndexedWorkQueue_Queue struct {
	*mock.Call
}

func (_m IndexedWorkQueue_Queue) Return(_a0 error) *IndexedWorkQueue_Queue {
	return &IndexedWorkQueue_Queue{Call: _m.Call.Return(_a0)}
}

func (_m *IndexedWorkQueue) OnQueue(ctx context.Context, id string, once workqueue.WorkItem) *IndexedWorkQueue_Queue {
	c_call := _m.On("Queue", ctx, id, once)
	return &IndexedWorkQueue_Queue{Call: c_call}
}

func (_m *IndexedWorkQueue) OnQueueMatch(matchers ...interface{}) *IndexedWorkQueue_Queue {
	c_call := _m.On("Queue", matchers...)
	return &IndexedWorkQueue_Queue{Call: c_call}
}

// Queue provides a mock function with given fields: ctx, id, once
func (_m *IndexedWorkQueue) Queue(ctx context.Context, id string, once workqueue.WorkItem) error {
	ret := _m.Called(ctx, id, once)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, workqueue.WorkItem) error); ok {
		r0 = rf(ctx, id, once)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type IndexedWorkQueue_Start struct {
	*mock.Call
}

func (_m IndexedWorkQueue_Start) Return(_a0 error) *IndexedWorkQueue_Start {
	return &IndexedWorkQueue_Start{Call: _m.Call.Return(_a0)}
}

func (_m *IndexedWorkQueue) OnStart(ctx context.Context) *IndexedWorkQueue_Start {
	c_call := _m.On("Start", ctx)
	return &IndexedWorkQueue_Start{Call: c_call}
}

func (_m *IndexedWorkQueue) OnStartMatch(matchers ...interface{}) *IndexedWorkQueue_Start {
	c_call := _m.On("Start", matchers...)
	return &IndexedWorkQueue_Start{Call: c_call}
}

// Start provides a mock function with given fields: ctx
func (_m *IndexedWorkQueue) Start(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
