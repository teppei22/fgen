package controller

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type TaskController interface {
	Interactor usecase.TaskInteractor
}

type createTaskRequest struct {
	Title       string
	Description string
	createdAt   time.Time
	updatedAt   time.Time
}

func (controller *TaskController) Create(c echo.Context) error {
	req := &createTaskRequest{}
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	tasks, err := controller.Interactor.Add(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "ERROR: failed to create tasks")

	}

	return c.JSON(http.StatusCreated, tasks)
}
