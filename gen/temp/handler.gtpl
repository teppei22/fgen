package {{ .Handler.pkgName }}

import (
	"net/http"
	"strconv"

	"sample_layered/domain/model"
	"sample_layered/usecase"

	"github.com/labstack/echo/v4"
)

type {{ .HandlerInterface.Name }} interface {
	FindByID(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type {{ .UseCase.StructName }} struct {
	{{ .UseCase.StructName }} usecase.{{ .UseCase.InterfaceName }}
}

func NewTask(u usecase.{{ .UseCase.InterfaceName }}) {{ .Handler.InterfaceName }} {
	return &{{ .Handler.StructName }}{
		{{ .UseCase.StructName }}: u,
	}
}

func (h *{{ .Handler.StructName }}) FindByID(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	{{ .ModelName }}, err := h.{{ .UseCase.StructName }}.FindByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, {{ .ModelName }})
}

func (h *{{ .Handler.StructName }}) Create(c echo.Context) error {

	var t model.{{ .ModelName }}
	err := c.Bind(&t)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	{{ .ModelName }}, err := h.{{ .UseCase.StructName }}.Create(&t)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())

	}

	return c.JSON(http.StatusCreated, {{ .ModelName }})
}

func (h *{{ .Handler.StructName }}) Update(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var t model.Task
	t.ID = id
	if err := c.Bind(&t); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	{{ .ModelName }}, err := h.{{ .UseCase.StructName }}.Update(&t)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())

	}

	return c.JSON(http.StatusOK, {{ .ModelName }})
}

func (h *{{ .Handler.StructName }}) Delete(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = h.{{ .UseCase.StructName }}.Delete(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
