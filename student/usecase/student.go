package usecase

import (
	"github.com/abdullahaaf/go-clean-arch-crudstudent/student/model"
	"github.com/abdullahaaf/go-clean-arch-crudstudent/student/repository"
)

type StudentUseCase interface {
	GetAll() ([]*model.Student, error)
	GetByName(name string) *model.Student
	Store(student *model.Student) (*model.Student, error)
	Update(student *model.Student) (*model.Student, error)
	Delete(name string) (bool, error)
}

type studentUseCase struct {
	studentRepos repository.StudentRepository
}

func NewStudentUseCase(stud repository.StudentRepository) StudentUseCase {
	return &studentUseCase{studentRepos: stud}
}

func (stud *studentUseCase) GetAll() ([]*model.Student, error) {
	listStudent, err := stud.studentRepos.GetAll()
	if err != nil {
		return nil, err
	}
	return listStudent, nil
}

func (stud *studentUseCase) GetByName(name string) *model.Student {
	detailStudent := stud.studentRepos.GetByName(name)
	return detailStudent
}

func (stud *studentUseCase) Store(student *model.Student) (*model.Student, error) {
	id, err := stud.studentRepos.Store(student)
	if err != nil {
		return nil, err
	}
	student.ID = id

	return student, nil
}

func (stud *studentUseCase) Update(student *model.Student) (*model.Student, error) {
	return stud.studentRepos.Update(student)
}
func (stud *studentUseCase) Delete(name string) (bool, error) {
	return stud.studentRepos.Delete(name)
}
