package main

import (
	"fmt"

	httpDeliver "github.com/abdullahaaf/go-clean-arch-crudstudent/student/delivery/http"
	"github.com/abdullahaaf/go-clean-arch-crudstudent/student/repository"
	"github.com/abdullahaaf/go-clean-arch-crudstudent/student/usecase"
	"github.com/labstack/echo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root@tcp(127.0.0.1:3306)/go-student?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	e := echo.New()
	studentRepo := repository.NewMysqlStudentRepository(db)
	studentUseCase := usecase.NewStudentUseCase(studentRepo)
	httpDeliver.NewStudentHandler(e, studentUseCase)

	e.Start("localhost:9003")
}
