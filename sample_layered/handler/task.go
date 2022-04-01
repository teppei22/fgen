package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/teppei22/fgen/sample_layered/usecase"
	"github.com/teppei22/fji-codegen/sample_layered/domain/model"
)

type TaskHandler interface {
	FindByID(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type taskHandler struct {
	taskUseCase usecase.TaskUseCase
}

func NewTaskHandler(u usecase.TaskUseCase) TaskHandler {
	return &taskHandler{
		taskUseCase: u,
	}
}

func (h *taskHandler) FindByID(c echo.Context) error {

	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("invalid request format"))
	}

	task, err := h.taskUseCase.FindByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("failed to get task"))
	}

	return c.JSON(http.StatusOK, task)
}

func (h *taskHandler) Create(c echo.Context) error {

	var req model.TaskRequest
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New(""))
	}

	tasks, err := h.taskUseCase.FindByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "ERROR: failed to get task")

	}

	return c.JSON(http.StatusCreated, tasks)
}

func (h *taskHandler) Update(c echo.Context) error {

	id := c.Param("id")
	task, err := h.taskUseCase.FindByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "ERROR: failed to get task")

	}

	req := model.TaskRequest
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New(""))
	}

	updatedTask, err := h.taskUsecase.Update(id, req.Title, req.Content)

	return c.JSON(http.StatusOK, tasks)
}

func (h *taskHandler) Delete(c echo.Context) error {

	var id = c.Param("id")

	task, err := h.taskUseCase.FindByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "ERROR: failed to get task")

	}

	err = th.taskUsecase.Delete(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "failed to delete task")
	}

	return c.NoContent(http.StatusOK)
}
