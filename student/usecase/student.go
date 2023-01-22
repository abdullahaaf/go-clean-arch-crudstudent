package usecase

import (
	"github.com/abdullahaaf/go-clean-arch-crudstudent/student/model"
	"github.com/abdullahaaf/go-clean-arch-crudstudent/student/repository"
)

type StudentUseCase interface {
	GetAll() ([]*model.Students, error)
	GetByName(name string) *model.Students
	Store(student *model.Students) (*model.Students, error)
	Update(student *model.Students) (*model.Students, error)
	Delete(name string) (bool, error)
}

type studentUseCase struct {
	studentRepos repository.StudentRepository
}

func NewStudentUseCase(stud repository.StudentRepository) StudentUseCase {
	return &studentUseCase{studentRepos: stud}
}

func (stud *studentUseCase) GetAll() ([]*model.Students, error) {
	listStudent, err := stud.studentRepos.GetAll()
	if err != nil {
		return nil, err
	}
	return listStudent, nil
}

func (stud *studentUseCase) GetByName(name string) *model.Students {
	detailStudent := stud.studentRepos.GetByName(name)
	return detailStudent
}

func (stud *studentUseCase) Store(student *model.Students) (*model.Students, error) {
	id, err := stud.studentRepos.Store(student)
	if err != nil {
		return nil, err
	}
	student.ID = id

	return student, nil
}

func (stud *studentUseCase) Update(student *model.Students) (*model.Students, error) {
	return stud.studentRepos.Update(student)
}
func (stud *studentUseCase) Delete(name string) (bool, error) {
	return stud.studentRepos.Delete(name)
}
