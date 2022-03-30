package router

import (
	"sample_layered/handler"
	"sample_layered/infra"
	"sample_layered/infra/persistence"
	"sample_layered/usecase"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	conn := infra.DBConnect()

	return e
}