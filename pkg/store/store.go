package store

import (
	"errors"
	"github.com/bilalislam/todo-app-api/pkg/handler/requests"
)

type DataStore struct {
	tasks []requests.Task
}

var NotFound = errors.New("Task was not found")

func NewDataStore() DataStore {
	return DataStore{}
}

func (ds *DataStore) GetTasks() []requests.Task {

	return ds.tasks
}

func (ds *DataStore) AddTask(task requests.Task) {
	ds.tasks = append(ds.tasks, task)
}

func (ds *DataStore) UpdateTask(id int, task requests.Task) error {
	for i, t := range ds.tasks {
		if t.ID == id {
			ds.tasks[i] = task
			return nil
		}
	}

	return NotFound
}

func (ds *DataStore) DeleteTask(id int) error {
	s := ds.tasks
	for i, t := range ds.tasks {
		if t.ID == id {
			s[len(s)-1], s[i] = s[i], s[len(s)-1]
			ds.tasks = s[:len(s)-1]
			return nil
		}
	}

	return NotFound
}
