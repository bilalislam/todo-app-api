package main

import (
	_ "github.com/bilalislam/todo-app-api/docs"
	"github.com/bilalislam/todo-app-api/pkg/handler"
	"github.com/bilalislam/todo-app-api/pkg/store"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
	"net/http"
)

// @title Todo Api
// @version 1.0
// @description This is a simple todo server
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email bilal.islam815@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"@timezone":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}","host":"${host}",` +
			`"method":"${method}","uri":"${uri}","status":${status},"error":"${error}","latency":${latency},` +
			`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
			`"bytes_out":${bytes_out}}` + "\n",
	}))

	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	ds := store.NewDataStore()
	h := handler.NewHandler(&ds)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/", Index)
	e.GET("/tasks", h.GetTasks)
	e.POST("/tasks", h.AddTask)
	e.PUT("/tasks/:id", h.UpdateTask)
	e.DELETE("/tasks/:id", h.DeleteTask)

	e.Logger.Fatal(e.Start(":8080"))
}

func Index(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
}
