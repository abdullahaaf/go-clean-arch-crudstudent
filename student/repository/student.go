package repository

import "github.com/abdullahaaf/go-clean-arch-crudstudent/student/model"

type StudentRepository interface {
	GetAll() ([]*model.Student, error)
	GetByName(name string) *model.Student
	Store(stud *model.Student) (int64, error)
	Update(stud *model.Student) (*model.Student, error)
	Delete(name string) (bool, error)
}
