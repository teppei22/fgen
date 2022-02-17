package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// routing
	// e.POST("/signup",)
	// e.POST("/login",)

	// api routes
	// api := e.Group("/api")

	return e
}
