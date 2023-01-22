package repository

import "github.com/abdullahaaf/go-clean-arch-crudstudent/student/model"

type StudentRepository interface {
	GetAll() ([]*model.Students, error)
	GetByName(name string) *model.Students
	Store(stud *model.Students) (int64, error)
	Update(stud *model.Students) (*model.Students, error)
	Delete(name string) (bool, error)
}
