package store

import "errors"

type Task struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	IsComplete bool   `json:"isComplete"`
}

type DataStore struct {
	tasks  []Task
	lastID int
}

func (ds *DataStore) GetTasks() ([]Task, error) {

	return ds.tasks, nil
}

var NotFound = errors.New("Task was not found")

func (ds *DataStore) AddTask(task Task) Task {
	if task.ID == 0 {
		ds.lastID++
		task.ID = ds.lastID
		ds.tasks = append(ds.tasks, task)
		return task
	}

	return task
}

func (ds *DataStore) UpdateTask(id int, task Task) error {
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
			return nil
		}
	}

	return NotFound
}
