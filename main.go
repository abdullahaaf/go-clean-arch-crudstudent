package main

import (
	"database/sql"
	"fmt"

	httpDeliver "github.com/abdullahaaf/go-clean-arch-crudstudent/student/delivery/http"
	"github.com/abdullahaaf/go-clean-arch-crudstudent/student/repository"
	"github.com/abdullahaaf/go-clean-arch-crudstudent/student/usecase"
	"github.com/labstack/echo"
)

func main() {

	dbConn, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/go-student")
	if err != nil {
		fmt.Println(err)
	}
	defer dbConn.Close()

	e := echo.New()
	studentRepo := repository.NewMysqlStudentRepository(dbConn)
	studentUseCase := usecase.NewStudentUseCase(studentRepo)
	httpDeliver.NewStudentHandler(e, studentUseCase)

	e.Start("localhost:9003")
}
