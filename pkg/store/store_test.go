package store

import (
	"github.com/bilalislam/todo-app-api/pkg/handler/requests"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDataStore_GetTasks_ShouldSuccess(t *testing.T) {

	ds := NewDataStore()
	task := requests.Task{
		ID:         1,
		Name:       "buy some milk",
		IsComplete: false,
	}
	ds.AddTask(task)
	getTasks := ds.GetTasks()

	assert.Equal(t, len(getTasks), 1)
}

func TestDataStore_AddTask_ShouldSuccess(t *testing.T) {

	ds := NewDataStore()

	task := requests.Task{
		ID:         1,
		Name:       "buy some milk",
		IsComplete: false,
	}
	ds.AddTask(task)

	assert.Equal(t, task.ID, ds.GetTasks()[0].ID)
}

func TestDataStore_UpdateTask_ShouldFail_WhenTaskNotFound(t *testing.T) {
	ds := NewDataStore()

	task := requests.Task{
		ID:         1,
		Name:       "buy some milk",
		IsComplete: false,
	}
	ds.AddTask(task)
	task.Name = "buy some water"
	err := ds.UpdateTask(2, task)

	assert.Equal(t, err, NotFound)
}

func TestDataStore_UpdateTask_ShouldSuccess(t *testing.T) {
	ds := NewDataStore()

	task := requests.Task{
		ID:         1,
		Name:       "buy some milk",
		IsComplete: false,
	}
	ds.AddTask(task)
	task.Name = "buy some water"
	_ = ds.UpdateTask(task.ID, task)

	assert.Equal(t, task.Name, ds.GetTasks()[0].Name)
}

func TestDataStore_DeleteTask_ShouldFail_WhenNotFound(t *testing.T) {
	ds := NewDataStore()

	task := requests.Task{
		ID:         1,
		Name:       "buy some milk",
		IsComplete: false,
	}
	ds.AddTask(task)
	err := ds.DeleteTask(2)

	assert.Equal(t, err, NotFound)
}

func TestDataStore_DeleteTask_ShouldSuccess(t *testing.T) {
	ds := NewDataStore()

	task := requests.Task{
		ID:         1,
		Name:       "buy some milk",
		IsComplete: false,
	}
	ds.AddTask(task)
	_ = ds.DeleteTask(1)

	assert.Equal(t, 0, len(ds.GetTasks()))
}
