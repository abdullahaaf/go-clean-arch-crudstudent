package ports

import "github.com/abdullahaaf/go-clean-arch-crudstudent/internal/core/domain"

type StudentRepository interface {
	GetAll() ([]*domain.Student, error)
	GetByName(name string) *domain.Student
	Store(stud *domain.Student) (int64, error)
	Update(stud *domain.Student) (*domain.Student, error)
	Delete(name string) (bool, error)
}
