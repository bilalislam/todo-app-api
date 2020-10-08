package handler

import (
	"errors"
	"github.com/bilalislam/todo-app-api/pkg/handler/mocks"
	"github.com/bilalislam/todo-app-api/pkg/handler/requests"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MockRequest struct {
	method string
	header string
	body   string
}

func getMockContext(mockRequest *MockRequest) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(mockRequest.method, "/", strings.NewReader(mockRequest.body))
	if mockRequest.header != "" {
		req.Header.Set(echo.HeaderContentType, mockRequest.header)
	}
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return c, rec
}

func TestHandler_GetTasks_ShouldSuccess_WhenTasksAreNotEmpty(t *testing.T) {
	c, rec := getMockContext(&MockRequest{
		method: http.MethodGet,
	})
	c.SetPath("/tasks")

	var tasks []requests.Task
	tasks = append(tasks, requests.Task{
		ID:         1,
		Name:       "buy some milk",
		IsComplete: false,
	})

	mockStore := &mocks.DataStore{}
	mockStore.On("GetTasks").
		Return(tasks)

	h := Handler{
		store: mockStore,
	}

	if assert.NoError(t, h.GetTasks(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestHandler_GetTasks_ShouldFail_WhenTasksAreEmpty(t *testing.T) {
	c, rec := getMockContext(&MockRequest{
		method: http.MethodGet,
	})
	c.SetPath("/tasks")

	mockStore := &mocks.DataStore{}
	mockStore.On("GetTasks").
		Return([]requests.Task{})

	h := Handler{
		store: mockStore,
	}

	if assert.NoError(t, h.GetTasks(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
	}
}

func TestHandler_AddTask_ShouldFail_WhenNotBind(t *testing.T) {
	c, rec := getMockContext(&MockRequest{
		method: http.MethodPost,
		body: `{
			"id": 1
			"name": "buy some milk",
			"isComplete": false
		}`,
		header: echo.MIMEApplicationJSON,
	})

	c.SetPath("/tasks")

	mockStore := &mocks.DataStore{}
	mockStore.On("AddTask")

	h := Handler{
		store: mockStore,
	}

	if assert.NoError(t, h.AddTask(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestHandler_AddTask_ShouldFail_WhenNotValidateToRequest(t *testing.T) {
	c, rec := getMockContext(&MockRequest{
		method: http.MethodPost,
		body: `{
			"id": 1,
			"name": "",
			"isComplete": false
		}`,
		header: echo.MIMEApplicationJSON,
	})
	c.SetPath("/tasks")

	mockStore := &mocks.DataStore{}
	mockStore.On("AddTask")

	h := Handler{
		store: mockStore,
	}

	if assert.NoError(t, h.AddTask(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestHandler_AddTask_ShouldSuccess_WhenValidateToRequest(t *testing.T) {
	var req = `{
			"id": 1,
			"name": "buy some milk",
			"isComplete": false
		}`

	c, rec := getMockContext(&MockRequest{
		method: http.MethodPost,
		body:   req,
		header: echo.MIMEApplicationJSON,
	})
	c.SetPath("/tasks")

	mockStore := &mocks.DataStore{}
	mockStore.On("AddTask", requests.Task{
		ID:         1,
		Name:       "buy some milk",
		IsComplete: false,
	})

	h := Handler{
		store: mockStore,
	}

	if assert.NoError(t, h.AddTask(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestHandler_UpdateTask_ShouldFail_WhenNotBind(t *testing.T) {
	var req = `{
			"id": 1,
			"name": "buy some milk"
			"isComplete": false
		}`

	c, rec := getMockContext(&MockRequest{
		method: http.MethodPut,
		body:   req,
		header: echo.MIMEApplicationJSON,
	})
	c.SetPath("/tasks")

	mockStore := &mocks.DataStore{}
	mockStore.On("UpdateTask", requests.Task{
		ID:         1,
		Name:       "buy some milk",
		IsComplete: false,
	})

	h := Handler{
		store: mockStore,
	}

	if assert.NoError(t, h.UpdateTask(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestHandler_UpdateTask_ShouldFail_WhenTaskNotFound(t *testing.T) {
	var req = `{
			"id": 1,
			"name": "buy some milk",
			"isComplete": false
		}`

	c, rec := getMockContext(&MockRequest{
		method: http.MethodPut,
		body:   req,
		header: echo.MIMEApplicationJSON,
	})
	c.SetPath("/tasks/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	notFound := errors.New("Task was not found")
	mockStore := &mocks.DataStore{}
	mockStore.On("UpdateTask", 1, requests.Task{
		ID:         1,
		Name:       "buy some milk",
		IsComplete: false,
	}).
		Return(notFound)

	h := Handler{
		store: mockStore,
	}

	if assert.NoError(t, h.UpdateTask(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestHandler_UpdateTask_ShouldFail_WhenNotValidateToRequest(t *testing.T) {
	var req = `{
			"id": 1,
			"name": "",
			"isComplete": false
		}`

	c, rec := getMockContext(&MockRequest{
		method: http.MethodPut,
		body:   req,
		header: echo.MIMEApplicationJSON,
	})
	c.SetPath("/tasks/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockStore := &mocks.DataStore{}

	h := Handler{
		store: mockStore,
	}

	if assert.NoError(t, h.UpdateTask(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestHandler_UpdateTask_ShouldSuccess_WhenValidateToRequest(t *testing.T) {
	var req = `{
			"id": 1,
			"name": "buy some milk",
			"isComplete": false
		}`

	c, rec := getMockContext(&MockRequest{
		method: http.MethodPut,
		body:   req,
		header: echo.MIMEApplicationJSON,
	})
	c.SetPath("/tasks/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockStore := &mocks.DataStore{}
	mockStore.On("UpdateTask", 1, requests.Task{
		ID:         1,
		Name:       "buy some milk",
		IsComplete: false,
	}).
		Return(nil)

	h := Handler{
		store: mockStore,
	}

	if assert.NoError(t, h.UpdateTask(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestHandler_DeleteTask_ShouldFail_WhenTaskNotFound(t *testing.T) {

	c, rec := getMockContext(&MockRequest{
		method: http.MethodDelete,
	})
	c.SetPath("/tasks/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	notFound := errors.New("Task was not found")
	mockStore := &mocks.DataStore{}
	mockStore.On("DeleteTask", 1).
		Return(notFound)

	h := Handler{
		store: mockStore,
	}

	if assert.NoError(t, h.DeleteTask(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestHandler_DeleteTask_ShouldSuccess_WhenDeleteTask(t *testing.T) {

	c, rec := getMockContext(&MockRequest{
		method: http.MethodDelete,
	})
	c.SetPath("/tasks/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockStore := &mocks.DataStore{}
	mockStore.On("DeleteTask", 1).
		Return(nil)

	h := Handler{
		store: mockStore,
	}

	if assert.NoError(t, h.DeleteTask(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
