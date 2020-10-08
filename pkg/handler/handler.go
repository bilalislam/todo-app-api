package handler

import (
	"errors"
	"github.com/bilalislam/todo-app-api/pkg/handler/requests"
	"github.com/bilalislam/todo-app-api/pkg/handler/responses"
	"github.com/bilalislam/todo-app-api/pkg/store"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type DataStore interface {
	GetTasks() []requests.Task
	AddTask(task requests.Task)
	UpdateTask(id int, task requests.Task) error
	DeleteTask(id int) error
}

type Handler struct {
	store DataStore
}

func NewHandler(store *store.DataStore) Handler {
	return Handler{
		store: store,
	}
}

// GetTasks godoc
// @Summary Get All Tasks
// @Tags tasks
// @Accept  json
// @Produce  json
// @Success 200 {object} store.Response
// @Router /tasks [get]
func (h Handler) GetTasks(c echo.Context) error {
	tasks := h.store.GetTasks()

	if len(tasks) == 0 {
		return c.JSON(http.StatusNoContent, responses.Response{
			Success: false,
		})
	}

	return c.JSON(http.StatusOK, tasks)
}

// AddTask godoc
// @Summary Add a Task
// @Tags tasks
// @Accept  json
// @Produce  json
// @Success 200 {object} store.Response
// @Param tasks body store.Task true "task info"
// @Router /tasks [post]
func (h Handler) AddTask(c echo.Context) error {

	task := new(requests.Task)
	err := c.Bind(task)

	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, responses.Response{
			Messages: "An error occurred",
		})
	}

	err = validateTask(*task)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, responses.Response{
			Messages: err.Error(),
		})
	}

	h.store.AddTask(*task)

	return c.JSON(http.StatusCreated, responses.Response{
		Result:  task,
		Success: true,
	})
}

func (h Handler) UpdateTask(c echo.Context) error {

	task := new(requests.Task)
	err := c.Bind(task)

	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, responses.Response{
			Messages: err.Error(),
		})
	}

	err = validateTask(*task)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, responses.Response{
			Messages: err.Error(),
		})
	}

	id, _ := strconv.Atoi(c.Param("id"))
	err = h.store.UpdateTask(id, *task)

	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.Response{
			Messages: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, responses.Response{
		Result:  task,
		Success: true,
	})
}

func (h Handler) DeleteTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.store.DeleteTask(id)

	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, responses.Response{
			Messages: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, responses.Response{
		Success: true,
	})
}

func validateTask(t requests.Task) error {
	if t.Name == "" {
		return errors.New("name could not be empty")
	}
	return nil
}
