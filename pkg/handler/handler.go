package handler

import (
	"errors"
	"github.com/bilalislam/todo-app-api/pkg/store"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

var ds = &store.DataStore{}

type Store interface {
	GetTasks() []store.Task
	AddTask(task store.Task) error
	UpdateTask(task store.Task) error
	DeleteTask(id int) error
}

// GetTasks godoc
// @Summary Get All Tasks
// @Tags tasks
// @Accept  json
// @Produce  json
// @Success 200 {object} store.Response
// @Router /tasks [get]
func GetTasks(c echo.Context) error {
	tasks, err := ds.GetTasks()

	if err != nil {
		return c.JSON(http.StatusBadRequest, store.Response{
			Messages: err.Error(),
		})
	}

	if len(tasks) > 0 {
		return c.JSON(http.StatusOK, tasks)
	}

	return c.JSON(http.StatusNotFound, tasks)
}

// AddTask godoc
// @Summary Add a Task
// @Tags tasks
// @Accept  json
// @Produce  json
// @Success 200 {object} store.Response
// @Param tasks body store.Task true "task info"
// @Router /tasks [post]
func AddTask(c echo.Context) error {

	task := new(store.Task)
	err := c.Bind(task)

	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, store.Response{
			Messages: err.Error(),
		})
	}

	err = validateTask(*task)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, store.Response{
			Messages: err.Error(),
		})
	}

	t := ds.AddTask(*task)

	return c.JSON(http.StatusOK, store.Response{
		Result:  t,
		Success: true,
	})
}

func UpdateTask(c echo.Context) error {

	task := new(store.Task)
	err := c.Bind(task)

	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, store.Response{
			Messages: err.Error(),
		})
	}

	err = validateTask(*task)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, store.Response{
			Messages: err.Error(),
		})
	}

	id, _ := strconv.Atoi(c.Param("id"))
	err = ds.UpdateTask(id, *task)

	if err != nil {
		return c.JSON(http.StatusBadRequest, store.Response{
			Messages: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, store.Response{
		Result:  task,
		Success: true,
	})
}

func DeleteTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := ds.DeleteTask(id)

	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, store.Response{
			Messages: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, store.Response{
		Success: true,
	})
}

func validateTask(t store.Task) error {
	if t.Name == "" {
		return errors.New("name could not be empty")
	}
	return nil
}