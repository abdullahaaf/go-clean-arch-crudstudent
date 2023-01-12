package main

import (
	"database/sql"
	"fmt"

	"github.com/abdullahaaf/go-clean-arch-crudstudent/internal/core/usecase"
	"github.com/abdullahaaf/go-clean-arch-crudstudent/internal/handlers/student"
	"github.com/abdullahaaf/go-clean-arch-crudstudent/internal/repositories"
	"github.com/labstack/echo"
)

func main() {

	dbConn, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/go-student")
	if err != nil {
		fmt.Println(err)
	}
	defer dbConn.Close()

	e := echo.New()
	studentRepo := repositories.NewMysqlStudentRepository(dbConn)
	studentUseCase := usecase.NewStudentUseCase(studentRepo)
	student.NewStudentHandler(e, studentUseCase)

	e.Start("localhost:9003")
}
