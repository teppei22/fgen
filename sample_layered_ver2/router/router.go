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

	// routing
	// e.POST("/signup",)
	// e.POST("/login",)

	// api routes
	// api := e.Group("/api")

	conn := infra.DBConnect()

	// taskHandler := handler.NewTaskHandler(usecase.NewTaskUseCase(persistence.NewTaskPersistence(conn)))

	// e.GET("task/:id", taskHandler.FindByID)
	// e.POST("task", taskHandler.Create)
	// e.PUT("task/:id", taskHandler.Update)
	// e.DELETE("task/:id", taskHandler.Delete)

	taskResource(e, conn)

	return e
}

func taskResource(e *echo.Echo, conn *gorm.DB) {

	taskHandler := handler.NewTask(usecase.NewTask(persistence.NewTask(conn)))

	e.GET("task/:id", taskHandler.FindByID)
	e.POST("task", taskHandler.Create)
	e.PUT("task/:id", taskHandler.Update)
	e.DELETE("task/:id", taskHandler.Delete)

}
