package http

import (
	"net/http"

	"github.com/abdullahaaf/go-clean-arch-crudstudent/student/model"
	"github.com/abdullahaaf/go-clean-arch-crudstudent/student/usecase"
	"github.com/labstack/echo"
)

type StudentHandler struct {
	studentUseCase usecase.StudentUseCase
}

func (stud *StudentHandler) GetAll(c echo.Context) error {
	listStudent, err := stud.studentUseCase.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, listStudent)
}

func (stud *StudentHandler) GetByName(c echo.Context) error {
	detail := stud.studentUseCase.GetByName(c.Param("name"))
	return c.JSON(http.StatusOK, detail)
}

func (stud *StudentHandler) Store(c echo.Context) error {
	var student model.Students
	err := c.Bind(&student)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	store, err := stud.studentUseCase.Store(&student)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, store)
}

func NewStudentHandler(e *echo.Echo, studusecase usecase.StudentUseCase) {
	handler := &StudentHandler{
		studentUseCase: studusecase,
	}

	e.GET("/student", handler.GetAll)
	e.GET("/student/:name", handler.GetByName)
	e.POST("/student", handler.Store)
}
