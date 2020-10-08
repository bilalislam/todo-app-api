package handler

import (
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
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	}
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return c, rec
}

func TestGetTasks_ShouldSuccess_WhenTasksAreNotEmpty(t *testing.T) {
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

func TestGetTasks_ShouldFail_WhenTasksAreEmpty(t *testing.T) {
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

