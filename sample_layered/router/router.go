package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/teppei22/fji-codegen/sample_layered/handler"
	"github.com/teppei22/fji-codegen/sample_layered/infra"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// routing
	// e.POST("/signup",)
	// e.POST("/login",)

	// api routes
	// api := e.Group("/api")

	conn := infra.DBConnect()

	taskHandler := handler.NewTaskHandler(
		usecase.NewTaskUseCase(
			persistence.NewTaskPersistence(conn)
		)
	)
	e.GET("task/:id",taskHandler.FindTask)

	return e
}
