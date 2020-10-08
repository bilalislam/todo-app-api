// Code generated by mockery v2.3.0. DO NOT EDIT.

package mocks

import (
	requests "github.com/bilalislam/todo-app-api/pkg/handler/requests"
	mock "github.com/stretchr/testify/mock"
)

// DataStore is an autogenerated mock type for the DataStore type
type DataStore struct {
	mock.Mock
}

// AddTask provides a mock function with given fields: task
func (_m *DataStore) AddTask(task requests.Task) {
	_m.Called(task)
}

// DeleteTask provides a mock function with given fields: id
func (_m *DataStore) DeleteTask(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTasks provides a mock function with given fields:
func (_m *DataStore) GetTasks() []requests.Task {
	ret := _m.Called()

	var r0 []requests.Task
	if rf, ok := ret.Get(0).(func() []requests.Task); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]requests.Task)
		}
	}

	return r0
}

// UpdateTask provides a mock function with given fields: id, task
func (_m *DataStore) UpdateTask(id int, task requests.Task) error {
	ret := _m.Called(id, task)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, requests.Task) error); ok {
		r0 = rf(id, task)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
