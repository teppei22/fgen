package handler

import (
	"net/http"
	"strconv"

	"sample_layered/domain/model"
	"sample_layered/usecase"

	"github.com/labstack/echo/v4"
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

func NewTask(u usecase.TaskUseCase) TaskHandler {
	return &taskHandler{
		taskUseCase: u,
	}
}

func (h *taskHandler) FindByID(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	task, err := h.taskUseCase.FindByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, task)
}

func (h *taskHandler) Create(c echo.Context) error {

	var t model.Task
	err := c.Bind(&t)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	task, err := h.taskUseCase.Create(&t)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())

	}

	return c.JSON(http.StatusCreated, task)
}

func (h *taskHandler) Update(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var t model.Task
	t.ID = id
	if err := c.Bind(&t); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	task, err := h.taskUseCase.Update(&t)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())

	}

	return c.JSON(http.StatusOK, task)
}

func (h *taskHandler) Delete(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = h.taskUseCase.Delete(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
