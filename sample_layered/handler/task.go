package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/teppei22/fji-codegen/sample_layered/usecase"
)

type TaskHandler interface {
	FindTask(c echo.Context) error
}

type taskHandler struct {
	taskUseCase usecase.TaskUseCase
}

func NewTaskHandler(u usecase.TaskUseCase) TaskHandler {
	return &taskHandler{
		taskUseCase: u,
	}
}

func (h *taskHandler) FindTask(c echo.Context) error {

	var id = c.Param("id")

	tasks, err := h.taskUseCase.FindTask(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "ERROR: failed to get task")

	}

	return c.JSON(http.StatusCreated, tasks)
}
